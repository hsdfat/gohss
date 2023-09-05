package cxdx

import (
	"fmt"
	"gohss/hss_models"
	"gohss/messages"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/datatype"

	"github.com/flosch/pongo2/v6"
)

func NewSAA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
) (*diam.Message, error) {
	var saa SAA
	var sar SAR
	if err := msg.Unmarshal(&sar); err != nil {
		return msg.Answer(diam.UnableToComply), fmt.Errorf("SAR Unmarshal failed for message: %v failed: %v", msg, err)
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
	template := srv.Config.CxTemplate
	if template != nil {
		out, err := template.Execute(pongo2.Context{"iFC_vars": map[string]string{
			"imsi": "1",
			"mnc":  "452",
			"mcc":  "04",
		}})
		if err != nil {
			return messages.ConstructFailureAnswer(msg, sar.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err
		}
		saa.UserData = datatype.OctetString(out)
		// fmt.Println(out)

	}

	return NewSuccessfulSAA(srv, sar.SessionID, msg, &saa), nil
}

// NewSuccessfulAIA outputs a successful authentication information answer (AIA) to reply to an
// authentication information request (AIR) message. It populates AIA with all of the mandatory fields
// and adds the authentication vectors.
func NewSuccessfulSAA(
	srv *hss_models.HomeSubscriberServer,
	sessionID datatype.UTF8String,
	msg *diam.Message,
	saa *SAA,
) *diam.Message {
	// vendorID := srv.GetVendorID()
	answer := messages.ConstructSuccessAnswer(msg, sessionID, srv.Config.Server, diam.CX_APP_ID)
	// answer := msg.Answer(2001)
	saa.SupportedFeatures = []*SupportedFeatures{
		{
			VendorId:      srv.GetVendorID(),
			FeatureListID: 1,
			FeatureList:   5,
		},
	}
	// fmt.Println(saa.UserData)
	answer.Marshal(saa)

	return answer
}
