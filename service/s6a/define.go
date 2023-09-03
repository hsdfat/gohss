package s6a

// package servce implements S6a GRPC proxy service which sends AIR, ULR messages over diameter connection,
// waits (blocks) for diameter's AIAs, ULAs & returns their RPC representation

import (
	"gohss/modules/go-diameter/v4/diam/datatype"
)

const (
	// 3GPP 29.273 5.2.3.6
	RadioAccessTechnologyType_EUTRAN = 1004
)

// Definitions for AIA, see sample below:
//
//Authentication-Information-Answer (AIA)
//{Code:318,Flags:0x40,Version:0x1,Length:556,ApplicationId:16777251,HopByHopId:0x9105bf89,EndToEndId:0x16c85bed}
//	Session-Id {Code:263,Flags:0x40,Length:28,VendorId:0,Value:UTF8String{session;3420619691},Padding:2}
//	Authentication-Info {Code:1413,Flags:0xc0,Length:144,VendorId:10415,Value:Grouped{
//		E-UTRAN-Vector {Code:1414,Flags:0xc0,Length:132,VendorId:10415,Value:Grouped{
//			RAND {Code:1447,Flags:0xc0,Length:28,VendorId:10415,Value:OctetString{0xf122047125e8372054d2b31643878866},Padding:0},
//			XRES {Code:1448,Flags:0xc0,Length:20,VendorId:10415,Value:OctetString{0x2c6e30243a103f0e},Padding:0},
//			AUTN {Code:1449,Flags:0xc0,Length:28,VendorId:10415,Value:OctetString{0xe73b5091a37080007d1726fd84830ecc},Padding:0},
//			KASME {Code:1450,Flags:0xc0,Length:44,VendorId:10415,Value:OctetString{0x08083ce5b62fdbe542ba0a19c415411cfaf1db35b8832b1f8a9c7cb525824c21},Padding:0},
//		}}
//	}}
//	Authentication-Info {Code:1413,Flags:0xc0,Length:144,VendorId:10415,Value:Grouped{
//		E-UTRAN-Vector {Code:1414,Flags:0xc0,Length:132,VendorId:10415,Value:Grouped{
//			RAND {Code:1447,Flags:0xc0,Length:28,VendorId:10415,Value:OctetString{0x12c7eb54f10c4007f65e14315545ed25},Padding:0},
//			XRES {Code:1448,Flags:0xc0,Length:20,VendorId:10415,Value:OctetString{0x22aeae2a4713ee62},Padding:0},
//			AUTN {Code:1449,Flags:0xc0,Length:28,VendorId:10415,Value:OctetString{0xfb97e19addee80002be44eee2df02059},Padding:0},
//			KASME {Code:1450,Flags:0xc0,Length:44,VendorId:10415,Value:OctetString{0x342a6173dda12c7902d2048d70fd83806a5e66b6fced874ccddfa106c9d4e03f},Padding:0},
//		}}
//	}}
//	Authentication-Info {Code:1413,Flags:0xc0,Length:144,VendorId:10415,Value:Grouped{
//		E-UTRAN-Vector {Code:1414,Flags:0xc0,Length:132,VendorId:10415,Value:Grouped{
//			RAND {Code:1447,Flags:0xc0,Length:28,VendorId:10415,Value:OctetString{0x23ea3e0ebd90b06b87e07554ac65d85d},Padding:0},
//			XRES {Code:1448,Flags:0xc0,Length:20,VendorId:10415,Value:OctetString{0x4c4f47cf85b84db9},Padding:0},
//			AUTN {Code:1449,Flags:0xc0,Length:28,VendorId:10415,Value:OctetString{0xcc49b7b25775800011079582097b2e48},Padding:0},
//			KASME {Code:1450,Flags:0xc0,Length:44,VendorId:10415,Value:OctetString{0x671c99d4aeca35a90c4bb26028df37a5151322c837c86189635da3ac24979d43},Padding:0},
//		}}
//	}}
//	Auth-Session-State {Code:277,Flags:0x40,Length:12,VendorId:0,Value:Enumerated{0}}
//	Origin-Host {Code:264,Flags:0x40,Length:28,VendorId:0,Value:DiameterIdentity{hss.openair4G.eur},Padding:3}
//	Origin-Realm {Code:296,Flags:0x40,Length:24,VendorId:0,Value:DiameterIdentity{openair4G.eur},Padding:3}
//	Result-Code {Code:268,Flags:0x40,Length:12,VendorId:0,Value:Unsigned32{2001}}
//

