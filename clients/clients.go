package clients

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ApogeeNetworking/ciscowireless/requests"
)

// Service ...
type Service struct {
	http *requests.Service
}

// NewService ...
func NewService(req *requests.Service) *Service {
	return &Service{http: req}
}

// Client ...
type Client struct {
	MacAddr               string `json:"client-mac"`
	ApName                string `json:"ap-name"`
	MsApSlotID            int    `json:"ms-ap-slot-id"`
	MsRadioType           string `json:"ms-radio-type"`
	WlanID                int    `json:"wlan-id"`
	ClientType            string `json:"client-type"`
	CoState               string `json:"co-state"`
	AaaOverridePassphrase bool   `json:"aaa-override-passphrase"`
	IsTviEnabled          bool   `json:"is-tvi-enabled"`
	WlanPolicy            struct {
		CurrentSwitchingMode  string `json:"current-switching-mode"`
		WlanSwitchingMode     string `json:"wlan-switching-mode"`
		CentralAuthentication string `json:"central-authentication"`
		CentralDhcp           bool   `json:"central-dhcp"`
		CentralAssocEnable    bool   `json:"central-assoc-enable"`
		VlanCentralSwitching  bool   `json:"vlan-central-switching"`
		IsFabricClient        bool   `json:"is-fabric-client"`
		IsGuestFabricClient   bool   `json:"is-guest-fabric-client"`
		UpnBitFlag            string `json:"upn-bit-flag"`
	} `json:"wlan-policy"`
	Username           string `json:"username"`
	GuestLanClientInfo struct {
		WiredVlan       int `json:"wired-vlan"`
		PhyIfid         int `json:"phy-ifid"`
		IdleTimeSeconds int `json:"idle-time-seconds"`
	} `json:"guest-lan-client-info"`
	MethodID                 string    `json:"method-id"`
	L3VlanOverrideReceived   bool      `json:"l3-vlan-override-received"`
	UpnID                    int       `json:"upn-id"`
	IsLocallyAdministeredMac bool      `json:"is-locally-administered-mac"`
	IdleTimeout              int       `json:"idle-timeout"`
	IdleTimestamp            time.Time `json:"idle-timestamp"`
}

// Get retrieve common client oper data
func (s *Service) Get() ([]Client, error) {
	uri := "/Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data"
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return nil, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resp := struct {
		Response []Client `json:"Cisco-IOS-XE-wireless-client-oper:common-oper-data"`
	}{}
	json.NewDecoder(res.Body).Decode(&resp)
	return resp.Response, nil
}

// Traffic ...
type Traffic struct {
	MacAddr   string `json:"ms-mac-address"`
	BytesRx   string `json:"bytes-rx"`
	BytesTx   string `json:"bytes-tx"`
	PacketsRx string `json:"pkts-rx"`
	PacketsTx string `json:"pkts-tx"`
	RSSI      int    `json:"most-recent-rssi"`
	SNR       int    `json:"most-recent-snr"`
	Speed     int    `json:"speed"`
}

// GetClientStats ...
func (s *Service) GetClientStats(mac string) (Traffic, error) {
	uri := fmt.Sprintf("/Cisco-IOS-XE-wireless-client-oper:client-oper-data/traffic-stats=%s", mac)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return Traffic{}, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return Traffic{}, err
	}
	defer res.Body.Close()
	resp := struct {
		Response Traffic `json:"Cisco-IOS-XE-wireless-client-oper:traffic-stats"`
	}{}
	json.NewDecoder(res.Body).Decode(&resp)
	return resp.Response, nil
}
