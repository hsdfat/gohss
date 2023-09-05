package cxdx

import (
	"fmt"
	"gohss/hss_models"
	"gohss/messages"
	"gohss/utils"
	"strings"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/datatype"

	"github.com/omec-project/openapi/models"
)

func NewMAA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
) (*diam.Message, error) {
	var maa MAA
	var mar MAR
	err := msg.Unmarshal(&mar)
	if err != nil {
		return msg.Answer(diam.UnableToComply), fmt.Errorf("AIR Unmarshal failed for message: %v failed: %v", msg, err)
	}
	if mar.UserName == "" {
	}
	strs := strings.Split(string(mar.UserName), "@")
	if len(strs) != 2 && len(strs[0]) < 5 {
		return messages.ConstructFailureAnswer(msg, mar.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err
	}
	imsi := strs[0]
	fmt.Println(imsi)
	// Get authendata
	subscriber, err := srv.Store.GetAuthenSubscriptionData(imsi)
	if err != nil {
		return messages.ConstructFailureAnswer(msg, mar.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err

	}
	// TODO: resync sqn
	var nextSQN uint64
	if mar.SIPAuthDataItem != nil {
		if mar.SIPAuthDataItem.SIPAuthorization != "" {
			resynInfo := mar.SIPAuthDataItem.SIPAuthorization.Serialize()

			nextSQN, err = utils.ResyncLteAuthSeq(
				subscriber, resynInfo, srv.Config.AuthOp)
			if err != nil {
				return messages.ConvertAuthErrorToFailureMessage(err, msg, mar.SessionID, srv.Config.Server), err
			}
			subscriber.SequenceNumber = utils.IntToHexFixedLength(nextSQN, 12)

		}
	}

	// TODO: gen auth data
	nItem := uint32(mar.SIPNumberAuthItems)
	if nItem == 0 {
		nItem = 1
	}
	vectors, lteAuthNextSeq, err :=
		utils.GenerateSIPAuthVectors(
			nItem,
			srv.Milenage,
			subscriber,
			srv.Config.AuthOp,
			srv.AuthSqnInd,
		)
	if err == nil {
		err = setAuthNextSeq(imsi, srv, subscriber, lteAuthNextSeq)
		if len(vectors) > 0 {
			count := 0
			maa.SIPAuthDataItem = []*SIPAuthDataItem{}
			for _, v := range vectors {
				authDataItem := &SIPAuthDataItem{}
				authDataItem.ConfidentialityKey = datatype.OctetString(v.ConfidentialityKey[:])
				authDataItem.IntegrityKey = datatype.OctetString(v.IntegrityKey[:])
				authDataItem.SIPAuthenticationScheme = datatype.UTF8String("Digest-AKAv1-MD5")
				authDataItem.SIPAuthorization = datatype.OctetString(v.Xres[:])
				authDataItem.SIPAuthencate = datatype.OctetString(append(v.Rand[:], v.Autn[:]...))
				authDataItem.SIPItemNumber = datatype.Unsigned32(count)
				count++
				maa.SIPAuthDataItem = append(maa.SIPAuthDataItem, authDataItem)
			}
		}

	}
	return NewSuccessfulMAA(srv, mar.SessionID, msg, &maa), nil
}

func setAuthNextSeq(id string, srv *hss_models.HomeSubscriberServer, subs *models.AuthenticationSubscription, lteAuthNextSeq uint64) error {
	subs.SequenceNumber = utils.IntToHexFixedLength(lteAuthNextSeq, 12)
	return srv.Store.UpdateAuthenSubscriptionData(id, subs)
}

// NewSuccessfulAIA outputs a successful authentication information answer (AIA) to reply to an
// authentication information request (AIR) message. It populates AIA with all of the mandatory fields
// and adds the authentication vectors.
func NewSuccessfulMAA(
	srv *hss_models.HomeSubscriberServer,
	sessionID datatype.UTF8String,
	msg *diam.Message,
	maa *MAA,
) *diam.Message {
	// vendorID := srv.GetVendorID()
	answer := messages.ConstructSuccessAnswer(msg, sessionID, srv.Config.Server, diam.TGPP_S6A_APP_ID)
	answer.Marshal(maa)

	return answer
}