type EUtranVector struct {
	ItemNumber datatype.Unsigned32  `avp:"Item-Number,omitempty"`
	RAND       datatype.OctetString `avp:"RAND,omitempty"`
	XRES       datatype.OctetString `avp:"XRES,omitempty"`
	AUTN       datatype.OctetString `avp:"AUTN,omitempty"`
	KASME      datatype.OctetString `avp:"KASME,omitempty"`
}

type UtranVector struct {
	ItemNumber datatype.Unsigned32  `avp:"Item-Number,omitempty"`
	RAND       datatype.OctetString `avp:"RAND,omitempty"`
	XRES       datatype.OctetString `avp:"XRES,omitempty"`
	AUTN       datatype.OctetString `avp:"AUTN,omitempty"`
	CK         datatype.OctetString `avp:"Confidentiality-Key,omitempty"`
	IK         datatype.OctetString `avp:"Integrity-Key,omitempty"`
}

type GeranVector struct {
	ItemNumber datatype.Unsigned32  `avp:"Item-Number,omitempty"`
	RAND       datatype.OctetString `avp:"RAND,omitempty"`
	SRES       datatype.OctetString `avp:"SRES,omitempty"`
	Kc         datatype.OctetString `avp:"Kc,omitempty"`
}

type ExperimentalResult struct {
	VendorId               uint32 `avp:"Vendor-Id,omitempty"`
	ExperimentalResultCode uint32 `avp:"Experimental-Result-Code,omitempty"`
}

type AuthenticationInfo struct {
	EUtranVectors []EUtranVector `avp:"E-UTRAN-Vector,omitempty"`
	UtranVectors  []UtranVector  `avp:"UTRAN-Vector,omitempty"`
	GeranVectors  []GeranVector  `avp:"GERAN-Vector,omitempty"`
}

type AIA struct {
	SessionID          string                    `avp:"Session-Id,omitempty"`
	ResultCode         uint32                    `avp:"Result-Code,omitempty"`
	OriginHost         datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm        datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState   int32                     `avp:"Auth-Session-State,omitempty"`
	ExperimentalResult ExperimentalResult        `avp:"Experimental-Result,omitempty"`
	AI                 AuthenticationInfo        `avp:"Authentication-Info,omitempty"`
}

