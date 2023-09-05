package cxdx

import (
	"fmt"
	"gohss/hss_models"
	"gohss/messages"
	"strings"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/datatype"
)

func NewUAA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
) (*diam.Message, error) {
	var uar UAR
	var uaa UAA
	err := msg.Unmarshal(&uar)
	if err != nil {
		// TODO:
		return msg.Answer(diam.UnableToComply), fmt.Errorf("AIR Unmarshal failed for message: %v failed: %v", msg, err)
	}
	// TODO: checking username present
	if uar.UserName == "" {
	}
	strs := strings.Split(string(uar.UserName), "@")
	if len(strs) != 2 {
		return messages.ConstructFailureAnswer(msg, uar.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err
	}
	// imsi := strs[0]
	// domain := strs[1]
	// TODO: Get ims subscribers
	//imsSubscriber :=
	UserDataType := uar.UserAuthorizationType

	if UserDataType != nil && *UserDataType == 1 {
		// This is Derigister
		// TODO: Update Serving S-CSCF address to empty

		return NewSuccessfulUAA(srv, uar.SessionID, msg, &uaa), nil
	}
	// Checking existing S-CSCF
	if true {
		// TODO: Existed S-CSCF address
		serverName := ""
		uaa.ServerName = datatype.UTF8String(serverName)
		uaa.ExperimentalResult = &ExperimentalResult{
			VendorId:               srv.GetVendorID(),
			ExperimentalResultCode: 2002,
		}
	} else {
		// TODO: No S-CSCF address assigned
		// Default server name
		serverName := ""
		uaa.ServerName = datatype.UTF8String(serverName)
		uaa.ExperimentalResult = &ExperimentalResult{
			VendorId:               srv.GetVendorID(),
			ExperimentalResultCode: 2001,
		}
	}

	return NewSuccessfulUAA(srv, uar.SessionID, msg, &uaa), nil
}

// NewSuccessfulAIA outputs a successful authentication information answer (AIA) to reply to an
// authentication information request (AIR) message. It populates AIA with all of the mandatory fields
// and adds the authentication vectors.
func NewSuccessfulUAA(
	srv *hss_models.HomeSubscriberServer,
	sessionID datatype.UTF8String,
	msg *diam.Message,
	uaa *UAA,
) *diam.Message {
	// vendorID := srv.GetVendorID()
	answer := messages.ConstructSuccessAnswer(msg, sessionID, srv.Config.Server, diam.CX_APP_ID)
	uaa.SupportedFeatures = []*SupportedFeatures{
		{
			VendorId:      srv.GetVendorID(),
			FeatureListID: 1,
			FeatureList:   5,
		},
	}
	answer.Marshal(uaa)

	return answer
}
