package main

import (
	"log"
	"net"
	"time"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
	"gohss/modules/go-diameter/v4/diam/datatype"
	"gohss/modules/go-diameter/v4/diam/dict"
	"gohss/modules/go-diameter/v4/diam/sm"
)

func main() {
	addr := "127.0.0.1:3768"
	cfg := &sm.Settings{
		OriginHost:       datatype.DiameterIdentity("123"),
		OriginRealm:      datatype.DiameterIdentity("234"),
		VendorID:         13,
		ProductName:      "go-diameter",
		OriginStateID:    datatype.Unsigned32(time.Now().Unix()),
		FirmwareRevision: 1,
		HostIPAddresses: []datatype.Address{
			datatype.Address(net.ParseIP("127.0.0.1")),
		},
	}
	mux := sm.New(cfg)
	client := &sm.Client{
		Dict:               dict.Default,
		Handler:            mux,
		MaxRetransmits:     3,
		RetransmitInterval: time.Second,
		EnableWatchdog:     true,
		WatchdogInterval:   5 * time.Second,
		AcctApplicationID: []*diam.AVP{
			// Advertise that we want support accounting application with id 999
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(diam.TGPP_S6A_APP_ID)),
		},
		AuthApplicationID: []*diam.AVP{
			// Advertise support for credit control application
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(diam.TGPP_S6A_APP_ID)), // RFC 4006
		},
	}
	_, err := client.DialNetwork("sctp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	quit := make(chan bool, 1)
	<-quit
}