// Definitions for ULA, see sample below:
//
// Update-Location-Answer (ULA)
// {Code:316,Flags:0x40,Version:0x1,Length:516,ApplicationId:16777251,HopByHopId:0x22910d0a,EndToEndId:0x8d330652}
//
//	Session-Id {Code:263,Flags:0x40,Length:24,VendorId:0,Value:UTF8String{session;89988919},Padding:0}
//	ULA-Flags {Code:1406,Flags:0xc0,Length:16,VendorId:10415,Value:Unsigned32{1}}
//	Subscription-Data {Code:1400,Flags:0xc0,Length:380,VendorId:10415,Value:Grouped{
//	   MSISDN {Code:701,Flags:0xc0,Length:20,VendorId:10415,Value:OctetString{0x33638060010f},Padding:2},
//	   Access-Restriction-Data {Code:1426,Flags:0xc0,Length:16,VendorId:10415,Value:Unsigned32{47}},
//	   Subscriber-Status {Code:1424,Flags:0xc0,Length:16,VendorId:10415,Value:Enumerated{0}},
//	   Network-Access-Mode {Code:1417,Flags:0xc0,Length:16,VendorId:10415,Value:Enumerated{2}},
//	   AMBR {Code:1435,Flags:0xc0,Length:44,VendorId:10415,Value:Grouped{
//	   	Max-Requested-Bandwidth-UL {Code:516,Flags:0xc0,Length:16,VendorId:10415,Value:Unsigned32{50000000}},
//	   	Max-Requested-Bandwidth-DL {Code:515,Flags:0xc0,Length:16,VendorId:10415,Value:Unsigned32{100000000}},
//	   }}
//	   APN-Configuration-Profile {Code:1429,Flags:0xc0,Length:240,VendorId:10415,Value:Grouped{
//	   	   Context-Identifier {Code:1423,Flags:0xc0,Length:16,VendorId:10415,Value:Unsigned32{0}},
//	   	   All-APN-Configurations-Included-Indicator {Code:1428,Flags:0xc0,Length:16,VendorId:10415,Value:Enumerated{0}},
//	   	   APN-Configuration {Code:1430,Flags:0xc0,Length:196,VendorId:10415,Value:Grouped{
//	   	   	  Context-Identifier {Code:1423,Flags:0xc0,Length:16,VendorId:10415,Value:Unsigned32{0}},
//	   	   	  PDN-Type {Code:1456,Flags:0xc0,Length:16,VendorId:10415,Value:Enumerated{0}},
//	   	   	  Service-Selection {Code:493,Flags:0xc0,Length:20,VendorId:10415,Value:UTF8String{oai.ipv4},Padding:0},
//	   	   	  EPS-Subscribed-QoS-Profile {Code:1431,Flags:0xc0,Length:88,VendorId:10415,Value:Grouped{
//	   	   		  QoS-Class-Identifier {Code:1028,Flags:0xc0,Length:16,VendorId:10415,Value:Enumerated{9}},
//	   	   		  Allocation-Retention-Priority {Code:1034,Flags:0x80,Length:60,VendorId:10415,Value:Grouped{
//	   	   			  Priority-Level {Code:1046,Flags:0x80,Length:16,VendorId:10415,Value:Unsigned32{15}},
//	   	   			  Pre-emption-Capability {Code:1047,Flags:0x80,Length:16,VendorId:10415,Value:Enumerated{1}},
//	   	   			  Pre-emption-Vulnerability {Code:1048,Flags:0x80,Length:16,VendorId:10415,Value:Enumerated{0}},
//	   	   		  }}
//	   	   	  }},
//	          TGPP-Charging-Characteristics {Code: 13,Flags:0x80,VendorId:10415,Value:UTF8String{12}},
//	   	   	  AMBR {Code:1435,Flags:0xc0,Length:44,VendorId:10415,Value:Grouped{
//	   	   	      Max-Requested-Bandwidth-UL {Code:516,Flags:0xc0,Length:16,VendorId:10415,Value:Unsigned32{50000000}},
//	   	   		  Max-Requested-Bandwidth-DL {Code:515,Flags:0xc0,Length:16,VendorId:10415,Value:Unsigned32{100000000}},
//	   	   	  }}
//	      }}
//	  }}
//	  Subscribed-Periodic-RAU-TAU-Timer {Code:1619,Flags:0x80,Length:16,VendorId:10415,Value:Unsigned32{120}},
//	}}
//	Auth-Session-State {Code:277,Flags:0x40,Length:12,VendorId:0,Value:Enumerated{0}}
//	Origin-Host {Code:264,Flags:0x40,Length:28,VendorId:0,Value:DiameterIdentity{hss.openair4G.eur},Padding:3}
//	Origin-Realm {Code:296,Flags:0x40,Length:24,VendorId:0,Value:DiameterIdentity{openair4G.eur},Padding:3}
//	Result-Code {Code:268,Flags:0x40,Length:12,VendorId:0,Value:Unsigned32{2001}}
type AMBR struct {
	MaxRequestedBandwidthUL *uint32 `avp:"Max-Requested-Bandwidth-UL,omitempty"`
	MaxRequestedBandwidthDL *uint32 `avp:"Max-Requested-Bandwidth-DL,omitempty"`
	ExtendMaxRequestedBwUL  *uint32 `avp:"Extended-Max-Requested-BW-UL,omitempty"`
	ExtendMaxRequestedBwDL  *uint32 `avp:"Extended-Max-Requested-BW-DL,omitempty"`
}

