package cxdx

import "gohss/modules/go-diameter/v4/diam/datatype"

type UAR struct {
	SessionID                datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost               datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm              datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState         *datatype.Enumerated      `avp:"Auth-Session-State,omitempty"`
	DestinationHost          datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm         datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName                 datatype.UTF8String       `avp:"User-Name,omitempty"`
	PublicIdentity           datatype.UTF8String       `avp:"Public-Identity,omitempty"`
	VisitedNetworkIdentifier datatype.OctetString      `avp:"Visited-Network-Identifier,omitempty"`
	UserAuthorizationType    *datatype.Enumerated      `avp:"User-Authorization-Type,omitempty"`
	UARFlags                 datatype.Unsigned32       `avp:"UAR-Flags,omitempty"`
	SupportedFeatures        datatype.UTF8String       `avp:"Supported-Features,omitempty"`
}

type UAA struct {
	SessionID          datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost         datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm        datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState   datatype.Enumerated       `avp:"Auth-Session-State,omitempty"`
	DestinationHost    datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm   datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName           datatype.UTF8String       `avp:"User-Name,omitempty"`
	ResultCode         datatype.Unsigned32       `avp:"Result-Code,omitempty"`
	ExperimentalResult *ExperimentalResult       `avp:"Experimental-Result,omitempty"`
	ServerName         datatype.UTF8String       `avp:"Server-Name,omitempty"`
	ServerCapabilities *ServerCapabilities       `avp:"Server-Capabilities,omitempty"`
	SupportedFeatures  []*SupportedFeatures      `avp:"Supported-Features,omitempty"`
}
type ServerCapabilities struct {
	MandatoryCapability datatype.Unsigned32 `avp:"Mandatory-Capability,omitempty"`
	OptionalCapability  datatype.Unsigned32 `avp:"Optional-Capability,omitempty"`
	ServerName          datatype.UTF8String `avp:"Server-Name,omitempty"`
}
type ExperimentalResult struct {
	VendorId               uint32 `avp:"Vendor-Id,omitempty"`
	ExperimentalResultCode uint32 `avp:"Experimental-Result-Code,omitempty"`
}

type SupportedFeatures struct {
	VendorId      uint32 `avp:"Vendor-Id,omitempty"`
	FeatureListID uint32 `avp:"Feature-List-ID,omitempty"`
	FeatureList   uint32 `avp:"Feature-List,omitempty"`
}

type SAR struct {
	SessionID                      datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost                     datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm                    datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState               *datatype.Enumerated      `avp:"Auth-Session-State,omitempty"`
	DestinationHost                datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm               datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName                       datatype.UTF8String       `avp:"User-Name,omitempty"`
	ServerName                     datatype.UTF8String       `avp:"Server-Name,omitempty"`
	UserDataAlreadyAvailable       *datatype.Enumerated      `avp:"User-Data-Already-Available,omitempty"`
	SCSCFRestorationInfo           *SCSCFRestorationInfo     `avp:"SCSCF-Restoration-Info,omitempty"`
	MultipleRegistrationIndication *datatype.Enumerated      `avp:"Multiple-Registration-Indication,omitempty"`
	SARFlag                        datatype.Unsigned32       `avp:"SAR-Flags,omitempty"`
	VisitedNetworkIdentifier       datatype.OctetString      `avp:"Visited-Network-Identifier,omitempty"`
	SupportedFeatures              []*SupportedFeatures      `avp:"Supported-Features,omitempty"`
}
type SCSCFRestorationInfo struct {
	UserName                datatype.UTF8String `avp:"User-Name,omitempty"`
	RestorationInfo         *RestorationInfo    `avp:"Restoration-Info,omitempty"`
	SIPAuthenticationScheme datatype.UTF8String `avp:"SIP-Authentication-Scheme,omitempty"`
}

