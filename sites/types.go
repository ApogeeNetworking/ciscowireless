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
		Timer uint32 `json:"stats-timer"`
	} `json:"stats-timer"`
	// TCP Maximum Segement Size
	TCPMSS struct {
		// Default: TRUE
		Adjust bool `json:"adjust-mss"`
		// Default: 1250
		AdjustMSS uint16 `json:"tcp-adjust-mss"`
	} `json:"tcp-mss,omitempty"`
	// DEFAULT: TRUE
	LEDState struct {
		State bool `json:"led-state"`
	} `json:"led-state"`
	LinkLatency struct {
		// DEFAULT: link-auditing-disable
		Flag string `json:"link-latency-flag"`
	} `json:"link-latency"`
	CapWapTunnel struct {
		PreferredMode string `json:"preferred-mode"`
		UDPLite       string `json:"udp-lite"`
	} `json:"tunnel"`
	CapWapTimeouts struct {
		// Default: 30
		HeartBeat uint8 `json:"heart-beat-timeout"`
		// Default: 10
		Discovery uint8 `json:"discovery-timeout"`
		// Default: 0
		FastHeartBeat uint8 `json:"fast-heart-beat-timeout"`
		// Default: 120
		PrimaryDiscovery uint8 `json:"primary-discovery-timeout"`
		// Default: 0
		PrimedJoin uint8 `json:"primed-join-timeout"`
	} `json:"capwap-timer"`
	RetransmitTimer struct {
		// Default: 5
		Count uint8 `json:"count"`
		// Default: 3
		Interval uint8 `json:"interval"`
	} `json:"retransmit-timer"`
	Dot1xCreds struct {
		Username     string   `json:"dot1x-username"`
		Password     string   `json:"dot1x-password"`
		PasswordType CrypType `json:"dot1x-password-type"`
	} `json:"login-credentials"`
	Dot1xEapInfo struct {
		Type string `json:"dot1x-eap-type"`
	} `json:"dot1x-eap-type-info"`
	LSCApAuthInfo struct {
		Type string `json:"lsc-ap-auth-type"`
	} `json:"lsc-ap-auth-type-info"`
	JumboMTU struct {
		// DEFAULT: FALSE
		MTU bool `json:"jumbo-mtu"`
	} `json:"jumbo-mtu"`
	UsbStatus  ApJoinProfileUsbModuleStatus  `json:"usb-module-status,omitempty"`
	DeviceMgmt ApJoinProfileDeviceMgmtParams `json:"device-mgmt,omitempty"`
	UserMgmt   ApJoinProfileUserMgmtParams   `json:"user-mgmt,omitempty"`
	CoreDump   struct {
		Flag           MemoryCoreDumpProp `json:"coredump-flag"`
		TFTPServerAddr string             `json:"tftp-server-name"`
		FileName       string             `json:"corefile-name"`
	} `json:"coredump"`
	Syslog struct {
		FacValue SyslogFacProp   `json:"facility-value"`
		LogLevel SyslogLevelProp `json:"log-level"`
		Host     string          `json:"host"`
		TLSMode  bool            `json:"tls-mode"`
	} `json:"syslog"`
	Cdp struct {
		Enabled bool `json:"cdp-enable"`
	} `json:"cdp"`
	NTPServer         string `json:"ntp-server"`
	BackupControllers struct {
		Enabled         bool   `json:"fallback-enabled"`
		PrimaryName     string `json:"primary-controller-name"`
		SecondaryName   string `json:"secondary-controller-name"`
		PrimaryIPAddr   string `json:"primary-controller-ip"`
		SecondaryIPAddr string `json:"secondary-controller-ip"`
	} `json:"backup-controllers"`
}

// ApJoinProfileUsbModuleStatus ...
type ApJoinProfileUsbModuleStatus struct {
	Enabled bool `json:"enable"`
}

// ApJoinProfileDeviceMgmtParams ...
type ApJoinProfileDeviceMgmtParams struct {
	Telnet bool `json:"telnet"`
	SSH    bool `json:"ssh"`
}

// ApJoinProfileUserMgmtParams ...
type ApJoinProfileUserMgmtParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// Default: CLEAR
	PasswordType CrypType `json:"password-type"`
	Secret       string   `json:"secret"`
	// Default: CLEAR
	SecretType CrypType `json:"secret-type"`
}

// CrypType ...
type CrypType string

// CrypTypes ...
type CrypTypes struct {
	Clear         CrypType
	SHA           CrypType
	MD5           CrypType
	IkeAES        CrypType
	Type7         CrypType
	AES           CrypType
	ClearToSha    CrypType
	ClearToMd5    CrypType
	ClearToIkeAes CrypType
	ClearToType7  CrypType
	ClearToAes    CrypType
}