type AllocationRetentionPriority struct {
	PriorityLevel           *uint32 `avp:"Priority-Level,omitempty"`
	PreemptionCapability    *int32  `avp:"Pre-emption-Capability,omitempty"`
	PreemptionVulnerability *int32  `avp:"Pre-emption-Vulnerability,omitempty"`
}

type EPSSubscribedQoSProfile struct {
	QoSClassIdentifier          *int32                       `avp:"QoS-Class-Identifier,omitempty"`
	AllocationRetentionPriority *AllocationRetentionPriority `avp:"Allocation-Retention-Priority,omitempty"`
}

type APNConfiguration struct {
	ContextIdentifier           *uint32                  `avp:"Context-Identifier,omitempty"`
	PDNType                     *uint32                  `avp:"PDN-Type,omitempty"`
	ServiceSelection            string                   `avp:"Service-Selection,omitempty"`
	EPSSubscribedQoSProfile     *EPSSubscribedQoSProfile `avp:"EPS-Subscribed-QoS-Profile,omitempty"`
	AMBR                        *AMBR                    `avp:"AMBR,omitempty"`
	TgppChargingCharacteristics string                   `avp:"TGPP-Charging-Characteristics,omitempty"`
}

type APNConfigurationProfile struct {
	ContextIdentifier                     *uint32             `avp:"Context-Identifier,omitempty"`
	AllAPNConfigurationsIncludedIndicator *int32              `avp:"All-APN-Configurations-Included-Indicator,omitempty"`
	APNConfigs                            []*APNConfiguration `avp:"APN-Configuration,omitempty"`
}

type SubscriptionData struct {
	MSISDN                        datatype.OctetString     `avp:"MSISDN,omitempty"`
	AccessRestrictionData         *uint32                  `avp:"Access-Restriction-Data,omitempty"`
	SubscriberStatus              *int32                   `avp:"Subscriber-Status,omitempty"`
	NetworkAccessMode             *int32                   `avp:"Network-Access-Mode,omitempty"`
	AMBR                          *AMBR                    `avp:"AMBR,omitempty"`
	APNConfigurationProfile       *APNConfigurationProfile `avp:"APN-Configuration-Profile,omitempty"`
	SubscribedPeriodicRauTauTimer *uint32                  `avp:"Subscribed-Periodic-RAU-TAU-Timer,omitempty"`
	TgppChargingCharacteristics   string                   `avp:"TGPP-Charging-Characteristics,omitempty"`
	RegionalSubscriptionZoneCode  []datatype.OctetString   `avp:"Regional-Subscription-Zone-Code,omitempty"`
}

type ULA struct {
	SessionID          string                    `avp:"Session-Id,omitempty"`
	ULAFlags           *uint32                   `avp:"ULA-Flags,omitempty"`
	SubscriptionData   *SubscriptionData         `avp:"Subscription-Data,omitempty"`
	AuthSessionState   *int32                    `avp:"Auth-Session-State,omitempty"`
	ResultCode         *uint32                   `avp:"Result-Code,omitempty"`
	OriginHost         datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm        datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	ExperimentalResult *ExperimentalResult       `avp:"Experimental-Result,omitempty"`
	SupportedFeatures  []*SupportedFeatures      `avp:"Supported-Features,omitempty"`
}

type CLR struct {
	SessionID        string                    `avp:"Session-Id,omitempty"`
	AuthSessionState int32                     `avp:"Auth-Session-State,omitempty"`
	OriginHost       datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm      datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	CancellationType int32                     `avp:"Cancellation-Type,omitempty"`
	DestinationHost  datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName         string                    `avp:"User-Name,omitempty"`
}

