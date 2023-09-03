package cxdx

import (
	"gohss/hss_models"
	"gohss/messages"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
	"gohss/modules/go-diameter/v4/diam/datatype"
)

func NewLIA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
) (*diam.Message, error) {
	var lia LIA

	return NewSuccessfulLIA(srv, msg, &lia), nil
}

// NewSuccessfulAIA outputs a successful authentication information answer (AIA) to reply to an
// authentication information request (AIR) message. It populates AIA with all of the mandatory fields
// and adds the authentication vectors.
func NewSuccessfulLIA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
	lia *LIA,
) *diam.Message {
	// vendorID := srv.GetVendorID()
	// answer := messages.ConstructSuccessAnswer(msg, sessionID, srv.Config.Server, diam.TGPP_S6A_APP_ID)
	answer := msg.Answer(2001)
	answer.Marshal(lia)
	answer.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(srv.GetVendorID())),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(diam.TGPP_S6A_APP_ID)),
		},
	})
	answer.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(messages.AuthSessionState_NO_STATE_MAINTAINED))
	return answer
}
