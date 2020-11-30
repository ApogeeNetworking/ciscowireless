package wlans

import (
	"encoding/json"

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

// NewWlan ...
func NewWlan() {}

// NewPolicyProfile ...
func NewPolicyProfile() {}

// NewPolicyTag ...
func NewPolicyTag() {}

// Get ...
func (s *Service) Get() ([]WLAN, error) {
	uri := "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries/wlan-cfg-entry"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	type resp struct {
		Response []WLAN `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entry"`
	}
	var cResp resp
	json.NewDecoder(res.Body).Decode(&cResp)
	return cResp.Response, nil
}

// GetPolicies ...
func (s *Service) GetPolicies() ([]Policy, error) {
	uri := "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-policies/wlan-policy"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	type resp struct {
		Response []Policy `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policy"`
	}
	var cResp resp
	json.NewDecoder(res.Body).Decode(&cResp)
	return cResp.Response, nil
}

// GetPolicyTags ...
func (s *Service) GetPolicyTags() ([]PolicyTag, error) {
	uri := "/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries/policy-list-entry"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	type resp struct {
		Response []PolicyTag `json:"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entry"`
	}
	var cResp resp
	json.NewDecoder(res.Body).Decode(&cResp)
	return cResp.Response, nil
}
