package messages

import (
	"encoding/hex"
	"fmt"
	"gohss/hss_models"
	"gohss/utils/tbcd"
	"strconv"
	"strings"
	"time"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/avp"
	"gohss/modules/go-diameter/v4/diam/datatype"

	"github.com/golang/glog"
)

func ConstructFailureAnswer(msg *diam.Message, sessionID datatype.UTF8String, serverCfg *hss_models.DiameterConfig, resultCode uint32) *diam.Message {
	newMsg := diam.NewMessage(
		msg.Header.CommandCode,
		msg.Header.CommandFlags&^diam.RequestFlag, // Reset the Request bit.
		msg.Header.ApplicationID,
		msg.Header.HopByHopID,
		msg.Header.EndToEndID,
		msg.Dictionary(),
	)
	AddStandardAnswerAVPS(newMsg, sessionID, serverCfg, resultCode)
	return newMsg
}

// AddStandardAnswerAVPS adds the SessionID, ExperimentalResult, OriginHost, OriginRealm, and OriginStateID AVPs to a message.
func AddStandardAnswerAVPS(answer *diam.Message, sessionID datatype.UTF8String, serverCfg *hss_models.DiameterConfig, resultCode uint32) {
	// SessionID is required to be the AVP in position 1
	answer.InsertAVP(diam.NewAVP(avp.SessionID, avp.Mbit, 0, sessionID))
	if resultCode != diam.Success {
		answer.NewAVP(avp.ExperimentalResult, avp.Mbit, 0, &diam.GroupedAVP{
			AVP: []*diam.AVP{
				diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(serverCfg.Vendor3GPP)),
				diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, datatype.Unsigned32(resultCode)),
			},
		})
	}
	answer.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity(serverCfg.DestHost))
	answer.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity(serverCfg.DestRealm))
	answer.NewAVP(avp.OriginStateID, avp.Mbit, 0, datatype.Unsigned32(time.Now().Unix()))
}

type replyFunc func(*hss_models.HomeSubscriberServer, *diam.Message) (*diam.Message, error)

// handleMessage processes incoming request messages and sends an answer message
// back which is constructed using a replyFunc.
func HandleMessage(srv *hss_models.HomeSubscriberServer, reply replyFunc) diam.HandlerFunc {
	return func(conn diam.Conn, msg *diam.Message) {
		fmt.Println("HandleMessage here: ")
		// Add client connection to connection manager in case of HSS initiated message
		// err := srv.AddClientConnection(conn, msg)
		// // If the connection cannot be added, the HSS should still respond properly
		// if err != nil {
		// 	fmt.Println(err)
		// }
		if msg == nil {
			fmt.Println("Received nil message")
			return
		}
		glog.V(2).Infof("Message received in hss service: %s}", msg.String())

		answer, err := reply(srv, msg)
		if err != nil {
			fmt.Println(err)
		}

		_, err = answer.WriteTo(conn)
		if err != nil {
			fmt.Printf("Failed to send response: %s", err.Error())
		}
		// fmt.Printf("Message sent from hss service: %s}", answer.String())
	}
}

func ConvertBitRate(str string) uint32 {
	strs := strings.Split(str, " ")
	if len(strs) < 2 {
		return 0
	}
	n, err := strconv.ParseUint(strs[0], 10, 32)
	num := uint32(n)
	if err != nil {
		return 0
	}
	switch strings.ToLower(strs[1]) {
	case "bps":
		return num
	case "kbps":
		return num * 1000
	case "mbps":
		return num * 1000000
	case "gbps":
		return num * 1000000000
	case "tbps":
		return 1 << 31
	}
	return 1000
}

func ConvertPDNType(str string) uint32 {
	switch str {
	case "IPV4":
		return 0
	case "IPV6":
		return 1
	case "IPV4V6":
		return 2
	case "UNSTRUCTURED":
	case "ETHERNET":
	}
	return 0
}

func ConvertPreeCapt(str string) uint32 {
	switch str {
	case "NOT_PREEMPT":
		return 1
	case "MAY_PREEMPT":
		return 0
	}
	return 1
}

func ConvertPreeVul(str string) uint32 {
	switch str {
	case "NOT_PREEMPTABLE":
		return 1
	case "PREEMPTABLE":
		return 0
	}
	return 0
}

func ConvertMSISDN(str string) string {
	// str = hex.DecodeString()
	msisdn, err := tbcd.EncodeToTBCD(str)
	if err != nil {
		return ""
	}
	hexstr, err := hex.DecodeString(msisdn)
	if err != nil {
		return ""
	}
	return string(hexstr)
}

func NewUint32(num uint32) *uint32 {
	return &num
}

func NewInt32(num int32) *int32 {
	return &num
}

func NewInt(num int) *int {
	return &num
}
