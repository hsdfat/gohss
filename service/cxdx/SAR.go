package cxdx

import (
	"fmt"
	"gohss/hss_models"
	"gohss/messages"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
	"gohss/modules/go-diameter/v4/diam/datatype"
)

func NewSAA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
) (*diam.Message, error) {
	var saa SAA
	var sar SAR
	if err := msg.Unmarshal(sar); err != nil {
		return msg.Answer(diam.UnableToComply), fmt.Errorf("AIR Unmarshal failed for message: %v failed: %v", msg, err)
	}

	if sar.UserName != "" {
		// TODO: Checking if the username exists
	}

	// Gen username
	imsi := ""
	domain := ""

	// TODO: Get ims subscribers details

	// TODO: gen userdata

	// TODO: Gen respone username
	saa.UserName = datatype.UTF8String(fmt.Sprintf("%s@%s", imsi, domain))
	// TODO: Gen charging information
	saa.ChargingInformation = &ChargingInformation{
		PrimaryEventChargingFunctionName: "123",
	}
	// TODO: Get server assign type
	ServerAssignType := 1
	if ServerAssignType == 1 || ServerAssignType == 2 {
		// Update serving CSCF from originHost
	} else {
		// Clear S-CSCF 
	}

	return NewSuccessfulSAA(srv, msg, &saa), nil
}

// NewSuccessfulAIA outputs a successful authentication information answer (AIA) to reply to an
// authentication information request (AIR) message. It populates AIA with all of the mandatory fields
// and adds the authentication vectors.
func NewSuccessfulSAA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
	uaa *SAA,
) *diam.Message {
	// vendorID := srv.GetVendorID()
	// answer := messages.ConstructSuccessAnswer(msg, sessionID, srv.Config.Server, diam.TGPP_S6A_APP_ID)
	answer := msg.Answer(2001)
	answer.Marshal(uaa)
	answer.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(srv.GetVendorID())),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(diam.TGPP_S6A_APP_ID)),
		},
	})
	answer.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(messages.AuthSessionState_NO_STATE_MAINTAINED))
	return answer
}
