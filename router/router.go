package router

import (
	hss_models "gohss/hss_models"
	"gohss/messages"
	"gohss/service/cxdx"
	"gohss/service/s6a"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
	"gohss/modules/go-diameter/v4/diam/datatype"
	"gohss/modules/go-diameter/v4/diam/sm"

	"github.com/emakeev/milenage"
)

func InitRouter(srv *hss_models.HomeSubscriberServer, stateMachine *sm.StateMachine) error {
	// stateMachine.HandleFunc(diam.AIR) // default handler
	stateMachine.HandleFunc(diam.AIR, messages.HandleMessage(srv, s6a.NewAIA))
	stateMachine.HandleFunc(diam.ULR, messages.HandleMessage(srv, s6a.NewULA))
	stateMachine.HandleFunc(diam.MAR, messages.HandleMessage(srv, cxdx.NewMAA))
	stateMachine.HandleFunc(diam.SAR, messages.HandleMessage(srv, cxdx.NewSAA))
	// mux.GetMux().Handle(diam.MAR, srv.handleMessage(NewMAA))
	// mux.GetMux().Handle(diam.SAR, srv.handleMessage(NewSAA))
	// mux.GetMux().HandleIdx(
	// 	diam.CommandIndex{AppID: diam.TGPP_SWX_APP_ID, Code: diam.RegistrationTermination, Request: false},
	// 	handleRTA(srv))
	return nil
}

func NewSuccessfulAIA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
	sessionID datatype.UTF8String,
	vectors []*milenage.EutranVector,
	utranVectors []*milenage.UtranVector,
) *diam.Message {
	vendorID := srv.GetVendorID()
	answer := messages.ConstructSuccessAnswer(msg, sessionID, srv.Config.Server, diam.TGPP_S6A_APP_ID)
	evs := []*diam.AVP{}

	answer.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit|avp.Vbit, vendorID, &diam.GroupedAVP{AVP: evs})
	return answer
}
