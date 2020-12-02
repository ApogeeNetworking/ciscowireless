package rf

// Tag ...
type Tag struct {
	// Tag Name: Once Created Cannot be Modified (only Deleted)
	Name string `json:"tag-name"`
	// Optional Field
	Description string `json:"description"`
	// Default: "default-rf-5gh"
	Dot11aProfileName string `json:"dot11a-rf-profile-name"`
	// Default: "default-rf-24gh"
	Dot11bProfileName string `json:"dot11b-rf-profile-name"`
}

// Profile ...
type Profile struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// Default: -10
	MinTxPwr float64 `json:"tx-power-min"`
	// Default: 30
	MaxTxPwr float64 `json:"tx-power-max"`
	// Default: -70
	TxPwrV1Threshold float64 `json:"tx-power-v1-threshold"`
	// Default: -67
	TxPwrV2Threshold int16 `json:"tx-power-v2-threshold"`
	// Default: false
	Status bool `json:"status"`
	// ENUM
	Band Dot11RadioBandProp `json:"band"`
	// Default: "apf-tx-rate-basic"
	DataRate1m DataRateStateProp `json:"data-rate-1m"`
	// Default: "apf-tx-rate-basic"
	DataRate2m DataRateStateProp `json:"data-rate-2m"`
	// Default: "apf-tx-rate-basic"
	DataRate5m DataRateStateProp `json:"data-rate-5-5m"`
	// Default: "apf-tx-rate-basic"
	DataRate11m DataRateStateProp `json:"data-rate-11m"`
	// Default: "apf-tx-rate-not-applicable"
	DataRate6m DataRateStateProp `json:"data-rate-6m"`
	// Default: "apf-tx-rate-supported"
	DataRate9m DataRateStateProp `json:"data-rate-9m"`
	// Default: "apf-tx-rate-not-applicable"
	DataRate12m DataRateStateProp `json:"data-rate-12m"`
	// Default: "apf-tx-rate-supported"
	DataRate18m DataRateStateProp `json:"data-rate-18m"`
	// Default: "apf-tx-rate-not-applicable"
	DataRate24m DataRateStateProp `json:"data-rate-24m"`
	// Default: "apf-tx-rate-supported"
	DataRate36m DataRateStateProp `json:"data-rate-36m"`
	// Default: "apf-tx-rate-supported"
	DataRate48m DataRateStateProp `json:"data-rate-48m"`
	// Default: "apf-tx-rate-supported"
	DataRate54m DataRateStateProp `json:"data-rate-54m"`
	// Default: -80
	CoverageDataPacketRSSIThreshold float64 `json:"coverage-data-packet-rssi-threshold"`
	// Default: -80
	CoverageVoicePacketRSSIThreshold float64 `json:"coverage-voice-packet-rssi-threshold"`
	// Default: 3
	MinNumClients float64 `json:"min-num-clients"`
	// Default: 200
	MaxNumClients uint16 `json:"max-radio-clients"`
	// Default: 25
	ExceptionLevel uint16 `json:"exception-level"`
	// Default: -80
	BandSelectClientRSSI int16 `json:"band-select-client-rssi"`
	// Default: -80
	BandSelectClientMidRSSI int16 `json:"band-select-client-mid-rssi"`
	// Default: 2
	BandSelectCycleCount uint16 `json:"band-select-cycle-count"`
	// Default: 200
	BandSelectCycleThreshold uint16 `json:"band-select-cycle-threshold"`
	// Default: 60
	BandSelectAgeOutDualBand uint16 `json:"band-select-age-out-dual-band"`
	// Default: 20
	BandSelectAgeOutSuppression uint16 `json:"band-select-age-out-suppression"`
	// Default: false
	BandSelectProbeResponse bool `json:"band-select-probe-response"`
	// Default: true
	DCAContribInterference bool `json:"dca-contribution-interference"`
	// Default: rf-dca-chan-width-best
	RfDcaChanWidth DynChannelWidthProp `json:"rf-dca-chan-width"`
	// Default: 5
	LoadBalancingWindow uint32 `json:"load-balancing-window"`
	// Default: 3
	LoadBalancingDenialCount uint32 `json:"load-balancing-denial-count"`
	// Default: 12
	TrapThresholdClients uint32 `json:"trap-threshold-clients"`
	// Default: 10
	TrapThresholdInterference uint16 `json:"trap-threshold-interference"`
	// Default: -70
	TrapThresholdNoise int32 `json:"trap-threshold-noise"`
	// Default: 80
	TrapThresholdUtilization uint16 `json:"trap-threshold-utilization"`
	// Default: "mcast-data-rate-default"
	MCastDataRate MCastDataRateProp `json:"multicast-data-rate"`
	// Default: "rrm-ewlc-rxsensop-threshold-auto"
	RxSenSOPThreshold RxSenSOPThresholdProp `json:"rx-sen-sop-threshold"`
	// Default: 0 (-85 & -60)
	RxSenSOPCustom int16 `json:"rx-sen-sop-custom,omitempty"`
	// Default: default
	ClientNetworkPref ClientNetworkPrefProp `json:"client-network-preference"`
	// High Speed Roam Mode (HSR)
	// Default: false
	HSRMode bool `json:"hsr-mode"`
	// Default: 5
	HSRNeighborTimeout uint16 `json:"hsr-neighbor-timeout"`
	// Default: -127 (Yang Prop|Field Bug: "opt-roam-rssi-treshold")
	OptRoamRSSIThreshold int16 `json:"opt-roam-rssi-threshold"`
	// Default: false
	OptRoamRSSICheckEnabled bool `json:"opt-roam-rssi-check-enable"`
	// Air Time Fairness:
	// Default: "apf-atf-mode-disable"
	ATFOperMode AtfProp `json:"atf-oper-mode"`
	// Default: "apf-atf-stealing-policy-disable"
	ATFOptimization AtfOptiProp `json:"atf-optimization"`
	// Default: false
	BridgeClientAccess bool `json:"bridge-client-access"`
	// Default: 5
	AirtimeAllocation uint8 `json:"airtime-allocation"`
	// Client aware FRA allows the Dual band to operate
	// On 5ghz or monitor mode depending on the load
	// On the dedicated 5Ghz Radio
	// Default: false
	ClientAwareFRA bool `json:"client-aware-fra"`
	// Client aware FRA utilization threshold for moving Dual
	// Band radio from Monitor Mode to Client Saving
	// Default: 50
	ClientSelectThreshold uint8 `json:"client-select-threshold"`
	// Client aware FRA utilization value for moving Dual
	// Band Radio from Client Serving Mode to Monitor Mode
	// Default: 5
	ClientResetThreshold uint8 `json:"client-reset-threshold"`
	// Modulation and Coding Scheme (MCS) Data Rate
	MCSEntries         McsItem               `json:"rf-mcs-entries,omitempty"`
	DcaAllowedChannels DcaAllowedChannelItem `json:"rf-dca-allowed-channels,omitempty"`
	DcaRemovedChannels DcaRemovedChannelItem `json:"rfdca-removed-channels,omitempty"`
}

