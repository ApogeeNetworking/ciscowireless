package ssh

import (
	"regexp"

	"github.com/ApogeeNetworking/ciscowireless/ciscotypes"
)

func (s *Service) parseCdp(output string) ciscotypes.ApCdp {
	intfRe := regexp.MustCompile(`Interface\s:\s(\S+)`)
	rmSwRe := regexp.MustCompile(`Device\sID\s:\s(.*)`)
	rmSwIntfRe := regexp.MustCompile(`Port\sID.*:\s(.*)`)
	rmSwIPRe := regexp.MustCompile(`Entry\saddress\(es\)\s:\s(.*)`)

	intfMatch := intfRe.FindStringSubmatch(output)
	rmSwMatch := rmSwRe.FindStringSubmatch(output)
	rmSwIntfMatch := rmSwIntfRe.FindStringSubmatch(output)
	rmSwIPMatch := rmSwIPRe.FindStringSubmatch(output)
	var cdp ciscotypes.ApCdp
	if len(intfMatch) == 2 {
		cdp.LocalIntf = replaceAll(intfMatch[1], "\r", "")
	}
	if len(rmSwMatch) == 2 {
		cdp.RemoteSw = replaceAll(rmSwMatch[1], "\r", "")
	}
	if len(rmSwIntfMatch) == 2 {
		cdp.RemoteIntf = replaceAll(rmSwIntfMatch[1], "\r", "")
	}
	if len(rmSwIPMatch) == 2 {
		cdp.RemoteSwIPAddr = replaceAll(rmSwIPMatch[1], "\r", "")
	}
	return cdp
}

func (s *Service) parseApEthIntf(output string) ciscotypes.ApEthIntf {
	intfRe := regexp.MustCompile(`Gigabit\S+`)
	lines := split(output, "\n")
	var ethIntf ciscotypes.ApEthIntf
	for _, line := range lines {
		line = trimWS(line)
		if intfRe.FindString(line) == "" {
			continue
		}
		ethSplit := split(line, " ")
		if len(ethSplit) == 8 {
			ethIntf = ciscotypes.ApEthIntf{
				Name:   ethSplit[0],
				Status: lower(ethSplit[1]),
				Speed:  ethSplit[2] + ethSplit[3],
				TxRcv:  ethSplit[5] + "/" + ethSplit[4],
				Drops:  ethSplit[6],
			}
		}
	}
	return ethIntf
}

// Trim Line for All MultiSpaced White Space for easier parsing
func trimWS(text string) string {
	tsRe := regexp.MustCompile(`\s+`)
	return tsRe.ReplaceAllString(text, " ")
}
