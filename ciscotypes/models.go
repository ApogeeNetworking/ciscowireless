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
