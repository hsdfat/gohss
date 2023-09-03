package sh

import "gohss/modules/go-diameter/v4/diam/datatype"

type ExperimentalResult struct {
	VendorId               uint32 `avp:"Vendor-Id,omitempty"`
	ExperimentalResultCode uint32 `avp:"Experimental-Result-Code,omitempty"`
}

type SupportedFeatures struct {
	VendorId      uint32 `avp:"Vendor-Id,omitempty"`
	FeatureListID uint32 `avp:"Feature-List-ID,omitempty"`
	FeatureList   uint32 `avp:"Feature-List,omitempty"`
}

type UDR struct {
	SessionID         datatype.UTF8String       `avp:"Session-Id,omitempty"`
	OriginHost        datatype.DiameterIdentity `avp:"Origin-Host,omitempty"`
	OriginRealm       datatype.DiameterIdentity `avp:"Origin-Realm,omitempty"`
	AuthSessionState  *datatype.Enumerated      `avp:"Auth-Session-State,omitempty"`
	DestinationHost   datatype.DiameterIdentity `avp:"Destination-Host,omitempty"`
	DestinationRealm  datatype.DiameterIdentity `avp:"Destination-Realm,omitempty"`
	UserName          datatype.UTF8String       `avp:"User-Name,omitempty"`
	UserIdentity      datatype.UTF8String       `avp:"User-Identity,omitempty"`
	ServerName        datatype.UTF8String       `avp:"Server-Name,omitempty"`
	ServiceIndication datatype.OctetString      `avp:"Service-Indication,omitempty"`
	DataReference     *datatype.Enumerated      `avp:"Data-Reference,omitempty"`
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
	UserData           datatype.OctetString      `avp:"User-Data,omitempty"`
	SupportedFeatures  []*SupportedFeatures      `avp:"Supported-Features,omitempty"`
}
