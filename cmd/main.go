package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ApogeeNetworking/ciscowireless"
	"github.com/subosito/gotenv"
)

var wlcHost, wlcUser, wlcPass string

func init() {
	gotenv.Load()
	wlcHost = os.Getenv("WLC_HOST")
	wlcUser = os.Getenv("WLC_USER")
	wlcPass = os.Getenv("WLC_PASS")
}

func main() {
	client := ciscowireless.NewService(wlcHost, wlcUser, wlcPass, true)
	wlans, err := client.Wlans.GetPolicyTags()
	if err != nil {
		log.Fatal(err)
	}
	for _, wlan := range wlans {
		fmt.Println(wlan)
	}
}
