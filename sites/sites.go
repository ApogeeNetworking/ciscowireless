package sites

import (
	"encoding/json"
	"reflect"

	"github.com/ApogeeNetworking/ciscowireless/requests"
)

// Service ...
type Service struct {
	http         *requests.Service
	jsonContains func(keys []string, val string) bool
}

// NewService ...
func NewService(req *requests.Service, contains func(keys []string, val string) bool) *Service {
	return &Service{
		http:         req,
		jsonContains: contains,
	}
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
	var tagsMap map[string][]interface{}
	json.NewDecoder(res.Body).Decode(&tagsMap)
	tags := tagsMap["Cisco-IOS-XE-wireless-site-cfg:site-tag-config"]
	var siteTags []Tag
	for _, tagMap := range tags {
		v := reflect.ValueOf(tagMap)
		if v.Kind() == reflect.Map {
			var siteTag Tag
			var keys []string
			for _, k := range v.MapKeys() {
				key := k.String()
				keys = append(keys, key)
				switch key {
				case "site-tag-name":
					val := (v.MapIndex(k)).Interface().(string)
					siteTag.Name = val
				case "description":
					val := (v.MapIndex(k)).Interface().(string)
					siteTag.Description = val
				case "ap-join-profile":
					val := (v.MapIndex(k)).Interface().(string)
					siteTag.ApJoinProfile = val
				case "is-local-site":
					val := (v.MapIndex(k)).Interface().(bool)
					siteTag.IsLocalSite = val
				case "flex-profile":
					val := (v.MapIndex(k)).Interface().(string)
					siteTag.FlexProfile = val
				case "fabric-control-plane-name":
					val := (v.MapIndex(k)).Interface().(string)
					siteTag.FabricCntrlPlaneName = val
				case "image-download-profile-name":
					val := (v.MapIndex(k)).Interface().(string)
					siteTag.ImageDownloadProfileName = val
				case "arp-caching":
					val := (v.MapIndex(k)).Interface().(bool)
					siteTag.ArpCaching = val
				}
			}
			isLocalSite := s.jsonContains(keys, "is-local-site")
			if !isLocalSite {
				siteTag.IsLocalSite = true
			}
			if !s.jsonContains(keys, "arp-caching") {
				siteTag.ArpCaching = true
			}
			if !s.jsonContains(keys, "flex-profile") && isLocalSite {
				siteTag.FlexProfile = "default-flex-profile"
			}
			siteTags = append(siteTags, siteTag)
		}
	}
	return siteTags, err
}

// GetApJoinProfiles ...
func (s *Service) GetApJoinProfiles() ([]ApJoinProfile, error) {
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
	var apJoinProfiles []ApJoinProfile
	var profilesMap map[string][]interface{}
	json.NewDecoder(res.Body).Decode(&profilesMap)
	profiles := profilesMap["Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profile"]
	for _, profileMap := range profiles {
		profile := s.handleJoinProfiles(profileMap)
		apJoinProfiles = append(apJoinProfiles, profile)
	}
	return apJoinProfiles, nil
}

