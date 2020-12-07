package requests

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Service ...
type Service struct {
	http     *http.Client
	baseURL  string
	user     string
	password string
}

// NewService ...
func NewService(host, user, pass string, insecureSSL bool) *Service {
	return &Service{
		user:     user,
		password: pass,
		baseURL:  fmt.Sprintf("https://%s/restconf/data", host),
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecureSSL,
				},
			},
			Timeout: 90 * time.Second,
		},
	}
}

// GenerateRequest ...
func (s *Service) GenerateRequest(uri, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, s.baseURL+uri, body)
	if err != nil {
		return nil, fmt.Errorf("unabled to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/yang-data+json")
	req.Header.Set("Accept", "application/yang-data+json")
	req.SetBasicAuth(s.user, s.password)
	return req, nil
}

// CreateReqBody ...
func (s *Service) CreateReqBody(v interface{}) (*bytes.Reader, error) {
	payload, err := json.Marshal(&v)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(payload), nil
}

// MakeRequest ...
func (s *Service) MakeRequest(req *http.Request) (*http.Response, error) {
	resp, err := s.http.Do(req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode == 401 {
		return resp, errors.New(resp.Status)
	}
	return s.http.Do(req)
}
