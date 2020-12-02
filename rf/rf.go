package rf

import (
	"encoding/json"
	"reflect"

	"github.com/ApogeeNetworking/ciscowireless/requests"
)

// Service ...
type Service struct {
	http         *requests.Service
	jsonContains func(k []string, v string) bool
}

// NewService ...
func NewService(r *requests.Service, c func(k []string, v string) bool) *Service {
	return &Service{http: r, jsonContains: c}
}

// GetProfiles ...
func (s *Service) GetProfiles() ([]Profile, error) {
	uri := "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-profiles/rf-profile"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var profilesMap map[string][]interface{}
	json.NewDecoder(res.Body).Decode(&profilesMap)
	profiles := profilesMap["Cisco-IOS-XE-wireless-rf-cfg:rf-profile"]
	var rfProfiles []Profile
	for _, profileMap := range profiles {
		v := reflect.ValueOf(profileMap)
		if v.Kind() != reflect.Map {
			continue
		}
		rfProfiles = append(rfProfiles, s.handleRfProfiles(v))
	}
	return rfProfiles, nil
}

func (s *Service) handleRfProfiles(v reflect.Value) Profile {
	var profile Profile
	var keys []string
	for _, k := range v.MapKeys() {
		key := k.String()
		keys = append(keys, key)
		val := v.MapIndex(k).Interface()
		switch key {
		case "name":
			profile.Name = val.(string)
		case "description":
			profile.Description = val.(string)
		case "tx-power-min":
			profile.MinTxPwr = val.(float64)
		case "tx-power-max":
			profile.MaxTxPwr = val.(float64)
		case "tx-power-v1-threshold":
			profile.TxPwrV1Threshold = val.(float64)
		case "tx-power-v2-threshold":
			profile.TxPwrV2Threshold = val.(int16)
		case "status":
			profile.Status = val.(bool)
		case "band":
			profile.Band = Dot11RadioBandProp(val.(string))
		case "data-rate-1m":
			profile.DataRate1m = DataRateStateProp(val.(string))
		case "data-rate-2m":
			profile.DataRate2m = DataRateStateProp(val.(string))
		case "data-rate-5-5m":
			profile.DataRate5m = DataRateStateProp(val.(string))
		case "data-rate-11m":
			profile.DataRate11m = DataRateStateProp(val.(string))
		case "data-rate-6m":
			profile.DataRate6m = DataRateStateProp(val.(string))
		case "data-rate-9m":
			profile.DataRate9m = DataRateStateProp(val.(string))
		case "data-rate-12m":
			profile.DataRate12m = DataRateStateProp(val.(string))
		case "data-rate-18m":
			profile.DataRate18m = DataRateStateProp(val.(string))
		case "data-rate-24m":
			profile.DataRate24m = DataRateStateProp(val.(string))
		case "data-rate-36m":
			profile.DataRate36m = DataRateStateProp(val.(string))
		case "data-rate-48m":
			profile.DataRate48m = DataRateStateProp(val.(string))
		case "data-rate-54m":
			profile.DataRate54m = DataRateStateProp(val.(string))
		case "coverage-data-packet-rssi-threshold":
			profile.CoverageDataPacketRSSIThreshold = val.(float64)
		case "coverage-voice-packet-rssi-threshold":
			profile.CoverageVoicePacketRSSIThreshold = val.(float64)
		case "min-num-clients":
			profile.MinNumClients = val.(float64)
		case "max-radio-clients":
			profile.MaxNumClients = val.(uint16)
		case "exception-level":
			profile.ExceptionLevel = val.(uint16)
		case "band-select-client-rssi":
			profile.BandSelectClientRSSI = val.(int16)
		case "band-select-client-mid-rssi":
			profile.BandSelectClientMidRSSI = val.(int16)
		case "band-select-cycle-count":
			profile.BandSelectCycleCount = val.(uint16)
		case "band-select-cycle-threshold":
			profile.BandSelectCycleThreshold = val.(uint16)
		case "band-select-age-out-dual-band":
			profile.BandSelectAgeOutDualBand = val.(uint16)
		case "band-select-age-out-suppression":
			profile.BandSelectAgeOutSuppression = val.(uint16)
		case "band-select-probe-response":
			profile.BandSelectProbeResponse = val.(bool)
		case "dca-contribution-interference":
			profile.DCAContribInterference = val.(bool)
		case "rf-dca-chan-width":
			profile.RfDcaChanWidth = DynChannelWidthProp(val.(string))
		case "load-balancing-window":
			profile.LoadBalancingWindow = val.(uint32)
		case "load-balancing-denial-count":
			profile.LoadBalancingDenialCount = val.(uint32)
		case "trap-threshold-clients":
			profile.TrapThresholdClients = val.(uint32)
		case "trap-threshold-interference":
			profile.TrapThresholdInterference = val.(uint16)
		case "trap-threshold-noise":
			profile.TrapThresholdNoise = val.(int32)
		case "trap-threshold-utilization":
			profile.TrapThresholdUtilization = val.(uint16)
		case "multicast-data-rate":
			profile.MCastDataRate = MCastDataRateProp(val.(string))
		case "rx-sen-sop-threshold":
			profile.RxSenSOPThreshold = RxSenSOPThresholdProp(val.(string))
		case "rx-sen-sop-custom":
			profile.RxSenSOPCustom = val.(int16)
		case "client-network-preference":
			profile.ClientNetworkPref = ClientNetworkPrefProp(val.(string))
		case "hsr-mode":
			profile.HSRMode = val.(bool)
		case "hsr-neighbor-timeout":
			profile.HSRNeighborTimeout = val.(uint16)
		case "opt-roam-rssi-threshold":
			profile.OptRoamRSSIThreshold = val.(int16)
		case "opt-roam-rssi-check-enable":
			profile.OptRoamRSSICheckEnabled = val.(bool)
		case "atf-oper-mode":
			profile.ATFOperMode = AtfProp(val.(string))
		case "atf-optimization":
			profile.ATFOptimization = AtfOptiProp(val.(string))
		case "bridge-client-access":
			profile.BridgeClientAccess = val.(bool)
		case "airtime-allocation":
			profile.AirtimeAllocation = val.(uint8)
		case "client-aware-fra":
			profile.ClientAwareFRA = val.(bool)
		case "client-select-threshold":
			profile.ClientSelectThreshold = val.(uint8)
		case "client-reset-threshold":
			profile.ClientResetThreshold = val.(uint8)
		case "rf-mcs-entries":
			var mcsEntries []McsEntry
			rme := val.(map[string]interface{})
			for _, rmeVal := range rme {
				rmeMaps := rmeVal.([]interface{})
				for _, mapVal := range rmeMaps {
					var mcsEntry McsEntry
					entryMap := mapVal.(map[string]interface{})
					var keys []string
					for key, val := range entryMap {
						keys = append(keys, key)
						switch key {
						case "rf-index":
							mcsEntry.ID = val.(float64)
						case "rf-80211n-mcs-enable":
							mcsEntry.Dot11nMcsEnable = val.(bool)
						}
					}
					if !s.jsonContains(keys, "rf-80211n-mcs-enable") {
						mcsEntry.Dot11nMcsEnable = true
					}
					mcsEntries = append(mcsEntries, mcsEntry)
				}
			}
			profile.MCSEntries.Entries = mcsEntries
		case "rf-dca-allowed-channels":
			var allowedChannels []DcaChannel
			rdac := val.(map[string]interface{})
			for _, rdacVal := range rdac {
				rdacMaps := rdacVal.([]interface{})
				for _, mapVal := range rdacMaps {
					var entry DcaChannel
					entryMap := mapVal.(map[string]interface{})
					for _, val := range entryMap {
						entry.Channel = val.(float64)
					}
					allowedChannels = append(allowedChannels, entry)
				}
			}
			profile.DcaAllowedChannels.Channels = allowedChannels
		case "rfdca-removed-channels":
			var removedChannels []DcaChannel
			rdac := val.(map[string]interface{})
			for _, rdacVal := range rdac {
				rdacMaps := rdacVal.([]interface{})
				for _, mapVal := range rdacMaps {
					var entry DcaChannel
					entryMap := mapVal.(map[string]interface{})
					for _, val := range entryMap {
						entry.Channel = val.(float64)
					}
					removedChannels = append(removedChannels, entry)
				}
			}
			profile.DcaRemovedChannels.Channels = removedChannels
		}
	}
	s.handleDefaultProfileValues(&profile, keys)
	return profile
}