func (s *Service) handleJoinProfiles(m interface{}) ApJoinProfile {
	var profile ApJoinProfile
	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Map {
		var keys []string
		for _, k := range v.MapKeys() {
			key := k.String()
			keys = append(keys, key)
			val := v.MapIndex(k).Interface()
			switch key {
			case "profile-name":
				profile.Name = val.(string)
			case "description":
				profile.Description = val.(string)
			// Default Value for this is FALSE
			// If found in RESP then the value is TRUE
			case "data-encryption-flag":
				profile.DataEncryptFlag = val.(bool)
			case "ap-packet-capture-profile":
				profile.ApPktCapProfile = val.(string)
			case "ap-trace-profile":
				profile.ApTraceProfile = val.(string)
			case "tcp-adjust-mss":
				tam := val.(map[string]interface{})
				for tamKey, tamVal := range tam {
					switch tamKey {
					case "adjust-mss":
						profile.TCPMSS.Adjust = tamVal.(bool)
					case "tcp-adjust-mss":
						profile.TCPMSS.AdjustMSS = tamVal.(uint16)
					}
				}
			case "led-state":
				ls := val.(map[string]interface{})
				for lsKey, lsVal := range ls {
					if lsKey == "led-state" {
						profile.LEDState.State = lsVal.(bool)
					}
				}
			case "link-latency":
				ll := val.(map[string]interface{})
				for llKey, llVal := range ll {
					if llKey == "link-latency-flag" {
						profile.LinkLatency.Flag = llVal.(string)
					}
				}
			case "device-mgmt":
				dm := val.(map[string]interface{})
				for dmKey, dmVal := range dm {
					switch dmKey {
					case "ssh":
						profile.DeviceMgmt.SSH = dmVal.(bool)
					case "telnet":
						profile.DeviceMgmt.Telnet = dmVal.(bool)
					}
				}
			case "user-mgmt":
				um := val.(map[string]interface{})
				var keys []string
				for umKey, umVal := range um {
					keys = append(keys, umKey)
					val := umVal.(string)
					switch umKey {
					case "username":
						profile.UserMgmt.Username = val
					case "password":
						profile.UserMgmt.Password = val
					case "password-type":
						profile.UserMgmt.PasswordType = umVal.(CrypType)
					case "secret":
						profile.UserMgmt.Secret = val
					case "secret-type":
						profile.UserMgmt.SecretType = umVal.(CrypType)
					}
				}
				if !s.jsonContains(keys, "secret-type") {
					profile.UserMgmt.SecretType = "clear"
				}
				if !s.jsonContains(keys, "password-type") {
					profile.UserMgmt.PasswordType = "clear"
				}
			case "tunnel":
				var keys []string
				tun := val.(map[string]interface{})
				for tunKey, tunVal := range tun {
					keys = append(keys, tunKey)
					val := tunVal.(string)
					switch tunKey {
					case "preferred-mode":
						profile.CapWapTunnel.PreferredMode = val
					case "udp-lite":
						profile.CapWapTunnel.UDPLite = val
					}
				}
				if !s.jsonContains(keys, "preferred-mode") {
					profile.CapWapTunnel.PreferredMode = "preferred-mode-unconfig"
				}
				if !s.jsonContains(keys, "udp-lite") {
					profile.CapWapTunnel.UDPLite = "udplite-checksum-disabled"
				}
			case "capwap-timer":
				var keys []string
				cwt := val.(map[string]interface{})
				for cwtKey, cwtVal := range cwt {
					keys = append(keys, cwtKey)
					val := cwtVal.(uint8)
					switch cwtKey {
					case "heart-beat-timeout":
						profile.CapWapTimeouts.HeartBeat = val
					case "discovery-timeout":
						profile.CapWapTimeouts.Discovery = val
					case "fast-heart-beat-timeout":
						profile.CapWapTimeouts.FastHeartBeat = val
					case "primary-discovery-timeout":
						profile.CapWapTimeouts.PrimaryDiscovery = val
					case "primed-join-timeout":
						profile.CapWapTimeouts.PrimedJoin = val
					}
				}
				if !s.jsonContains(keys, "heart-beat-timeout") {
					profile.CapWapTimeouts.HeartBeat = 30
				}
				if !s.jsonContains(keys, "discovery-timeout") {
					profile.CapWapTimeouts.Discovery = 10
				}
				if !s.jsonContains(keys, "primery-discovery-timeout") {
					profile.CapWapTimeouts.PrimaryDiscovery = 120
				}
			case "retransmit-timer":
				var keys []string
				rt := val.(map[string]interface{})
				for rtKey, rtVal := range rt {
					keys = append(keys, rtKey)
					val := rtVal.(uint8)
					if rtKey == "count" {
						profile.RetransmitTimer.Count = val
					}
					if rtKey == "interval" {
						profile.RetransmitTimer.Interval = val
					}
				}
				if !s.jsonContains(keys, "count") {
					profile.RetransmitTimer.Count = 5
				}
				if !s.jsonContains(keys, "interval") {
					profile.RetransmitTimer.Interval = 3
				}
			case "login-credentials":
				dlc := val.(map[string]interface{})
				for dlcKey, dlcVal := range dlc {
					val := dlcVal.(string)
					switch dlcKey {
					case "dot1x-username":
						profile.Dot1xCreds.Username = val
					case "dot1x-password":
						profile.Dot1xCreds.Password = val
					case "dot1x-password-type":
						profile.Dot1xCreds.PasswordType = dlcVal.(CrypType)
					}
				}
			case "dot1x-eap-type-info":
				deti := val.(map[string]interface{})
				for detiKey, detiVal := range deti {
					val := detiVal.(string)
					if detiKey == "dot1x-eap-type" {
						profile.Dot1xEapInfo.Type = val
					}
				}
				if profile.Dot1xEapInfo.Type == "" {
					profile.Dot1xEapInfo.Type = "dot1x-eap-fast"
				}
			case "lsc-ap-auth-type-info":
				lati := val.(map[string]interface{})
				for latiKey, latiVal := range lati {
					val := latiVal.(string)
					if latiKey == "lsc-ap-auth-type" {
						profile.LSCApAuthInfo.Type = val
					}
				}
				if profile.LSCApAuthInfo.Type == "" {
					profile.LSCApAuthInfo.Type = "lsc-ap-auth-capwap-dtls"
				}
			case "usb-module-status":
				uv := val.(map[string]interface{})
				for uKey, ukVal := range uv {
					switch uKey {
					case "enable":
						profile.UsbStatus.Enabled = ukVal.(bool)
					}
				}
			// Default Value: False
			case "jumbo-mtu":
				jm := val.(map[string]interface{})
				for _, jmVal := range jm {
					val := jmVal.(bool)
					profile.JumboMTU.MTU = val
				}
			case "coredump":
				cd := val.(map[string]interface{})
				for cdKey, cdVal := range cd {
					switch cdKey {
					case "coredump-flag":
						profile.CoreDump.Flag = cdVal.(MemoryCoreDumpProp)
					case "tftp-server-address":
						profile.CoreDump.TFTPServerAddr = cdVal.(string)
					case "corefile-name":
						profile.CoreDump.FileName = cdVal.(string)
					}
				}
			case "syslog":
				var keys []string
				sl := val.(map[string]interface{})
				for slKey, slVal := range sl {
					keys = append(keys, slKey)
					switch slKey {
					case "facility-value":
						profile.Syslog.FacValue = slVal.(SyslogFacProp)
					case "log-level":
						profile.Syslog.LogLevel = slVal.(SyslogLevelProp)
					case "host":
						profile.Syslog.Host = slVal.(string)
					case "tls-mode":
						profile.Syslog.TLSMode = slVal.(bool)
					}
				}
				if !s.jsonContains(keys, "facility-value") {
					profile.Syslog.FacValue = SyslogFacility.Kern
				}
				if !s.jsonContains(keys, "log-level") {
					profile.Syslog.LogLevel = SyslogLevel.Info
				}
				if !s.jsonContains(keys, "host") {
					profile.Syslog.Host = "255.255.255.255"
				}
			case "backup-controller":
				bc := val.(map[string]interface{})
				for bcKey, bcVal := range bc {
					switch bcKey {
					case "fallback-enabled":
						profile.BackupControllers.Enabled = bcVal.(bool)
					case "primary-controller-name":
						profile.BackupControllers.PrimaryName = bcVal.(string)
					case "secondary-controller-name":
						profile.BackupControllers.SecondaryName = bcVal.(string)
					case "primary-controller-ip":
						profile.BackupControllers.PrimaryIPAddr = bcVal.(string)
					case "secondary-controller-ip":
						profile.BackupControllers.SecondaryIPAddr = bcVal.(string)
					}
				}
			case "hyperlocation":
			case "rogue-detection":
			case "lag-info":
			case "tftp-down-grade":
			case "traffic-limit":
			case "cdp":
				c := val.(map[string]interface{})
				for cKey, cVal := range c {
					if cKey == "cdp-enable" {
						profile.Cdp.Enabled = cVal.(bool)
					}
				}
			case "capwap-window":
			case "mesh":
			case "ntp-server":
				profile.NTPServer = val.(string)
			case "reporting-interval":
			case "ext-module":
			case "icap":
			case "lawful-interception":
			case "persistent-ssid-broadcast":
			case "dhcp-server":
			case "halo-ble-entries":
			case "icap-full-packet-trace-client-mac-address-entries":
			case "icap-partial-packet-trace-client-mac-address-entries":
			case "icap-anomaly-detection-client-mac-address-entries":
			case "icap-client-statistics-client-mac-address-entries":
			case "gas-rate-limit":
			case "qosmap":
			case "qosmap-dscp-to-ups":
			case "qosmap-dscp-to-up-exceptions":
			case "client-rssi":
			case "ntp-server-info":
			case "awips-enabled":
			case "accounting":
			case "apphost":
			case "aux-client-interface":
			case "proxy":
			case "grpc-enable":
			case "bssid-enable-stats":
			case "bssid-stats-frequency":
			case "priviate-ip-discovery":
			case "public-ip-discovery":
			case "led-flash":
			case "bssid-neighbor-stats-enable":
			case "bssid-neighbor-stats-frequency":
			case "traffic-distribution":
			case "dhcp-fallback":
			}
		}
		// Fill Default Values for JSON Keys
		if !s.jsonContains(keys, "backup-controllers") {
			profile.BackupControllers.Enabled = true
		}
		if !s.jsonContains(keys, "ntp-server") {
			profile.NTPServer = "0.0.0.0"
		}
		if !s.jsonContains(keys, "cdp") {
			profile.Cdp.Enabled = true
		}
		if !s.jsonContains(keys, "syslog") {
			profile.Syslog.FacValue = SyslogFacility.Kern
			profile.Syslog.LogLevel = SyslogLevel.Info
			profile.Syslog.Host = "255.255.255.255"
		}
		if !s.jsonContains(keys, "coredump") {
			profile.CoreDump.Flag = MemCoreDump.Disable
		}
		if !s.jsonContains(keys, "tcp-adjust-mss") {
			profile.TCPMSS.Adjust = true
			profile.TCPMSS.AdjustMSS = 1250
		}
		if !s.jsonContains(keys, "led-state") {
			profile.LEDState.State = true
		}
		if !s.jsonContains(keys, "tunnel") {
			profile.CapWapTunnel.PreferredMode = "preferred-mode-unconfig"
			profile.CapWapTunnel.UDPLite = "udplite-checksum-disabled"
		}
		if !s.jsonContains(keys, "capwap-timer") {
			profile.CapWapTimeouts.HeartBeat = 30
			profile.CapWapTimeouts.PrimaryDiscovery = 120
			profile.CapWapTimeouts.Discovery = 10
		}
		if !s.jsonContains(keys, "retransmit-timer") {
			profile.RetransmitTimer.Count = 5
			profile.RetransmitTimer.Interval = 3
		}
		if !s.jsonContains(keys, "login-credentials") {
			profile.Dot1xCreds.PasswordType = CryptTypes.Clear
		}
		if !s.jsonContains(keys, "dot1x-eap-type-info") {
			profile.Dot1xEapInfo.Type = "dot1x-eap-fast"
		}
		if !s.jsonContains(keys, "lsc-ap-auth-type-info") {
			profile.LSCApAuthInfo.Type = "lsc-ap-auth-capwap-dtls"
		}
		if !s.jsonContains(keys, "link-latency") {
			profile.LinkLatency.Flag = "link-auditing-disable"
		}
		if !s.jsonContains(keys, "usb-module-status") {
			profile.UsbStatus.Enabled = true
		}
	}
	return profile
}
