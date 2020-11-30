package sites

// ApJoinProfileSummary ...
type ApJoinProfileSummary struct {
	Name        string `json:"profile-name"`
	Description string `json:"description,omitempty"`
}

// ApJoinProfile ...
type ApJoinProfile struct {
	Name        string `json:"profile-name"`
	Description string `json:"description,omitempty"`
	// Default = 1
	BLEBeaconInterval int `json:"ble-beacon-interval"`
	// Default = 59
	BLEBeaconPwr int `json:"ble-beacon-advpwr"`
	// Default: FALSE
	DataEncryptFlag bool   `json:"data-encryption-flag"`
	ApPktCapProfile string `json:"ap-packet-capture-profile,omitempty"`
	ApTraceProfile  string `json:"ap-trace-profile,omitempty"`
	// DEFAULT: 180
	StatsTimer struct {
		Timer int32 `json:"stats-timer"`
	} `json:"stats-timer"`
	// DEFAULT: TRUE
	LEDState struct {
		State bool `json:"led-state"`
	} `json:"led-state"`
	// DEFAULT: link-auditing-disable
	LinkLatency struct {
		Flag string `json:"link-latency-flag"`
	} `json:"link-latency"`
	// DEFAULT: FALSE
	JumboMTU struct {
		MTU bool `json:"jumbo-mtu"`
	} `json:"jumbo-mtu"`
	UsbStatus  ApUsbModuleStatus `json:"usb-module-status,omitempty"`
	DeviceMgmt DeviceMgmtParams  `json:"device-mgmt,omitempty"`
	UserMgmt   UserMgmtParams    `json:"user-mgmt,omitempty"`
}

// ApUsbModuleStatus ...
type ApUsbModuleStatus struct {
	Enabled bool `json:"enable"`
}

// DeviceMgmtParams ...
type DeviceMgmtParams struct {
	SSH bool `json:"ssh"`
}

// UserMgmtParams ...
type UserMgmtParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}
