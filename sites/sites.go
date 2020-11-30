package sites

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

// GetApJoinProfiles ...
func (s *Service) GetApJoinProfiles() ([]ApJoinProfileSummary, error) {
	uri := "/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/ap-cfg-profiles/ap-cfg-profile"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	type Resp struct {
		Response []ApJoinProfileSummary `json:"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profile"`
	}
	var cResp Resp
	err = json.NewDecoder(res.Body).Decode(&cResp)
	if err != nil {
		fmt.Println(err)
	}
	return cResp.Response, nil
}

// Tag ...
type Tag struct {
	Name          string `json:"site-tag-name"`
	ApJoinProfile string `json:"ap-join-profile"`
}

// GetTags ...
func (s *Service) GetTags() ([]Tag, error) {
	req, err := s.http.GenerateRequest(
		"/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs/site-tag-config",
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	type resp struct {
		Response []Tag `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-config"`
	}
	var cResp resp
	json.NewDecoder(res.Body).Decode(&cResp)
	return cResp.Response, nil
}
