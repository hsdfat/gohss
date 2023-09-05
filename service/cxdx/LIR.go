package cxdx

import (
	"fmt"
	"gohss/hss_models"
	"gohss/messages"
	"strings"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/datatype"
)

func NewLIA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
) (*diam.Message, error) {
	var lia LIA
	var lir LIR
	err := msg.Unmarshal(&lir)
	if err != nil {
		return msg.Answer(diam.UnableToComply), fmt.Errorf("AIR Unmarshal failed for message: %v failed: %v", msg, err)
	}
	if lir.UserName == "" {

		return messages.ConstructFailureAnswer(msg, lir.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err
	}
	strs := strings.Split(string(lir.UserName), "@")
	if len(strs) != 2 {
		return messages.ConstructFailureAnswer(msg, lir.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err
	}
	//TODO: get Ims subscriber

	// TODO: check existence of server name
	serverName := ""
	if serverName != "" {
		lia.ServerName = datatype.UTF8String(serverName)
	} else {
		// Default
	}

	return NewSuccessfulLIA(srv, lir.SessionID, msg, &lia), nil
}

// NewSuccessfulAIA outputs a successful authentication information answer (AIA) to reply to an
// authentication information request (AIR) message. It populates AIA with all of the mandatory fields
// and adds the authentication vectors.
func NewSuccessfulLIA(
	srv *hss_models.HomeSubscriberServer,
	sessionID datatype.UTF8String,
	msg *diam.Message,
	lia *LIA,
) *diam.Message {
	// vendorID := srv.GetVendorID()
	answer := messages.ConstructSuccessAnswer(msg, sessionID, srv.Config.Server, diam.TGPP_S6A_APP_ID)
	// answer := msg.Answer(2001)
	lia.LIAFlags = 1
	answer.Marshal(lia)

	return answer
}