type RestorationInfo struct {
	Path                      datatype.OctetString `avp:"Path,omitempty"`
	Contact                   datatype.OctetString `avp:"Contact,omitempty"`
	InitialCSeqSequenceNumber datatype.Unsigned32  `avp:"Initial-CSeq-Sequence-Number,omitempty"`
	CallIDSIPHeader           datatype.OctetString `avp:"Call-ID-SIP-Header,omitempty"`
	SubscriptionInfo          *SubscriptionInfo    `avp:"Subscription-Info,omitempty"`
}

type SubscriptionInfo struct {
	CallIDSIPHeader datatype.OctetString `avp:"Call-ID-SIP-Header,omitempty"`
	FromSIPHeader   datatype.OctetString `avp:"From-SIP-Header,omitempty"`
	ToSIPHeader     datatype.OctetString `avp:"To-SIP-Header,omitempty"`
	RecordRoute     datatype.OctetString `avp:"Record-Route,omitempty"`
	Contact         datatype.OctetString `avp:"Contact,omitempty"`
}

type SAA struct {
	SessionID            datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost           datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm          datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState     *datatype.Enumerated      `avp:"Auth-Session-State,omitempty"`
	DestinationHost      datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm     datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName             datatype.UTF8String       `avp:"User-Name,omitempty"`
	ResultCode           datatype.Unsigned32       `avp:"Result-Code,omitempty"`
	UserData             datatype.OctetString      `avp:"User-Data,omitempty"`
	ChargingInformation  *ChargingInformation      `avp:"Charging-Information,omitempty"`
	AssociatedIdentities *AssociatedIdentities     `avp:"Associated-Identities,omitempty"`
	SCSCFRestorationInfo *SCSCFRestorationInfo     `avp:"SCSCF-Restoration-Info,omitempty"`
	ServerName           datatype.UTF8String       `avp:"Server-Name,omitempty"`
	LooseRouteIndication *datatype.Enumerated      `avp:"Loose-Route-Indication,omitempty"`
	ExperimentalResult   *ExperimentalResult       `avp:"Experimental-Result,omitempty"`
	SupportedFeatures    []*SupportedFeatures      `avp:"Supported-Features,omitempty"`
}

type ChargingInformation struct {
	PrimaryEventChargingFunctionName        datatype.DiameterURI `avp:"Primary-Event-Charging-Function-Name,omitempty"`
	SecondaryEventChargingFunctionName      datatype.DiameterURI `avp:"Secondary-Event-Charging-Function-Name,omitempty"`
	PrimaryChargingCollectionFunctionName   datatype.DiameterURI `avp:"Primary-Charging-CollectionFunction-Name,omitempty"`
	SecondaryChargingCollectionFunctionName datatype.DiameterURI `avp:"Secondary-Charging-CollectionFunction-Name,omitempty"`
}

type AssociatedIdentities struct {
	UserName datatype.UTF8String `avp:"User-Name,omitempty"`
}

type LIR struct {
	SessionID                datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost               datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm              datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState         *datatype.Enumerated      `avp:"Auth-Session-State,omitempty"`
	DestinationHost          datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm         datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName                 datatype.UTF8String       `avp:"User-Name,omitempty"`
	PublicIdentity           datatype.UTF8String       `avp:"Public-Identity,omitempty"`
	SessionPriority          *datatype.Enumerated      `avp:"Session-Priority,omitempty"`
	VisitedNetworkIdentifier *datatype.Enumerated      `avp:"Visited-Network-Identifier,omitempty"`
	UserAuthorizationType    *datatype.Enumerated      `avp:"User-Authorization-Type,omitempty"`
	SupportedFeatures        []*SupportedFeatures      `avp:"Supported-Features,omitempty"`
}

