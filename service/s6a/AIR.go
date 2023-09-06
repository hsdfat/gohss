package s6a

import (
	"fmt"
	"gohss/hss_models"
	"gohss/messages"
	"gohss/utils"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
	"gohss/modules/go-diameter/v4/diam/datatype"

	"gohss/modules/models"

	"github.com/emakeev/milenage"
)

func NewAIA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
) (*diam.Message, error) {
	// if err := ValidateAIR(msg, srv); err != nil {
	// 	return msg.Answer(diam.MissingAVP), err
	// }

	var air AIR
	if err := msg.Unmarshal(&air); err != nil {
		return msg.Answer(diam.UnableToComply), fmt.Errorf("AIR Unmarshal failed for message: %v failed: %v", msg, err)
	}

	subscriber, err := srv.Store.GetAuthenSubscriptionData(air.UserName)
	if err != nil {
		// if _, ok := err.(storage.UnknownSubscriberError); ok {
		// 	return messages.ConstructFailureAnswer(msg, air.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err
		// }
		return messages.ConstructFailureAnswer(msg, air.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err
	}

	// subscriber.Lock()
	// defer subscriber.Unlock()
	var nextSQN uint64
	if air.RequestedEUTRANAuthInfo != nil {
		resynInfo := air.RequestedEUTRANAuthInfo.ResyncInfo.Serialize()
		if len(resynInfo) > 0 {

			nextSQN, err = utils.ResyncLteAuthSeq(
				subscriber, air.RequestedEUTRANAuthInfo.ResyncInfo.Serialize(), srv.Config.AuthOp)
			if err != nil {
				return messages.ConvertAuthErrorToFailureMessage(err, msg, air.SessionID, srv.Config.Server), err
			}
			subscriber.SequenceNumber = utils.IntToHexFixedLength(nextSQN, 12)
		}
	}

	if air.RequestedUtranGeranAuthInfo != nil && len(air.RequestedUtranGeranAuthInfo.ResyncInfo) > 0 {
		lteAuthNextUtranSeq, err := utils.ResyncLteAuthSeq(
			subscriber, air.RequestedUtranGeranAuthInfo.ResyncInfo.Serialize(), srv.Config.AuthOp)
		if err != nil {
			return messages.ConvertAuthErrorToFailureMessage(err, msg, air.SessionID, srv.Config.Server), err
		}
		if air.RequestedEUTRANAuthInfo != nil && len(air.RequestedEUTRANAuthInfo.ResyncInfo) == 0 || lteAuthNextUtranSeq > nextSQN {
			nextSQN = lteAuthNextUtranSeq
		}
		subscriber.SequenceNumber = utils.IntToHexFixedLength(nextSQN, 12)
	}

	plmn := air.VisitedPLMNID.Serialize()
	var lteNumVector, utranNumVector uint32
	if air.RequestedEUTRANAuthInfo != nil {
		lteNumVector = uint32(air.RequestedEUTRANAuthInfo.NumVectors)
	}
	if air.RequestedUtranGeranAuthInfo != nil {
		lteNumVector = uint32(air.RequestedUtranGeranAuthInfo.NumVectors)
	}
	vectors, utranVectors, lteAuthNextSeq, err :=
		utils.GenerateLteAuthVectors(
			lteNumVector,
			utranNumVector,
			srv.Milenage,
			subscriber,
			plmn,
			srv.Config.AuthOp,
			srv.AuthSqnInd,
		)
	if err == nil {
		err = setAuthNextSeq(air.UserName, srv, subscriber, lteAuthNextSeq)
	}
	if err != nil {
		return messages.ConvertAuthErrorToFailureMessage(err, msg, air.SessionID, srv.Config.Server), err
	}

	return NewSuccessfulAIA(srv, msg, air.SessionID, vectors, utranVectors), nil
}

func setAuthNextSeq(id string, srv *hss_models.HomeSubscriberServer, subs *models.AuthenticationSubscription, lteAuthNextSeq uint64) error {
	subs.SequenceNumber = utils.IntToHexFixedLength(lteAuthNextSeq, 12)
	return srv.Store.UpdateAuthenSubscriptionData(id, subs)
}

// NewSuccessfulAIA outputs a successful authentication information answer (AIA) to reply to an
// authentication information request (AIR) message. It populates AIA with all of the mandatory fields
// and adds the authentication vectors.
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
	for itemNumber, vector := range vectors {
		evs = append(evs, diam.NewAVP(avp.EUTRANVector, avp.Mbit|avp.Vbit, vendorID, &diam.GroupedAVP{
			AVP: []*diam.AVP{
				diam.NewAVP(avp.ItemNumber, avp.Mbit|avp.Vbit, vendorID, datatype.Unsigned32(itemNumber+1)),
				diam.NewAVP(avp.RAND, avp.Mbit|avp.Vbit, vendorID, datatype.OctetString(vector.Rand[:])),
				diam.NewAVP(avp.XRES, avp.Mbit|avp.Vbit, vendorID, datatype.OctetString(vector.Xres[:])),
				diam.NewAVP(avp.AUTN, avp.Mbit|avp.Vbit, vendorID, datatype.OctetString(vector.Autn[:])),
				diam.NewAVP(avp.KASME, avp.Mbit|avp.Vbit, vendorID, datatype.OctetString(vector.Kasme[:])),
			},
		}))
	}
	for itemNumber, vector := range utranVectors {
		evs = append(evs, diam.NewAVP(avp.UTRANVector, avp.Mbit|avp.Vbit, vendorID, &diam.GroupedAVP{
			AVP: []*diam.AVP{
				diam.NewAVP(avp.ItemNumber, avp.Mbit|avp.Vbit, vendorID, datatype.Unsigned32(itemNumber+1)),
				diam.NewAVP(avp.RAND, avp.Mbit|avp.Vbit, vendorID, datatype.OctetString(vector.Rand[:])),
				diam.NewAVP(avp.XRES, avp.Mbit|avp.Vbit, vendorID, datatype.OctetString(vector.Xres[:])),
				diam.NewAVP(avp.AUTN, avp.Mbit|avp.Vbit, vendorID, datatype.OctetString(vector.Autn[:])),
				diam.NewAVP(
					avp.ConfidentialityKey,
					avp.Mbit|avp.Vbit,
					vendorID,
					datatype.OctetString(vector.ConfidentialityKey[:])),
				diam.NewAVP(
					avp.IntegrityKey,
					avp.Mbit|avp.Vbit,
					vendorID,
					datatype.OctetString(vector.IntegrityKey[:])),
			},
		}))
	}
	answer.NewAVP(avp.AuthenticationInfo, avp.Mbit|avp.Vbit, vendorID, &diam.GroupedAVP{AVP: evs})
	return answer
}
