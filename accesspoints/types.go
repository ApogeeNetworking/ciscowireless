package accesspoints

// Ap ...
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

// ApOptions ...
type ApOptions struct {
	ApName string `json:"ap-name,omitempty"`
	// RADIO MAC ADDRESS (not ETH | WTP MAC)
	MacAddr string `json:"mac-addr,omitempty"`
	// Only Use for SetApName Method
	NewName string `json:"name,omitempty"`
}

// ApNameUpdate ...
type ApNameUpdate struct {
	OldName string `json:"ap-name"`
	// Currently supports Names with Dashes(-) or Without
	// Does NOT support Names with Dots(.)
	NewName string `json:"name"`
}

// ApTagCfg ...
type ApTagCfg struct {
	MacAddr string `json:"ap-mac"`
	Policy  string `json:"policy-tag,omitempty"`
	Site    string `json:"site-tag,omitempty"`
	Rf      string `json:"rf-tag,omitempty"`
}

type capWapResp struct {
	Name    string             `json:"name"`
	IPAddr  string             `json:"ip-addr"`
	Detail  capWapDeviceDetail `json:"device-detail"`
	TagInfo capWapTags         `json:"tag-info"`
}
type capWapDeviceDetail struct {
	Info capWapInfo `json:"static-info"`
}
type capWapInfo struct {
	Board capWapBoard `json:"board-data"`
	Model capWapModel `json:"ap-models"`
}
type capWapBoard struct {
	Serial  string `json:"wtp-serial-num"`
	MacAddr string `json:"wtp-enet-mac"`
}
type capWapModel struct {
	Model string `json:"model"`
}
type capWapTags struct {
	Policy capWapPolicyInfo `json:"policy-tag-info"`
	Site   capWapSiteInfo   `json:"site-tag"`
	Rf     capWapRfInfo     `json:"rf-tag"`
}
type capWapPolicyInfo struct {
	Name string `json:"policy-tag-name"`
}
type capWapSiteInfo struct {
	Name string `json:"site-tag-name"`
}
type capWapRfInfo struct {
	Name string `json:"rf-tag-name"`
}