// MemoryCoreDumpProp ...
type MemoryCoreDumpProp string

// MemoryCoreDumpFlag ...
type MemoryCoreDumpFlag struct {
	Compress   MemoryCoreDumpProp
	Uncompress MemoryCoreDumpProp
	Disable    MemoryCoreDumpProp
}

// SyslogFacProp ...
type SyslogFacProp string

// SyslogFacEnum ...
type SyslogFacEnum struct {
	Kern   SyslogFacProp
	User   SyslogFacProp
	Mail   SyslogFacProp
	Daemon SyslogFacProp
	Auth   SyslogFacProp
	Syslog SyslogFacProp
	Lpr    SyslogFacProp
	News   SyslogFacProp
	UUCP   SyslogFacProp
	Sys9   SyslogFacProp
	Sys10  SyslogFacProp
	Sys11  SyslogFacProp
	Sys12  SyslogFacProp
	Sys13  SyslogFacProp
	Sys14  SyslogFacProp
	Cron   SyslogFacProp
	Local0 SyslogFacProp
	Local1 SyslogFacProp
	Local2 SyslogFacProp
	Local3 SyslogFacProp
	Local4 SyslogFacProp
	Local5 SyslogFacProp
	Local6 SyslogFacProp
	Local7 SyslogFacProp
}

// SyslogLevelProp ...
type SyslogLevelProp string

// SyslogLevelEnum ...
type SyslogLevelEnum struct {
	Emergency    SyslogLevelProp
	Alert        SyslogLevelProp
	Critical     SyslogLevelProp
	Errors       SyslogLevelProp
	Warning      SyslogLevelProp
	Notification SyslogLevelProp
	Info         SyslogLevelProp
	Debug        SyslogLevelProp
}

var (
	// CryptTypes ...
	CryptTypes = CrypTypes{
		Clear:         "clear",
		SHA:           "sha",
		MD5:           "md5",
		IkeAES:        "ike-aes",
		Type7:         "type7",
		AES:           "aes",
		ClearToSha:    "clear-to-sha",
		ClearToMd5:    "clear-to-md5",
		ClearToIkeAes: "clear-to-ike-aes",
		ClearToType7:  "clear-to-type7",
		ClearToAes:    "clear-to-aes",
	}
	// MemCoreDump ...
	MemCoreDump = MemoryCoreDumpFlag{
		Disable:    "tftp-coredump-disable",
		Compress:   "tftp-coredump-compress",
		Uncompress: "tftp-coredump-uncompress",
	}
	// SyslogFacility ...
	SyslogFacility = SyslogFacEnum{
		Kern:   "facility-kern",
		User:   "facility-user",
		Mail:   "facility-mail",
		Daemon: "facility-daemon",
		Auth:   "facility-auth",
		Syslog: "facility-syslog",
		Lpr:    "facility-lpr",
		News:   "facility-news",
		UUCP:   "facility-uucp",
		Sys9:   "facility-sys9",
		Sys10:  "facility-sys10",
		Sys11:  "facility-sys11",
		Sys12:  "facility-sys12",
		Sys13:  "facility-sys13",
		Sys14:  "facility-sys14",
		Cron:   "facility-cron",
		Local0: "facility-local0",
		Local1: "facility-local1",
		Local2: "facility-local2",
		Local3: "facility-local3",
		Local4: "facility-local4",
		Local5: "facility-local5",
		Local6: "facility-local6",
		Local7: "facility-local7",
	}
	// SyslogLevel ...
	SyslogLevel = SyslogLevelEnum{
		Emergency: "syslog-level-emergency",
		Alert:     "syslog-level-alert",
		Critical:  "syslog-level-critical",
		Errors:    "syslog-level-errors",
		Warning:   "syslog-level-warning",
		Info:      "syslog-level-information",
		Debug:     "syslog-level-debug",
	}
)

// Tag ...
type Tag struct {
	Name        string `json:"site-tag-name"`
	Description string `json:"description"`
	// Only Enabled if Site IS NOT Local
	FlexProfile   string `json:"flex-profile,omitempty"`
	ApJoinProfile string `json:"ap-join-profile"`
	// Default: TRUE
	IsLocalSite          bool   `json:"is-local-site"`
	FabricCntrlPlaneName string `json:"fabric-control-plane-name,omitempty"`
	// Default: default
	ImageDownloadProfileName string `json:"image-download-profile-name,omitempty"`
	// Default: TRUE
	ArpCaching bool `json:"arp-caching"`
}
