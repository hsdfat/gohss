package hss_models

type DiameterConfig struct {
	Protocol     string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`                             // tcp/sctp/...
	Address      string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`                               // host:port
	LocalAddress string `protobuf:"bytes,3,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"` // IP:port or :port
	DestHost     string `protobuf:"bytes,4,opt,name=dest_host,json=destHost,proto3" json:"dest_host,omitempty"`             // diameter host
	DestRealm    string `protobuf:"bytes,5,opt,name=dest_realm,json=destRealm,proto3" json:"dest_realm,omitempty"`          // diameter realm
	Vendor3GPP   uint32
}

type DiameterServerConnConfig struct {
	Addr      string // host:port
	Protocol  string // tcp/sctp
	LocalAddr string // IP:port or :port
}

type DiameterServerConfig struct {
	DiameterServerConnConfig
	DestHost          string
	DestRealm         string
	DisableDestHost   bool
	OverwriteDestHost bool
}

// DiameterClientConfig holds information for connecting with a diameter server
type DiameterClientConfig struct {
	Host               string // diameter host
	Realm              string // diameter realm
	ProductName        string
	AppID              uint32
	AuthAppID          uint32
	Retransmits        uint
	WatchdogInterval   uint
	RetryCount         uint // number of times to reconnect after connection lost
	RequestTimeout     uint
	SupportedVendorIDs string
	ServiceContextId   string
}
