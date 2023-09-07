package unittest

import (
	"gohss/store/mongoDb"
	"testing"
)

func TestUpdateServerName(t *testing.T) {
	// privateId := "sip@452.04"
	publicId := "sip@++452.04"
	// ueId := "imsi-452040000000022"

	db := &mongoDb.MongoDBClient{
		DBName: "free5gc",
		URL:    "mongodb://admin:123@localhost:27017/free5gc?authsource=admin",
	}
	db.Init()
	db.UpdateSIPServerName(publicId, "sp01")
	str, err := db.GetSIPServerName(publicId)
	if err != nil {
		t.Fail()
	}
	t.Log(str)
}
