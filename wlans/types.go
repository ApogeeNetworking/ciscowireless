package wlans

// WLAN ...
type WLAN struct {
	Name        string `json:"profile-name"`
	Description string `json:"description"`
	ID          int    `json:"wlan-id"`
	// Default: true
	SecurityWPA      bool     `json:"security-wpa"`
	AuthKeyMgmtPSK   bool     `json:"auth-key-mgmt-psk"`
	AuthKeyMgmtDot1x bool     `json:"auth-key-mgmt-dot1x"`
	PSK              string   `json:"psk,omitempty"`
	MacFilterList    string   `json:"mac-filtering-list,omitempty"`
	AuthList         string   `json:"authentication-list"`
	WPA2Enabled      bool     `json:"wpa2-enabled"`
	WPA2AES          bool     `json:"wpa2-aes"`
	RadioPolicy      string   `json:"radio-policy"`
	Info             SSIDInfo `json:"apf-vap-id-data"`
}

// SSIDInfo ...
type SSIDInfo struct {
	Name           string `json:"ssid"`
	Broadcast      bool   `json:"broadcast-ssid"`
	CCXAironetIE   bool   `json:"ccx-aironet-id,omitempty"`
	P2PBlockAction string `json:"p2p-block-action,omitempty"`
	Dot11aDTIM     int    `json:"dot11a-dtim,omitempty"`
	Dot11bDTIM     int    `json:"dot11b-dtim,omitempty"`
	Chd            string `json:"chd,omitempty"`
	Status         bool   `json:"wlan-status"`
}

// Policy ...
type Policy struct {
	Name            string          `json:"policy-profile-name"`
	Description     string          `json:"description"`
	Status          bool            `json:"status"`
	InterfaceName   string          `json:"interface-name"`
	WlanTimeout     Timeout         `json:"wlan-timeout,omitempty"`
	LocalProfiling  LocalProfiling  `json:"wlan-local-profiling"`
	BlackListParams BlackListParams `json:"blacklist-params"`
	AAAPolicyParams AAAPolicyParams `json:"aaa-policy-params"`
	AccountingList  string          `json:"account-list"`
	UPN             UPNProps        `json:"upn,omitempty"`
}

// Timeout ...
type Timeout struct {
	Session       int32 `json:"session-timeout,omitempty"`
	Idle          int32 `json:"idle-timeout,omitempty"`
	IdleThreshold int32 `json:"idle-threshold,omitempty"`
}

// LocalProfiling ...
type LocalProfiling struct {
	DeviceClass          bool   `json:"device-classification,omitempty"`
	SubscriberPolicyName string `json:"subscriber-policy-name,omitempty"`
	Radius               bool   `json:"radius-profiling,omitemtpy"`
	HTTPTlvCaching       bool   `json:"http-tlv-caching,omitemtpy"`
	DHCPTlvCaching       bool   `json:"dhcp-tlv-caching,omitempty"`
}

// BlackListParams ...
type BlackListParams struct {
	IsEnabled bool  `json:"is-blacklist-enabled,omitempty"`
	Timeout   int32 `json:"blacklist-timeout,omitempty"`
}

// AAAPolicyParams ...
type AAAPolicyParams struct {
	Override bool   `json:"aaa-override,omitempty"`
	NAC      bool   `json:"nac,omitempty"`
	NACType  string `json:"nac-type,omitempty"`
}

// UPNProps ...
type UPNProps struct {
	IsRestrictEnabled bool `json:"is-upn-restrict-enabled"`
	UnicastDisable    bool `json:"upn-unicast-disable"`
}

// PolicyTag ...
type PolicyTag struct {
	Name        string                  `json:"tag-name"`
	Policies    ConfigPolicySummary     `json:"wlan-policies"`
	RlanConfigs RlanConfigPolicySummary `json:"tag-child-rlan-policy-configs"`
}

// ConfigPolicySummary ...
type ConfigPolicySummary struct {
	WlanPolicy []WlanConfigPolicy `json:"wlan-policy"`
}

// WlanConfigPolicy ...
type WlanConfigPolicy struct {
	WlanProfileName   string `json:"wlan-profile-name"`
	PolicyProfileName string `json:"policy-profile-name"`
}

// RlanConfigPolicySummary ...
type RlanConfigPolicySummary struct {
	RlanConfig []RlanProfilePolicy `json:"tag-child-rlan-policy-config"`
}

// RlanProfilePolicy ...
type RlanProfilePolicy struct {
	PortID            int    `json:"port-id"`
	ProfileName       string `json:"rlan-profile-name"`
	PolicyProfileName string `json:"rlan-policy-profile-name"`
}