type LIA struct {
	SessionID          datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost         datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm        datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState   *datatype.Enumerated      `avp:"Auth-Session-State,omitempty"`
	DestinationHost    datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm   datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName           datatype.UTF8String       `avp:"User-Name,omitempty"`
	ServerName         datatype.UTF8String       `avp:"Server-Name,omitempty"`
	ServerCapabilities *ServerCapabilities       `avp:"Server-Capabilities,omitempty"`
	ResultCode         datatype.Unsigned32       `avp:"Result-Code,omitempty"`
	ExperimentalResult *ExperimentalResult       `avp:"Experimental-Result,omitempty"`
	SupportedFeatures  []*SupportedFeatures      `avp:"Supported-Features,omitempty"`
	LIAFlags           datatype.Unsigned32       `avp:"LIA-Flags,omitempty"`
}

type MAR struct {
	SessionID        datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost       datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm      datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState *datatype.Enumerated      `avp:"Auth-Session-State,omitempty"`
	DestinationHost  datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName         datatype.UTF8String       `avp:"User-Name,omitempty"`
	PublicIdentity   datatype.UTF8String       `avp:"Public-Identity,omitempty"`
	// RATType                  *datatype.Enumerated      `avp:"RAT-Type,omitempty"`
	VisitedNetworkIdentifier *datatype.Enumerated `avp:"Visited-Network-Identifier,omitempty"`
	ServerName               datatype.UTF8String  `avp:"Server-Name,omitempty"`
	SIPAuthDataItem          *SIPAuthDataItem     `avp:"SIP-Auth-Data-Item,omitempty"`
	SIPNumberAuthItems       datatype.Unsigned32  `avp:"SIP-Number-Auth-Items,omitempty"`
	SupportedFeatures        []*SupportedFeatures `avp:"Supported-Features,omitempty"`
}

type MAA struct {
	SessionID                datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost               datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm              datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState         *datatype.Enumerated      `avp:"Auth-Session-State,omitempty"`
	DestinationHost          datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm         datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName                 datatype.UTF8String       `avp:"User-Name,omitempty"`
	PublicIdentity           datatype.UTF8String       `avp:"Public-Identity,omitempty"`
	RATType                  *datatype.Enumerated      `avp:"RAT-Type,omitempty"`
	VisitedNetworkIdentifier *datatype.Enumerated      `avp:"Visited-Network-Identifier,omitempty"`
	ServerName               datatype.UTF8String       `avp:"Server-Name,omitempty"`
	SIPAuthDataItem          []*SIPAuthDataItem        `avp:"SIP-Auth-Data-Item,omitempty"`
	SIPNumberAuthItems       datatype.Unsigned32       `avp:"SIP-Number-Auth-Items,omitempty"`
	SupportedFeatures        []*SupportedFeatures      `avp:"Supported-Features,omitempty"`
}

type SIPAuthDataItem struct {
	SIPItemNumber           datatype.Unsigned32    `avp:"SIP-Item-Number,omitempty"`
	SIPAuthenticationScheme datatype.UTF8String    `avp:"SIP-Authentication-Scheme,omitempty"`
	SIPAuthencate           datatype.OctetString   `avp:"SIP-Authenticate,omitempty"`
	SIPAuthorization        datatype.OctetString   `avp:"SIP-Authorization,omitempty"`
	SIPAuthorizationContext datatype.OctetString   `avp:"SIP-Authorization-Context,omitempty"`
	ConfidentialityKey      datatype.OctetString   `avp:"Confidentiality-Key,omitempty"`
	IntegrityKey            datatype.OctetString   `avp:"Integrity-Key,omitempty"`
	SIPDigestAuthenticate   *SIPDigestAuthenticate `avp:"SIP-Digest-Authenticate,omitempty"`
}

type SIPDigestAuthenticate struct {
	DigestRealm     datatype.UTF8String `avp:"Digest-Realm,omitempty"`
	DigestAlgorithm datatype.UTF8String `avp:"Digest-Algorithm,omitempty"`
	DigestQoP       datatype.UTF8String `avp:"Digest-QoP,omitempty"`
	DigestHA1       datatype.UTF8String `avp:"Digest-HA1,omitempty"`
}
