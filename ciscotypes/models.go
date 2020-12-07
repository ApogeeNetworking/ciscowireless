package ciscotypes

// Ap is a Cisco Access Point general properties
type Ap struct {
	Name    string
	MacAddr string
	IPAddr  string
	Serial  string
	Model   string
	Tag     ApTag
}

// ApTag ...
type ApTag struct {
	Policy string
	Site   string
	Rf     string
}

// ApTagCfg ...
type ApTagCfg struct {
	MacAddr string `json:"ap-mac"`
	Policy  string `json:"policy-tag,omitempty"`
	Site    string `json:"site-tag,omitempty"`
	Rf      string `json:"rf-tag,omitempty"`
}

// ApCdp ...
type ApCdp struct {
	LocalIntf      string
	RemoteSw       string
	RemoteIntf     string
	RemoteSwIPAddr string
}

// ApEthIntf ...
type ApEthIntf struct {
	Name   string
	Status string
	Speed  string
	TxRcv  string
	Drops  string
}

// ApLanPort ...
type ApLanPort struct {
	ID     int
	Status LanPortState
}

// LanPortState string enable|disable
type LanPortState string

// LPState {Enable: LanPortState, Disable: LanPortState}
type LPState struct {
	Enable  LanPortState
	Disable LanPortState
}

// ApLanPortState enum for ENABLE|DISABLE (do not offer other options)
var ApLanPortState = LPState{
	Enable:  "enable",
	Disable: "disable",
}
