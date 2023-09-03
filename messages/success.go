package messages

import (
	"gohss/hss_models"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
	"gohss/modules/go-diameter/v4/diam/datatype"
)

// ConstructSuccessAnswer returns a message response with a success result code
// and with the server config AVPs already added.
func ConstructSuccessAnswer(msg *diam.Message, sessionID datatype.UTF8String, serverCfg *hss_models.DiameterConfig, authApplicationID uint32) *diam.Message {
	answer := msg.Answer(diam.Success)
	AddStandardAnswerAVPS(answer, sessionID, serverCfg, diam.Success)
	answer.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(serverCfg.Vendor3GPP)),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(authApplicationID)),
		},
	})
	answer.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(AuthSessionState_NO_STATE_MAINTAINED))
	return answer
}
