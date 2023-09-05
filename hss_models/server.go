package hss_models

import (
	"gohss/modules/go-diameter/v4/diam/sm"

	"github.com/emakeev/milenage"
	"github.com/flosch/pongo2/v6"
)

type HomeSubscriberServer struct {
	Store    SubscriberStore
	Config   HSSConfig
	Milenage *milenage.Cipher
	SmClient *sm.Client
	// ConnMan        *diameter.ConnectionManager
	// RequestTracker *RequestTracker
	ClientMapping map[string]string

	// authSqnInd is an index used in the array scheme described by 3GPP TS 33.102 Appendix C.1.2 and C.2.2.
	// SQN consists of two parts (SQN = SEQ||IND).
	AuthSqnInd uint64
}

type HSSConfig struct {
	Name       string
	Server     *DiameterConfig
	AuthOp     []byte
	AuthAmf    []byte
	CxTemplate *pongo2.Template
	ShTemplate *pongo2.Template
}

func (srv *HomeSubscriberServer) GetVendorID() uint32 {
	return srv.Config.Server.Vendor3GPP
}
