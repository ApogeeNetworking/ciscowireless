package ciscowireless

import (
	"github.com/ApogeeNetworking/ciscowireless/accesspoints"
	"github.com/ApogeeNetworking/ciscowireless/requests"
	"github.com/ApogeeNetworking/ciscowireless/rf"
	"github.com/ApogeeNetworking/ciscowireless/sites"
	"github.com/ApogeeNetworking/ciscowireless/ssh"
	"github.com/ApogeeNetworking/ciscowireless/wlans"
)

// Service ...
type Service struct {
	AccessPoints *accesspoints.Service
	Wlans        *wlans.Service
	Sites        *sites.Service
	Rf           *rf.Service
	SSH          *ssh.Service
}

// NewService ...
func NewService(host, user, pass, enablePass string, insecureSSL bool) *Service {
	req := requests.NewService(host, user, pass, insecureSSL)
	return &Service{
		AccessPoints: accesspoints.NewService(req),
		Wlans:        wlans.NewService(req),
		Sites:        sites.NewService(req, jsonContains),
		Rf:           rf.NewService(req, jsonContains),
		SSH:          ssh.NewService(host, user, pass, enablePass),
	}
}

func jsonContains(keys []string, value string) bool {
	for _, key := range keys {
		if value == key {
			return true
		}
	}
	return false
}
