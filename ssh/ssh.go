package ssh

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ApogeeNetworking/ciscowireless/ciscotypes"
	"github.com/ApogeeNetworking/gonetssh"
	"github.com/ApogeeNetworking/gonetssh/universal"
)

var (
	contains   = strings.Contains
	split      = strings.Split
	lower      = strings.ToLower
	replaceAll = strings.ReplaceAll
)

// Service ...
type Service struct {
	universal.Device
}

// NewService ...
func NewService(host, user, pass, enablePass string) *Service {
	client, _ := gonetssh.NewDevice(
		host, user, pass, enablePass, gonetssh.DType.CiscoIOS,
	)
	var service Service
	service.Device = client
	return &service
}

// SetApName ...
func (s *Service) SetApName(oldName, newName string) {
	cmd := fmt.Sprintf("ap name %s name %s", oldName, newName)
	s.SendCmd(cmd)
}

// GetApCdp ...
func (s *Service) GetApCdp(apName string) ciscotypes.ApCdp {
	cmd := fmt.Sprintf("sh ap name %s cdp neighbor detail", apName)
	output, _ := s.SendCmd(cmd)
	// Workaround for Cisco BugID: CSCvp81958
	erstr := "% No connections to Shell Manager available"
	if contains(output, erstr) {
		time.Sleep(250 * time.Millisecond)
		return s.GetApCdp(apName)
	}
	return s.parseCdp(output)
}

// GetApEthIntf ...
func (s *Service) GetApEthIntf(apName string) ciscotypes.ApEthIntf {
	cmd := fmt.Sprintf("sh ap name %s eth stat", apName)
	output, _ := s.SendCmd(cmd)
	// Workaround for Cisco BugID: CSCvp81958
	erstr := "% No connections to Shell Manager available"
	if contains(output, erstr) {
		time.Sleep(250 * time.Millisecond)
		return s.GetApEthIntf(apName)
	}
	return s.parseApEthIntf(output)
}

// GetApLanPorts used to get Port Status on Hospitality APs
// Non-Hosp APs will return an ApLanPort with LEN 0
func (s *Service) GetApLanPorts(apName string) []ciscotypes.ApLanPort {
	cmd := fmt.Sprintf("sh ap name %s lan port summary", apName)
	output, _ := s.SendCmd(cmd)
	var apLanPorts []ciscotypes.ApLanPort
	lanPortRe := regexp.MustCompile(`LAN(\d)`)
	lines := split(output, "\n")
	for _, line := range lines {
		switch {
		case lanPortRe.MatchString(line):
			line = trimWS(line)
			psplit := split(line, " ")
			portMatch := lanPortRe.FindStringSubmatch(psplit[0])
			if len(portMatch) > 1 {
				portID, _ := strconv.Atoi(portMatch[1])
				lanState := ciscotypes.LanPortState(lower(psplit[1]))
				apLanPorts = append(apLanPorts, ciscotypes.ApLanPort{
					ID:     portID,
					Status: lanState,
				})
			}
		}
	}
	return apLanPorts
}

// UpdateApLanPortState for hospitality AP's enable|disable LAN eth Ports
func (s *Service) UpdateApLanPortState(apName string, state ciscotypes.LanPortState, portID int) error {
	// State can Either be ENABLE|DISABLE
	cmd := fmt.Sprintf("ap name %s lan port-id %d %s", apName, portID, state)
	out, _ := s.SendCmd(cmd)
	if strings.Contains(out, "Invalid") {
		errStr := fmt.Sprintf("error setting lan port for %s", apName)
		return errors.New(errStr)
	}
	return nil
}

// GetClientCount ...
func (s *Service) GetClientCount() int {
	var numClients int
	re := regexp.MustCompile(`Total\sClients\s+:\s+(\S+)`)
	out, _ := s.SendCmd("sh wireless sum")
	matches := re.FindStringSubmatch(out)
	if len(matches) == 2 {
		numClients, _ = strconv.Atoi(matches[1])
	}
	return numClients
}

// GetApsSupported retrieve total number of APs WLC can support
func (s *Service) GetApsSupported() int {
	var apsSupported int
	re := regexp.MustCompile(`Max\s+APs\s+supported\s+:\s+(\S+)`)
	out, _ := s.SendCmd("sh wireless sum")
	matches := re.FindStringSubmatch(out)
	if len(matches) == 2 {
		apsSupported, _ = strconv.Atoi(matches[1])
	}
	return apsSupported
}