// DcaAllowedChannelItem ...
type DcaAllowedChannelItem struct {
	Channels []DcaChannel `json:"rf-dca-allowed-channel,omitempty"`
}

// DcaRemovedChannelItem ...
type DcaRemovedChannelItem struct {
	Channels []DcaChannel `json:"rfdca-removed-channel,omitempty"`
}

// DcaChannel ...
type DcaChannel struct {
	Channel float64 `json:"channel"`
}

// McsItem ...
type McsItem struct {
	Entries []McsEntry `json:"rf-mcs-entry"`
}

// McsEntry ...
type McsEntry struct {
	ID float64 `json:"rf-index"`
	// Default: true
	Dot11nMcsEnable bool `json:"rf-80211n-mcs-enable"`
}

// Dot11RadioBandProp ...
type Dot11RadioBandProp string

// Dot11RadioBand ...
type Dot11RadioBand struct {
	TwoFourGhz Dot11RadioBandProp
	FiveGhz    Dot11RadioBandProp
}

// DataRateStateProp ...
type DataRateStateProp string

// DataRateState ...
type DataRateState struct {
	Basic         DataRateStateProp
	Supported     DataRateStateProp
	Unsupported   DataRateStateProp
	NotApplicable DataRateStateProp
}

// DynChannelWidthProp ...
type DynChannelWidthProp string

// DynChanWidth ...
type DynChanWidth struct {
	Rf20Mhz   DynChannelWidthProp
	Rf40Mhz   DynChannelWidthProp
	Rf80Mhz   DynChannelWidthProp
	RfBest    DynChannelWidthProp
	Rf160Mhz  DynChannelWidthProp
	Rf8080Mhz DynChannelWidthProp
}