// Definitions for PU
//
// PUR is Go representation of Purge-UE-Request message
//
//	< Purge-UE-Request> ::=	< Diameter Header: 321, REQ, PXY, 16777251 >
//	< Session-Id >
//	[ DRMP ]
//	[ Vendor-Specific-Application-Id ]
//	{ Auth-Session-State }
//	{ Origin-Host }
//	{ Origin-Realm }
//	[ Destination-Host ]
//	{ Destination-Realm }
//	{ User-Name }
//	[ OC-Supported-Features ]
//	[ PUR-Flags ]
//	*[ Supported-Features ]
//	[ EPS-Location-Information ]
//	*[ AVP ]
//	*[ Proxy-Info ]
//	*[ Route-Record ]
type PUR struct {
	SessionID                   string                      `avp:"Session-Id,omitempty"`
	DRMP                        uint32                      `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            int32                       `avp:"Auth-Session-State,omitempty"`
	OriginHost                  datatype.DiameterIdentity   `avp:"Origin-Host,omitempty"`
	OriginRealm                 datatype.DiameterIdentity   `avp:"Origin-Realm,omitempty"`
	DestinationHost             datatype.DiameterIdentity   `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity   `avp:"Destination-Realm,omitempty"`
	UserName                    datatype.UTF8String         `avp:"User-Name,omitempty"`
	OCSupportedFeatures         OCSupportedFeatures         `avp:"OC-Supported-Features,omitempty"`
	PURFlags                    uint32                      `avp:"PUR-Flags,omitempty"`
	SupportedFeatures           []SupportedFeatures         `avp:"Supported-Features,omitempty"`
}

// PUA is Go representation of Purge-UE-Answer message
//
//	< Purge-UE-Answer> ::=	< Diameter Header: 321, PXY, 16777251 >
//	< Session-Id >
//	[ DRMP ]
//	[ Vendor-Specific-Application-Id ]
//	*[ Supported-Features ]
//	[ Result-Code ]
//	[ Experimental-Result ]
//	{ Auth-Session-State }
//	{ Origin-Host }
//	{ Origin-Realm }
//	[ OC-Supported-Features ]
//	[ OC-OLR ]
//	*[ Load ]
//	[ PUA-Flags ]
//	*[ AVP ]
//	[ Failed-AVP ]
//	*[ Proxy-Info ]
//	*[ Route-Record ]
type PUA struct {
	SessionID                   string                      `avp:"Session-Id,omitempty"`
	DRMP                        uint32                      `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
	SupportedFeatures           []SupportedFeatures         `avp:"Supported-Features,omitempty"`
	ResultCode                  uint32                      `avp:"Result-Code,omitempty"`
	ExperimentalResult          ExperimentalResult          `avp:"Experimental-Result,omitempty"`
	AuthSessionState            int32                       `avp:"Auth-Session-State,omitempty"`
	OriginHost                  datatype.DiameterIdentity   `avp:"Origin-Host,omitempty"`
	OriginRealm                 datatype.DiameterIdentity   `avp:"Origin-Realm,omitempty"`
	OCSupportedFeatures         OCSupportedFeatures         `avp:"OC-Supported-Features,omitempty"`
	OC_OLR                      OC_OLR                      `avp:"OC-OLR,omitempty"`
	PUAFlags                    uint32                      `avp:"PUA-Flags,omitempty"`
}

// VendorSpecificApplicationId -> Vendor-Specific-Application-Id AVP
type VendorSpecificApplicationId struct {
	VendorId          uint32 `avp:"Vendor-Id,omitempty"`
	AuthApplicationId uint32 `avp:"Auth-Application-Id,omitempty"`
	AcctApplicationId uint32 `avp:"Acct-Application-Id,omitempty"`
}

// SupportedFeatures -> Supported-Features AVP
type SupportedFeatures struct {
	VendorId      uint32 `avp:"Vendor-Id,omitempty"`
	FeatureListID uint32 `avp:"Feature-List-ID,omitempty"`
	FeatureList   uint32 `avp:"Feature-List,omitempty"`
}

// OCSupportedFeatures -> OC-Supported-Features AVP
type OCSupportedFeatures struct {
	OCFeatureVector uint64 `avp:"OC-Feature-Vector,omitempty"`
}

// OC_OLR -> OC-OLR AVP
type OC_OLR struct {
	OCSequenceNumber      uint64 `avp:"OC-Sequence-Number,omitempty"`
	OCReportType          uint32 `avp:"OC-Report-Type,omitempty"`
	OCReductionPercentage uint32 `avp:"OC-Reduction-Percentage,omitempty"`
	OCValidityDuration    uint32 `avp:"OC-Validity-Duration,omitempty"`
}

// RSR is Go representation of Reset-Request message
//
// < Reset-Request> ::= < Diameter Header: 322, REQ, PXY, 16777251 >
//
// < Session-Id >
// [ Vendor-Specific-Application-Id ]
// { Auth-Session-State }
// { Origin-Host }
// { Origin-Realm }
// { Destination-Host }
// { Destination-Realm }
// *[ Supported-Features ]
// *[ User-Id ]
// *[ AVP ]
// *[ Proxy-Info ]
// *[ Route-Record ]
type RSR struct {
	SessionID                   string                      `avp:"Session-Id,omitempty"`
	VendorSpecificApplicationId VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            int32                       `avp:"Auth-Session-State,omitempty"`
	OriginHost                  datatype.DiameterIdentity   `avp:"Origin-Host,omitempty"`
	OriginRealm                 datatype.DiameterIdentity   `avp:"Origin-Realm,omitempty"`
	DestinationHost             datatype.DiameterIdentity   `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity   `avp:"Destination-Realm,omitempty"`
	SupportedFeatures           []SupportedFeatures         `avp:"Supported-Features,omitempty"`
	UserId                      []datatype.UTF8String       `avp:"User-Id,omitempty"`
}

