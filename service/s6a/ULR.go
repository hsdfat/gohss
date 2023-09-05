package s6a

import (
	"fmt"
	"gohss/hss_models"
	"gohss/messages"
	"strings"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
	"gohss/modules/go-diameter/v4/diam/datatype"
)

func NewULA(
	srv *hss_models.HomeSubscriberServer,
	msg *diam.Message,
) (*diam.Message, error) {
	var ulr ULR
	var ula ULA
	if err := msg.Unmarshal(&ulr); err != nil {
		return msg.Answer(diam.UnableToComply), fmt.Errorf("AIR Unmarshal failed for message: %v failed: %v", msg, err)
	}

	ula.SessionID = string(ulr.SessionID)
	smData, err := srv.Store.GetSMSubscriptionData(string(ulr.UserName))
	if err != nil {
		return messages.ConstructFailureAnswer(msg, ulr.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err
	}
	amData, err := srv.Store.GetAMSubscriptionData(string(ulr.UserName))
	if err != nil {
		return messages.ConstructFailureAnswer(msg, ulr.SessionID, srv.Config.Server, uint32(messages.ErrorCode_USER_UNKNOWN)), err
	}
	// TODO: Update serving MME
	// TODO: Convert AMBR
	if amData.SubscribedUeAmbr != nil {
		ula.SubscriptionData = &SubscriptionData{
			AMBR: &AMBR{},
		}
		if messages.ConvertBitRate(amData.SubscribedUeAmbr.Uplink) != 0 {
			ula.SubscriptionData.AMBR.MaxRequestedBandwidthUL = messages.NewUint32(messages.ConvertBitRate(amData.SubscribedUeAmbr.Uplink))
		}
		if messages.ConvertBitRate(amData.SubscribedUeAmbr.Downlink) != 0 {
			ula.SubscriptionData.AMBR.MaxRequestedBandwidthDL = messages.NewUint32(messages.ConvertBitRate(amData.SubscribedUeAmbr.Downlink))
		}
	}
	//TODO: Convert APN
	if smData.DnnConfigurations != nil && len(smData.DnnConfigurations) > 0 {
		if ula.SubscriptionData == nil {
			ula.SubscriptionData = &SubscriptionData{}
		}
		fmt.Println(smData)
		// ula.SubscriptionData.APNConfigurationProfile.AllAPNConfigurationsIncludedIndicator = messages.NewInt32(0)
		// ula.SubscriptionData.APNConfigurationProfile.ContextIdentifier = messages.NewUint32(1)
		ula.SubscriptionData.APNConfigurationProfile = &APNConfigurationProfile{}
		ula.SubscriptionData.APNConfigurationProfile.APNConfigs = []*APNConfiguration{}
		count := 1
		for key, dnn := range smData.DnnConfigurations {
			apn := &APNConfiguration{}
			ula.SubscriptionData.APNConfigurationProfile.APNConfigs = append(ula.SubscriptionData.APNConfigurationProfile.APNConfigs, apn)
			apn.ServiceSelection = key
			apn.ContextIdentifier = messages.NewUint32(uint32(count))
			count++
			if dnn.PduSessionTypes != nil && dnn.PduSessionTypes.DefaultSessionType != "" {
				apn.PDNType = messages.NewUint32(messages.ConvertPDNType(string(dnn.PduSessionTypes.DefaultSessionType)))
			}
			if dnn.SessionAmbr != nil {
				apn.AMBR = &AMBR{}
				dl := messages.ConvertBitRate(dnn.SessionAmbr.Downlink)
				if dl != 0 {
					apn.AMBR.MaxRequestedBandwidthDL = messages.NewUint32(dl)
				}

				ul := messages.ConvertBitRate(dnn.SessionAmbr.Uplink)
				if dl != 0 {
					apn.AMBR.MaxRequestedBandwidthUL = messages.NewUint32(ul)
				}
				if dnn.Var5gQosProfile != nil {
					apn.EPSSubscribedQoSProfile = &EPSSubscribedQoSProfile{}
					if dnn.Var5gQosProfile.Arp != nil {
						apn.EPSSubscribedQoSProfile.AllocationRetentionPriority = &AllocationRetentionPriority{}
						apn.EPSSubscribedQoSProfile.AllocationRetentionPriority.PriorityLevel = messages.NewUint32(uint32(dnn.Var5gQosProfile.Arp.PriorityLevel))
						if dnn.Var5gQosProfile.Arp.PreemptCap != "" {
							apn.EPSSubscribedQoSProfile.AllocationRetentionPriority.PreemptionCapability = messages.NewInt32(int32(messages.ConvertPreeCapt(string(dnn.Var5gQosProfile.Arp.PreemptCap))))
						}
						if dnn.Var5gQosProfile.Arp.PreemptVuln != "" {
							apn.EPSSubscribedQoSProfile.AllocationRetentionPriority.PreemptionVulnerability = messages.NewInt32(int32(messages.ConvertPreeVul(string(dnn.Var5gQosProfile.Arp.PreemptVuln))))
						}
						if dnn.Var5gQosProfile.Var5qi != 0 {
							apn.EPSSubscribedQoSProfile.QoSClassIdentifier = messages.NewInt32(int32(dnn.Var5gQosProfile.Var5qi))
						}
					}
				}
				if dnn.Var3gppChargingCharacteristics != "" {
					apn.TgppChargingCharacteristics = dnn.Var3gppChargingCharacteristics
				}
			}
		}

		// Todo: Convert MSISDN
		if amData.Gpsis != nil && amData.Gpsis[0] != "" {
			msisdn := strings.Trim(amData.Gpsis[0], "msisdn-")
			if ula.SubscriptionData == nil {
				ula.SubscriptionData = &SubscriptionData{}
			}
			ula.SubscriptionData.MSISDN = datatype.OctetString(messages.ConvertMSISDN(msisdn))
		}
		if ula.SubscriptionData == nil {
			ula.SubscriptionData = &SubscriptionData{}
		}
		ula.SubscriptionData.AccessRestrictionData = messages.NewUint32(0)
		ula.SubscriptionData.SubscriberStatus = messages.NewInt32(0)
		ula.SubscriptionData.NetworkAccessMode = messages.NewInt32(0)
		ula.SubscriptionData.SubscribedPeriodicRauTauTimer = messages.NewUint32(720)
	}

	return NewSuccessfulULA(srv, ulr.SessionID, msg, &ula), nil
}

// NewSuccessfulAIA outputs a successful authentication information answer (AIA) to reply to an
// authentication information request (AIR) message. It populates AIA with all of the mandatory fields
// and adds the authentication vectors.
func NewSuccessfulULA(
	srv *hss_models.HomeSubscriberServer,
	sessionID datatype.UTF8String,
	msg *diam.Message,
	ula *ULA,
) *diam.Message {
	// vendorID := srv.GetVendorID()
	answer := messages.ConstructSuccessAnswer(msg, sessionID, srv.Config.Server, diam.TGPP_S6A_APP_ID)
	ula.ULAFlags = messages.NewUint32(1)
	answer.Marshal(ula)
	answer.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(srv.GetVendorID())),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(diam.TGPP_S6A_APP_ID)),
		},
	})
	answer.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(messages.AuthSessionState_NO_STATE_MAINTAINED))
	return answer
}