func (s *Service) handleDefaultProfileValues(profile *Profile, keys []string) {
	if !s.jsonContains(keys, "tx-power-min") {
		profile.MinTxPwr = -10
	}
	if !s.jsonContains(keys, "tx-power-max") {
		profile.MaxTxPwr = 30
	}
	if !s.jsonContains(keys, "tx-power-v1-threshold") {
		profile.TxPwrV1Threshold = -70
	}
	if !s.jsonContains(keys, "tx-power-v2-threshold") {
		profile.TxPwrV2Threshold = -67
	}
	if !s.jsonContains(keys, "coverage-data-packet-rssi-threshold") {
		profile.CoverageDataPacketRSSIThreshold = -80
	}
	if !s.jsonContains(keys, "coverage-voice-packet-rssi-threshold") {
		profile.CoverageVoicePacketRSSIThreshold = -80
	}
	if !s.jsonContains(keys, "min-num-clients") {
		profile.MinNumClients = 3
	}
	if !s.jsonContains(keys, "max-radio-clients") {
		profile.MaxNumClients = 200
	}
	if !s.jsonContains(keys, "exception-level") {
		profile.ExceptionLevel = 25
	}
	if !s.jsonContains(keys, "band-select-client-rssi") {
		profile.BandSelectClientRSSI = -80
	}
	if !s.jsonContains(keys, "band-select-client-mid-rssi") {
		profile.BandSelectClientMidRSSI = -80
	}
	if !s.jsonContains(keys, "band-select-cycle-count") {
		profile.BandSelectCycleCount = 2
	}
	if !s.jsonContains(keys, "band-select-cycle-threshold") {
		profile.BandSelectCycleThreshold = 200
	}
	if !s.jsonContains(keys, "band-select-age-out-dual-band") {
		profile.BandSelectAgeOutDualBand = 60
	}
	if !s.jsonContains(keys, "band-select-age-out-suppression") {
		profile.BandSelectAgeOutSuppression = 20
	}
	if !s.jsonContains(keys, "dca-contribution-interference") {
		profile.DCAContribInterference = true
	}
	if !s.jsonContains(keys, "rf-dca-chan-width") {
		profile.RfDcaChanWidth = DynChannelWidths.RfBest
	}
	if !s.jsonContains(keys, "load-balancing-window") {
		profile.LoadBalancingWindow = 5
	}
	if !s.jsonContains(keys, "load-balancing-denial-count") {
		profile.LoadBalancingDenialCount = 3
	}
	if !s.jsonContains(keys, "trap-threshold-clients") {
		profile.TrapThresholdClients = 12
	}
	if !s.jsonContains(keys, "trap-threshold-interference") {
		profile.TrapThresholdInterference = 10
	}
	if !s.jsonContains(keys, "trap-threshold-noise") {
		profile.TrapThresholdNoise = -70
	}
	if !s.jsonContains(keys, "trap-threshold-utilization") {
		profile.TrapThresholdUtilization = 80
	}
	if !s.jsonContains(keys, "multicast-data-rate") {
		profile.MCastDataRate = MCastDataRates.Default
	}
	if !s.jsonContains(keys, "rx-sen-sop-threshold") {
		profile.RxSenSOPThreshold = RxSenSOPThresholds.Auto
	}
	if !s.jsonContains(keys, "client-network-preference") {
		profile.ClientNetworkPref = ClientNetworkPreferences.Default
	}
	if !s.jsonContains(keys, "hsr-neighbor-timeout") {
		profile.HSRNeighborTimeout = 5
	}
	if !s.jsonContains(keys, "opt-roam-rssi-threshold") {
		profile.OptRoamRSSIThreshold = -127
	}
	if !s.jsonContains(keys, "atf-oper-mode") {
		profile.ATFOperMode = AirtimeFairnessMode.Disable
	}
	if !s.jsonContains(keys, "atf-optimization") {
		profile.ATFOptimization = AtfOptimizationState.Disabled
	}
	if !s.jsonContains(keys, "client-select-threshold") {
		profile.ClientSelectThreshold = 50
	}
	if !s.jsonContains(keys, "client-reset-threshold") {
		profile.ClientResetThreshold = 5
	}
	if !s.jsonContains(keys, "data-rate-1m") {
		profile.DataRate1m = "apf-tx-rate-basic"
	}
	if !s.jsonContains(keys, "data-rate-2m") {
		profile.DataRate2m = "apf-tx-rate-basic"
	}
	if !s.jsonContains(keys, "data-rate-5-5m") {
		profile.DataRate5m = "apf-tx-rate-basic"
	}
	if !s.jsonContains(keys, "data-rate-11m") {
		profile.DataRate11m = "apf-tx-rate-basic"
	}
	if !s.jsonContains(keys, "data-rate-6m") {
		profile.DataRate6m = "apf-tx-rate-not-applicable"
	}
	if !s.jsonContains(keys, "data-rate-9m") {
		profile.DataRate9m = "apf-tx-rate-supported"
	}
	if !s.jsonContains(keys, "data-rate-12m") {
		profile.DataRate12m = "apf-tx-rate-not-applicable"
	}
	if !s.jsonContains(keys, "data-rate-18m") {
		profile.DataRate18m = "apf-tx-rate-supported"
	}
	if !s.jsonContains(keys, "data-rate-24m") {
		profile.DataRate24m = "apf-tx-rate-not-applicable"
	}
	if !s.jsonContains(keys, "data-rate-36m") {
		profile.DataRate36m = "apf-tx-rate-supported"
	}
	if !s.jsonContains(keys, "data-rate-48m") {
		profile.DataRate48m = "apf-tx-rate-supported"
	}
	if !s.jsonContains(keys, "data-rate-54m") {
		profile.DataRate54m = "apf-tx-rate-supported"
	}
}

