package server

import (
	"gohss/hss_models"
	"gohss/router"
	"log"
	"time"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
	"gohss/modules/go-diameter/v4/diam/datatype"
	"gohss/modules/go-diameter/v4/diam/dict"
	"gohss/modules/go-diameter/v4/diam/sm"

	"github.com/emakeev/milenage"
)

func Start(srv *hss_models.HomeSubscriberServer, stateMachine *sm.StateMachine, started chan string) error {
	serverCfg := srv.Config.Server

	server := &diam.Server{
		Network: serverCfg.Protocol,
		Addr:    serverCfg.Address,
		Handler: stateMachine,
	}
	listener, err := diam.MultistreamListen(serverCfg.Protocol, serverCfg.Address)
	if err != nil {
		return err
	}
	localAddress := listener.Addr().String()
	if cap(started) > len(started) {
		started <- localAddress
	} else {
		go func() {
			started <- localAddress
		}() // non-buffered/full chan -> use a dedicated routine, it may block
	}
	return server.Serve(listener)
}

func SetupHssServer(
	config *hss_models.HSSConfig, store hss_models.SubscriberStore,
) *hss_models.HomeSubscriberServer {

	servicer, err := NewHomeSubscriberServer(*config)
	if err != nil {
		log.Fatalf("Error creating home subscriber server: %s", err)
	}
	servicer.Store = store
	err = store.Init()
	if err != nil {
		log.Fatal(err)
	}
	return servicer

}

func NewHomeSubscriberServer(config hss_models.HSSConfig) (*hss_models.HomeSubscriberServer, error) {
	mcipher, err := milenage.NewCipher(config.AuthAmf)
	if err != nil {
		return nil, err
	}

	return &hss_models.HomeSubscriberServer{
		Config:   config,
		Milenage: mcipher,
	}, nil
}

func StartDiameterServer(server *hss_models.HomeSubscriberServer) {
	startedChan := make(chan string, 1)
	go func() {
		localAddr := <-startedChan
		log.Printf("Started home subscriber server @ %s", localAddr)
	}()
	log.Printf("Starting home subscriber server with configs:\n\t%+v\n", server.Config)

	stateMachine, err := NewStateMachine(server)
	if err != nil {
		log.Printf("Error when start server:\n\t%+v\n", err.Error())

	}
	router.InitRouter(server, stateMachine)
	err = Start(server, stateMachine, startedChan) // blocks

	log.Fatal(err)
}

func NewStateMachine(srv *hss_models.HomeSubscriberServer) (*sm.StateMachine, error) {
	serverCfg := srv.Config.Server
	settings := &sm.Settings{
		OriginHost:       datatype.DiameterIdentity(serverCfg.DestHost),
		OriginRealm:      datatype.DiameterIdentity(serverCfg.DestRealm),
		VendorID:         datatype.Unsigned32(serverCfg.Vendor3GPP),
		ProductName:      datatype.UTF8String(srv.Config.Name),
		OriginStateID:    datatype.Unsigned32(time.Now().Unix()),
		FirmwareRevision: 1,
	}

	stateMachine := sm.New(settings)

	clientCfg := hss_models.DiameterClientConfig{}
	if clientCfg.WatchdogInterval == 0 {
		clientCfg.WatchdogInterval = 10
	}
	srv.SmClient = &sm.Client{
		Dict:               dict.Default,
		Handler:            stateMachine,
		MaxRetransmits:     clientCfg.Retransmits,
		RetransmitInterval: time.Second,
		EnableWatchdog:     true,
		WatchdogInterval:   time.Second * time.Duration(clientCfg.WatchdogInterval),
		SupportedVendorID: []*diam.AVP{
			diam.NewAVP(avp.SupportedVendorID, avp.Mbit, 0, datatype.Unsigned32(serverCfg.Vendor3GPP)),
		},
		VendorSpecificApplicationID: []*diam.AVP{
			diam.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(diam.TGPP_S6A_APP_ID)),
					diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(serverCfg.Vendor3GPP)),
				},
			}),
		},
	}

	return stateMachine, nil
}
