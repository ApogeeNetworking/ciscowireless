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

// GetOne ...
func (s *Service) GetOne(macAddr string) (Ap, error) {
	uri := fmt.Sprintf(
		"/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data=%s",
		macAddr,
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
	type resp struct {
		Response capWapResp `json:"Cisco-IOS-XE-wireless-access-point-oper:capwap-data"`
	}
	var cResp resp
	json.NewDecoder(res.Body).Decode(&cResp)
	ap := Ap{
		Name:    cResp.Response.Name,
		MacAddr: cResp.Response.Mac,
		Serial:  cResp.Response.Detail.Info.Board.Serial,
		Model:   cResp.Response.Detail.Info.Model.Model,
		Tag: ApTag{
			Policy: cResp.Response.TagInfo.Policy.Name,
			Site:   cResp.Response.TagInfo.Site.Name,
			Rf:     cResp.Response.TagInfo.Rf.Name,
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
			MacAddr: capWap.Mac,
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
	type resp struct {
		TagResp ApTagCfg `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
	}
	var r resp
	json.NewDecoder(res.Body).Decode(&r)
	return r.TagResp, nil
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