// GetTags ...
func (s *Service) GetTags() ([]Tag, error) {
	req, err := s.http.GenerateRequest(
		"/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag",
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
	tags := tagsMap["Cisco-IOS-XE-wireless-rf-cfg:rf-tag"]
	var rfTags []Tag
	for _, tagMap := range tags {
		v := reflect.ValueOf(tagMap)
		if v.Kind() != reflect.Map {
			continue
		}
		var rfTag Tag
		var keys []string
		for _, k := range v.MapKeys() {
			key := k.String()
			keys = append(keys, key)
			val := v.MapIndex(k).Interface().(string)
			switch key {
			case "tag-name":
				rfTag.Name = val
			case "description":
				rfTag.Description = val
			case "dot11a-rf-profile-name":
				rfTag.Dot11aProfileName = val
			case "dot11b-rf-profile-name":
				rfTag.Dot11bProfileName = val
			}
		}
		if !s.jsonContains(keys, "dot11a-rf-profile-name") {
			rfTag.Dot11aProfileName = "default-rf-5gh"
		}
		if !s.jsonContains(keys, "dot11b-rf-profile-name") {
			rfTag.Dot11bProfileName = "default-rf-24gh"
		}
		rfTags = append(rfTags, rfTag)
	}
	return rfTags, nil
}