// MCastDataRateProp ...
type MCastDataRateProp string

// MCastDataRate ...
type MCastDataRate struct {
	Default      MCastDataRateProp
	SixM         MCastDataRateProp
	NineM        MCastDataRateProp
	TwelveM      MCastDataRateProp
	EighteenM    MCastDataRateProp
	TwentyFourM  MCastDataRateProp
	ThirtySixM   MCastDataRateProp
	FourtyEightM MCastDataRateProp
	FiftyFourM   MCastDataRateProp
}

// RxSenSOPThresholdProp ...
type RxSenSOPThresholdProp string

// RxSenSOPThreshold ...
type RxSenSOPThreshold struct {
	Auto   RxSenSOPThresholdProp
	Low    RxSenSOPThresholdProp
	Medium RxSenSOPThresholdProp
	High   RxSenSOPThresholdProp
	Custom RxSenSOPThresholdProp
}

// ClientNetworkPrefProp ...
type ClientNetworkPrefProp string

// ClientNetworkPreference ...
type ClientNetworkPreference struct {
	Default      ClientNetworkPrefProp
	Connectivity ClientNetworkPrefProp
	Throughput   ClientNetworkPrefProp
}

// AtfProp ...
type AtfProp string

// AirtimeFairnessModes ...
type AirtimeFairnessModes struct {
	Disable AtfProp
	SSID    AtfProp
	Monitor AtfProp
}

// AtfOptiProp ...
type AtfOptiProp string

// AtfOptimizationStates ...
type AtfOptimizationStates struct {
	Enabled  AtfOptiProp
	Disabled AtfOptiProp
}

var (
	// AtfOptimizationState ...
	AtfOptimizationState = AtfOptimizationStates{
		Enabled:  "apf-atf-stealing-policy-enable",
		Disabled: "apf-atf-stealing-policy-disable",
	}
	// AirtimeFairnessMode ...
	AirtimeFairnessMode = AirtimeFairnessModes{
		Disable: "apf-atf-mode-disable",
		SSID:    "apf-atf-mode-ssid",
		Monitor: "apf-atf-mode-monitor",
	}
	// ClientNetworkPreferences ...
	ClientNetworkPreferences = ClientNetworkPreference{
		Default:      "default",
		Connectivity: "connectivity",
		Throughput:   "throughput",
	}
	// RxSenSOPThresholds ...
	RxSenSOPThresholds = RxSenSOPThreshold{
		Auto:   "rrm-ewlc-rxsensop-threshold-auto",
		Low:    "rrm-ewlc-rxsensop-threshold-low",
		Medium: "rrm-ewlc-rxsensop-threshold-medium",
		High:   "rrm-ewlc-rxsensop-threshold-high",
		Custom: "rrm-ewlc-rxsensop-threshold-custom",
	}
	// MCastDataRates ...
	MCastDataRates = MCastDataRate{
		Default:      "mcast-data-rate-default",
		SixM:         "mcast-data-rate-6m",
		NineM:        "mcast-data-rate-9m",
		TwelveM:      "mcast-data-rate-12m",
		TwentyFourM:  "mcast-data-rate-24m",
		ThirtySixM:   "mcast-data-rate-36m",
		FourtyEightM: "mcast-data-rate-48m",
		FiftyFourM:   "mcast-data-rate-54m",
	}
	// DynChannelWidths ...
	DynChannelWidths = DynChanWidth{
		Rf20Mhz:  "rf-dca-chan-width-20-mhz",
		Rf40Mhz:  "rf-dca-chan-width-40-mhz",
		Rf80Mhz:  "rf-dca-chan-width-80-mhz",
		RfBest:   "rf-dca-chan-width-best",
		Rf160Mhz: "rf-dca-chan-width-160-mhz",
		// 80+80 BANDS
		Rf8080Mhz: "rf-dca-chan-width-8080-mhz",
	}
	// DataRateStates ...
	DataRateStates = DataRateState{
		Basic:         "apf-tx-rate-basic",
		Supported:     "apf-tx-rate-supported",
		Unsupported:   "apf-tx-rate-unsupported",
		NotApplicable: "apf-tx-rate-not-applicable",
	}
	// Dot11RadioBands ...
	Dot11RadioBands = Dot11RadioBand{
		TwoFourGhz: "dot11-2-dot4-ghz-band",
		FiveGhz:    "dot11-5-ghz-band",
	}
)
