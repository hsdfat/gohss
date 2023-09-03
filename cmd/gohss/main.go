package main

import (
	"gohss/hss_models"
	"gohss/server"
	"gohss/store/mongoDb"
)

func main() {
	config := &hss_models.HSSConfig{
		Server: &hss_models.DiameterConfig{
			Protocol:     "tcp",
			Address:      "0.0.0.0:3768",
			LocalAddress: "0.0.0.0:3768",
			DestHost:     "hss",
			DestRealm:    "hss",
			Vendor3GPP:   10415,
		},
		AuthOp:  []byte{},
		AuthAmf: []byte{0x80, 0x00},
		Name:    "gohss",
	}
	// store := models.SubscriberStore{}
	db := &mongoDb.MongoDBClient{
		DBName: "free5gc",
		URL:    "mongodb://admin:123@localhost:27017/free5gc?authsource=admin",
	}
	hssServer := server.SetupHssServer(config, db)
	// endless loop
	server.StartDiameterServer(hssServer)
}
