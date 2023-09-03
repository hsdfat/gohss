package s6a

import (
	"errors"
	"gohss/hss_models"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
)

// ValidateAIR returns an error if the message is missing any mandatory AVPs.
// Mandatory AVPs are specified in 3GPP TS 29.272 Table 5.2.3.1.1/1
func ValidateAIR(msg *diam.Message, srv *hss_models.HomeSubscriberServer) error {
	vendorID := srv.GetVendorID()
	_, err := msg.FindAVP(avp.UserName, 0)
	if err != nil {
		return errors.New("Missing IMSI in message")
	}
	_, err = msg.FindAVP(avp.VisitedPLMNID, vendorID)
	if err != nil {
		return errors.New("Missing Visited PLMN ID in message")
	}
	_, err = msg.FindAVP(avp.RequestedEUTRANAuthenticationInfo, vendorID)
	if err != nil {
		_, err = msg.FindAVP(avp.RequestedUTRANGERANAuthenticationInfo, vendorID)
		if err != nil {
			return errors.New("Missing requested E-UTRAN and UTRAN/GERAN authentication info in message")
		}
	}
	_, err = msg.FindAVP(avp.SessionID, 0)
	if err != nil {
		return errors.New("Missing SessionID in message")
	}
	return nil
}
