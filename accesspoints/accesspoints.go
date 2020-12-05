package accesspoints

import (
	"encoding/json"
	"fmt"

	"github.com/ApogeeNetworking/ciscowireless/requests"
)

// Service ...
type Service struct {
	http *requests.Service
}

// NewService ...
func NewService(req *requests.Service) *Service {
	return &Service{http: req}
}

func (s *Service) getWtpMac(ethMac string) (string, error) {
	uri := fmt.Sprintf(
		"/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-mac-wtp-mac-map=%s",
		ethMac,
	)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return "", err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	resp := struct {
		Response struct {
			EthMac string `json:"ethernet-mac"`
			WtpMac string `json:"wtp-mac"`
		} `json:"Cisco-IOS-XE-wireless-access-point-oper:ethernet-mac-wtp-mac-map"`
	}{}
	json.NewDecoder(res.Body).Decode(&resp)
	return resp.Response.WtpMac, nil
}

// GetApSummary ...
func (s *Service) GetApSummary() ([]Ap, error) {
	uri := "/openconfig-ap-manager:joined-aps/joined-ap"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resp := struct {
		Response []struct {
			Hostname string `json:"hostname"`
			State    struct {
				MacAddr string `json:"mac"`
				Serial  string `json:"serial"`
				Model   string `json:"model"`
				IPAddr  string `json:"ipv4"`
			} `json:"state"`
		} `json:"openconfig-ap-manager:joined-ap"`
	}{}
	json.NewDecoder(res.Body).Decode(&resp)
	var aps []Ap
	for _, rAp := range resp.Response {
		aps = append(aps, Ap{
			Name:    rAp.Hostname,
			MacAddr: rAp.State.MacAddr,
			Serial:  rAp.State.Serial,
			Model:   rAp.State.Model,
		})
	}
	return aps, nil
}

// GetOne ...
func (s *Service) GetOne(macAddr string) (Ap, error) {
	// the mac address needed is the WTP Radio MAC Address
	// however, we are receiving the Eth MAC so first we have
	// to take the Eth MAC to find the WTP MAC
	wtpMac, err := s.getWtpMac(macAddr)
	if err != nil {
		return Ap{}, err
	}
	uri := fmt.Sprintf(
		"/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data=%s",
		wtpMac,
	)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return Ap{}, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return Ap{}, fmt.Errorf("unable to make request: %s", err)
	}
	if res.StatusCode == 401 {
		fmt.Println(res.Status)
		return s.GetOne(macAddr)
	}
	defer res.Body.Close()
	resp := struct {
		Response capWapResp `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-data"`
	}{}
	json.NewDecoder(res.Body).Decode(&resp)
	ap := Ap{
		Name:    resp.Response.Name,
		MacAddr: resp.Response.Detail.Info.Board.MacAddr,
		Serial:  resp.Response.Detail.Info.Board.Serial,
		Model:   resp.Response.Detail.Info.Model.Model,
		Tag: ApTag{
			Policy: resp.Response.TagInfo.Policy.Name,
			Site:   resp.Response.TagInfo.Site.Name,
			Rf:     resp.Response.TagInfo.Rf.Name,
		},
	}
	return ap, nil
}

// Get ...
func (s *Service) Get() ([]Ap, error) {
	uri := "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("unable to make request: %s", err)
	}
	defer res.Body.Close()
	// d, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(d))
	type resp struct {
		Response []capWapResp `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-data"`
	}
	var cResp resp
	json.NewDecoder(res.Body).Decode(&cResp)
	var aps []Ap
	for _, capWap := range cResp.Response {
		aps = append(aps, Ap{
			Name:    capWap.Name,
			MacAddr: capWap.Detail.Info.Board.MacAddr,
			IPAddr:  capWap.IPAddr,
			Serial:  capWap.Detail.Info.Board.Serial,
			Model:   capWap.Detail.Info.Model.Model,
			Tag: ApTag{
				Policy: capWap.TagInfo.Policy.Name,
				Site:   capWap.TagInfo.Site.Name,
				Rf:     capWap.TagInfo.Rf.Name,
			},
		})
	}
	return aps, nil
}

// SetName ...
func (s *Service) SetName(oldName, newName string) (int, error) {
	ap := ApNameUpdate{OldName: oldName, NewName: newName}
	uri := "/Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-name"
	reqBody := struct {
		Ap ApNameUpdate `json:"Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-name"`
	}{Ap: ap}
	body, err := s.http.CreateReqBody(&reqBody)
	if err != nil {
		return 0, err
	}
	req, err := s.http.GenerateRequest(uri, "POST", body)
	if err != nil {
		return 0, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return 0, fmt.Errorf("unable to make request: %s", err)
	}
	return res.StatusCode, nil
}

// Reboot ...
func (s *Service) Reboot(options ApOptions) {
	uri := "/Cisco-IOS-XE-wireless-access-point-cmd-rpc:ap-reset"
	reqBody := struct {
		Options ApOptions `json:"Cisco-IOS-XE-wireless-access-point-cmd-rpc:ap-reset"`
	}{Options: options}
	body, err := s.http.CreateReqBody(&reqBody)
	if err != nil {
	}
	req, err := s.http.GenerateRequest(uri, "POST", body)
	if err != nil {
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
	}
	fmt.Println(res.StatusCode)
}

// GetTagsFromAp ...
func (s *Service) GetTagsFromAp(macAddr string) (ApTagCfg, error) {
	uri := fmt.Sprintf(
		"/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag=%s",
		macAddr,
	)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return ApTagCfg{}, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return ApTagCfg{}, fmt.Errorf("unable to make request: %s", err)
	}
	defer res.Body.Close()
	resp := struct {
		TagResp ApTagCfg `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
	}{}
	json.NewDecoder(res.Body).Decode(&resp)
	return resp.TagResp, nil
}

// UpdateApTagCfg ...
func (s *Service) UpdateApTagCfg(tagCfg ApTagCfg) (status int, err error) {
	uri := "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag"
	reqBody := struct {
		Cfg ApTagCfg `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
	}{Cfg: tagCfg}
	body, err := s.http.CreateReqBody(&reqBody)
	if err != nil {
		return 0, err
	}
	req, err := s.http.GenerateRequest(uri, "PATCH", body)
	if err != nil {
		return 0, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return 0, fmt.Errorf("unable to make request: %s", err)
	}
	return res.StatusCode, nil
}

// BulkUpdateTagCfg ...
func (s *Service) BulkUpdateTagCfg(cfgs []ApTagCfg) (status int, err error) {
	uri := "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag"
	reqBody := struct {
		Cfgs []ApTagCfg `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
	}{Cfgs: cfgs}
	body, err := s.http.CreateReqBody(&reqBody)
	if err != nil {
		return 0, err
	}
	req, err := s.http.GenerateRequest(uri, "PATCH", body)
	if err != nil {
		return 0, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return 0, fmt.Errorf("unable to make request: %s", err)
	}
	return res.StatusCode, nil
}