// RequestedEUTRANAuthInfo contains the information needed for authentication requests
// for E-UTRAN.
type RequestedEUTRANAuthInfo struct {
	NumVectors        datatype.Unsigned32  `avp:"Number-Of-Requested-Vectors,omitempty"`
	ImmediateResponse datatype.Unsigned32  `avp:"Immediate-Response-Preferred,omitempty"`
	ResyncInfo        datatype.OctetString `avp:"Re-synchronization-Info,omitempty"`
}

// RequestedUtranGeranAuthInfo contains the information needed for authentication requests
// for UTRAN/GERAN.
type RequestedUtranGeranAuthInfo struct {
	NumVectors        datatype.Unsigned32  `avp:"Number-Of-Requested-Vectors,omitempty"`
	ImmediateResponse datatype.Unsigned32  `avp:"Immediate-Response-Preferred,omitempty"`
	ResyncInfo        datatype.OctetString `avp:"Re-synchronization-Info,omitempty"`
}

// AIR encapsulates all of the information contained in an authentication information request.
// This information is sent to fetch data in order to authenticate a subscriber.
type AIR struct {
	SessionID                   datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost                  datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm                 datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState            datatype.UTF8String       `avp:"Auth-Session-State,omitempty"`
	UserName                    string                    `avp:"User-Name,omitempty"`
	VisitedPLMNID               datatype.OctetString      `avp:"Visited-PLMN-Id,omitempty"`
	RequestedEUTRANAuthInfo     *RequestedEUTRANAuthInfo  `avp:"Requested-EUTRAN-Authentication-Info,omitempty"`
	RequestedUtranGeranAuthInfo *RequestedEUTRANAuthInfo  `avp:"Requested-UTRAN-GERAN-Authentication-Info,omitempty"`
}

// ULR is an update location request. It is used to update location information in the HSS.
type ULR struct {
	SessionID         datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost        datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm       datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState  datatype.Unsigned32       `avp:"Auth-Session-State,omitempty"`
	UserName          datatype.UTF8String       `avp:"User-Name,omitempty"`
	VisitedPLMNID     datatype.OctetString      `avp:"Visited-PLMN-Id,omitempty"`
	RATType           datatype.Unsigned32       `avp:"RAT-Type,omitempty"`
	ULRFlags          datatype.Unsigned32       `avp:"ULR-Flags,omitempty"`
	SupportedFeatures []SupportedFeatures       `avp:"Supported-Features,omitempty"`
}
