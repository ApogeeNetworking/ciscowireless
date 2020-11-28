package ciscowireless

import (
	"github.com/ApogeeNetworking/ciscowireless/accesspoints"
	"github.com/ApogeeNetworking/ciscowireless/requests"
	"github.com/ApogeeNetworking/ciscowireless/wlans"
)

// Service ...
type Service struct {
	AccessPoints *accesspoints.Service
	Wlans        *wlans.Service
}

// NewService ...
func NewService(host, user, pass string, insecureSSL bool) *Service {
	req := requests.NewService(host, user, pass, insecureSSL)
	return &Service{
		AccessPoints: accesspoints.NewService(req),
		Wlans:        wlans.NewService(req),
	}
}
