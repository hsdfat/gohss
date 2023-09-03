package main

import (
	"encoding/json"
	"log"

	"github.com/omec-project/MongoDBLibrary"
	"github.com/omec-project/openapi/models"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	url := "mongodb://admin:123@localhost:27017/free5gc?authsource=admin"
	MongoDBLibrary.SetMongoDB("free5gc", url)
	InsertAuthSubscriptionToMongoDB(ueId, authSubsData)
	InsertAccessAndMobilitySubscriptionDataToMongoDB(ueId, amDataData, "45204")
	InsertSessionManagementSubscriptionDataToMongoDB(ueId, "45204", smDataData)
}

var (
	authSubsData  models.AuthenticationSubscription
	amDataData    models.AccessAndMobilitySubscriptionData
	smDataData    []models.SessionManagementSubscriptionData
	smfSelData    models.SmfSelectionSubscriptionData
	amPolicyData  models.AmPolicyData
	smPolicyData  models.SmPolicyData
	servingPlmnId string
	ueId          string
)

func init() {

	authSubsData = models.AuthenticationSubscription{
		AuthenticationManagementField: "8000",
		AuthenticationMethod:          "5G_AKA", // "5G_AKA", "EAP_AKA_PRIME"
		Milenage: &models.Milenage{
			Op: &models.Op{
				EncryptionAlgorithm: 0,
				EncryptionKey:       0,
				OpValue:             "c9e8763286b5b9ffbdf56e1297d0887b", // Required
			},
		},
		Opc: &models.Opc{
			EncryptionAlgorithm: 0,
			EncryptionKey:       0,
			OpcValue:            "981d464c7c52eb6e5036234984ad0bcf", // Required
		},
		PermanentKey: &models.PermanentKey{
			EncryptionAlgorithm: 0,
			EncryptionKey:       0,
			PermanentKeyValue:   "5122250214c33e723a5dd523fc145fc0", // Required
		},
		SequenceNumber: "16f3b3f70fc2",
	}

	amDataData = models.AccessAndMobilitySubscriptionData{
		Gpsis: []string{
			"msisdn-0900000000",
		},
		Nssai: &models.Nssai{
			DefaultSingleNssais: []models.Snssai{
				{
					Sd:  "010203",
					Sst: 1,
				},
				{
					Sd:  "112233",
					Sst: 1,
				},
			},
			SingleNssais: []models.Snssai{
				{
					Sd:  "010203",
					Sst: 1,
				},
				{
					Sd:  "112233",
					Sst: 1,
				},
			},
		},
		SubscribedUeAmbr: &models.AmbrRm{
			Downlink: "1000 Kbps",
			Uplink:   "1000 Kbps",
		},
	}

	smDataData = []models.SessionManagementSubscriptionData{
		{
			SingleNssai: &models.Snssai{
				Sst: 1,
				Sd:  "010203",
			},
			DnnConfigurations: map[string]models.DnnConfiguration{
				"internet": {
					PduSessionTypes: &models.PduSessionTypes{
						DefaultSessionType:  models.PduSessionType_IPV4,
						AllowedSessionTypes: []models.PduSessionType{models.PduSessionType_IPV4},
					},
					SscModes: &models.SscModes{
						DefaultSscMode:  models.SscMode__1,
						AllowedSscModes: []models.SscMode{models.SscMode__1},
					},
					SessionAmbr: &models.Ambr{
						Downlink: "1000 Kbps",
						Uplink:   "1000 Kbps",
					},
					Var5gQosProfile: &models.SubscribedDefaultQos{
						Var5qi: 9,
						Arp: &models.Arp{
							PriorityLevel: 8,
						},
						PriorityLevel: 8,
					},
				},
			},
		},
		{
			SingleNssai: &models.Snssai{
				Sst: 1,
				Sd:  "112233",
			},
			DnnConfigurations: map[string]models.DnnConfiguration{
				"internet": {
					PduSessionTypes: &models.PduSessionTypes{
						DefaultSessionType:  models.PduSessionType_IPV4,
						AllowedSessionTypes: []models.PduSessionType{models.PduSessionType_IPV4},
					},
					SscModes: &models.SscModes{
						DefaultSscMode:  models.SscMode__1,
						AllowedSscModes: []models.SscMode{models.SscMode__1},
					},
					SessionAmbr: &models.Ambr{
						Downlink: "2000 Kbps",
						Uplink:   "1000 Kbps",
					},
					Var5gQosProfile: &models.SubscribedDefaultQos{
						Var5qi: 9,
						Arp: &models.Arp{
							PriorityLevel: 8,
							PreemptCap:    models.PreemptionCapability_MAY_PREEMPT,
						},
						PriorityLevel: 8,
					},
				},
			},
		},
	}

	smfSelData = models.SmfSelectionSubscriptionData{
		SubscribedSnssaiInfos: map[string]models.SnssaiInfo{
			"01010203": {
				DnnInfos: []models.DnnInfo{
					{
						Dnn: "internet",
					},
				},
			},
			"01112233": {
				DnnInfos: []models.DnnInfo{
					{
						Dnn: "internet",
					},
				},
			},
		},
	}

	amPolicyData = models.AmPolicyData{
		SubscCats: []string{
			"free5gc",
		},
	}

	smPolicyData = models.SmPolicyData{
		SmPolicySnssaiData: map[string]models.SmPolicySnssaiData{
			"01010203": {
				Snssai: &models.Snssai{
					Sd:  "010203",
					Sst: 1,
				},
				SmPolicyDnnData: map[string]models.SmPolicyDnnData{
					"internet": {
						Dnn: "internet",
					},
				},
			},
			"01112233": {
				Snssai: &models.Snssai{
					Sd:  "112233",
					Sst: 1,
				},
				SmPolicyDnnData: map[string]models.SmPolicyDnnData{
					"internet": {
						Dnn: "internet",
					},
				},
			},
		},
	}

	servingPlmnId = "45204"
	ueId = "imsi-452040000000022"
}

func InsertAuthSubscriptionToMongoDB(ueId string, authSubs models.AuthenticationSubscription) {

	collName := "subscriptionData.authenticationData.authenticationSubscription"
	filter := bson.M{"ueId": ueId}
	putData := toBsonM(authSubs)
	putData["ueId"] = ueId
	MongoDBLibrary.RestfulAPIPutOne(collName, filter, putData)
}

func toBsonM(data interface{}) bson.M {
	tmp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Marshal error in toBsonM: %+v", err)
	}
	var putData = bson.M{}
	err = json.Unmarshal(tmp, &putData)
	if err != nil {
		log.Fatalf("Unmarshal error in toBsonM: %+v", err)
	}
	return putData
}

func DelAuthSubscriptionToMongoDB(ueId string) {
	collName := "subscriptionData.authenticationData.authenticationSubscription"
	filter := bson.M{"ueId": ueId}
	MongoDBLibrary.RestfulAPIDeleteMany(collName, filter)
}

func InsertAccessAndMobilitySubscriptionDataToMongoDB(
	ueId string, amData models.AccessAndMobilitySubscriptionData, servingPlmnId string) {
	collName := "subscriptionData.provisionedData.amData"
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	putData := toBsonM(amData)
	putData["ueId"] = ueId
	putData["servingPlmnId"] = servingPlmnId
	MongoDBLibrary.RestfulAPIPutOne(collName, filter, putData)
}

func GetAccessAndMobilitySubscriptionDataFromMongoDB(
	ueId string, servingPlmnId string) (amData *models.AccessAndMobilitySubscriptionData) {
	collName := "subscriptionData.provisionedData.amData"
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	getData := MongoDBLibrary.RestfulAPIGetOne(collName, filter)
	if getData == nil {
		return
	}
	tmp, err := json.Marshal(getData)
	if err != nil {
		return
	}
	amData = new(models.AccessAndMobilitySubscriptionData)
	err = json.Unmarshal(tmp, amData)
	if err != nil {
		log.Fatalf("Unmarshal error in GetAccessAndMobilitySubscriptionDataFromMongoDB: %+v", err)
	}
	return
}

func DelAccessAndMobilitySubscriptionDataFromMongoDB(ueId string, servingPlmnId string) {
	collName := "subscriptionData.provisionedData.amData"
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	MongoDBLibrary.RestfulAPIDeleteMany(collName, filter)
}

func InsertSessionManagementSubscriptionDataToMongoDB(
	ueId string, servingPlmnId string, smDatas []models.SessionManagementSubscriptionData) {
	var putDatas = make([]interface{}, 0, len(smDatas))
	collName := "subscriptionData.provisionedData.smData"
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	for _, smData := range smDatas {
		putData := toBsonM(smData)
		putData["ueId"] = ueId
		putData["servingPlmnId"] = servingPlmnId
		putDatas = append(putDatas, putData)
	}
	MongoDBLibrary.RestfulAPIPostMany(collName, filter, putDatas)
}

func GetSessionManagementDataFromMongoDB(
	ueId string, servingPlmnId string) (amData *models.SessionManagementSubscriptionData) {
	collName := "subscriptionData.provisionedData.smData"
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	getData := MongoDBLibrary.RestfulAPIGetOne(collName, filter)
	if getData == nil {
		return
	}
	tmp, err := json.Marshal(getData)
	if err != nil {
		return
	}
	amData = new(models.SessionManagementSubscriptionData)
	err = json.Unmarshal(tmp, amData)
	if err != nil {
		log.Fatalf("Unmarshal error in GetSessionManagementDataFromMongoDB: %+v", err)
	}
	return
}

func DelSessionManagementSubscriptionDataFromMongoDB(ueId string, servingPlmnId string) {
	collName := "subscriptionData.provisionedData.smData"
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	MongoDBLibrary.RestfulAPIDeleteMany(collName, filter)
}

func InsertSmfSelectionSubscriptionDataToMongoDB(
	ueId string, smfSelData models.SmfSelectionSubscriptionData, servingPlmnId string) {
	collName := "subscriptionData.provisionedData.smfSelectionSubscriptionData"
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	putData := toBsonM(smfSelData)
	putData["ueId"] = ueId
	putData["servingPlmnId"] = servingPlmnId
	MongoDBLibrary.RestfulAPIPutOne(collName, filter, putData)
}

func GetSmfSelectionSubscriptionDataFromMongoDB(
	ueId string, servingPlmnId string) (smfSelData *models.SmfSelectionSubscriptionData) {
	collName := "subscriptionData.provisionedData.smfSelectionSubscriptionData"
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	getData := MongoDBLibrary.RestfulAPIGetOne(collName, filter)
	if getData == nil {
		return
	}
	tmp, err := json.Marshal(getData)
	if err != nil {
		return
	}
	smfSelData = new(models.SmfSelectionSubscriptionData)
	err = json.Unmarshal(tmp, smfSelData)
	if err != nil {
		log.Fatalf("Unmarshal error in GetSmfSelectionSubscriptionDataFromMongoDB: %+v", err)
	}
	return
}

func DelSmfSelectionSubscriptionDataFromMongoDB(ueId string, servingPlmnId string) {
	collName := "subscriptionData.provisionedData.smfSelectionSubscriptionData"
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	MongoDBLibrary.RestfulAPIDeleteMany(collName, filter)
}

func InsertAmPolicyDataToMongoDB(ueId string, amPolicyData models.AmPolicyData) {
	collName := "policyData.ues.amData"
	filter := bson.M{"ueId": ueId}
	putData := toBsonM(amPolicyData)
	putData["ueId"] = ueId
	MongoDBLibrary.RestfulAPIPutOne(collName, filter, putData)
}

func GetAmPolicyDataFromMongoDB(ueId string) (amPolicyData *models.AmPolicyData) {
	collName := "policyData.ues.amData"
	filter := bson.M{"ueId": ueId}
	getData := MongoDBLibrary.RestfulAPIGetOne(collName, filter)
	if getData == nil {
		return
	}
	tmp, err := json.Marshal(getData)
	if err != nil {
		return
	}
	amPolicyData = new(models.AmPolicyData)
	err = json.Unmarshal(tmp, amPolicyData)
	if err != nil {
		log.Fatalf("Unmarshal error in GetAmPolicyDataFromMongoDB: %+v", err)
	}
	return
}

func DelAmPolicyDataFromMongoDB(ueId string) {
	collName := "policyData.ues.amData"
	filter := bson.M{"ueId": ueId}
	MongoDBLibrary.RestfulAPIDeleteMany(collName, filter)
}

func InsertSmPolicyDataToMongoDB(ueId string, smPolicyData models.SmPolicyData) {
	collName := "policyData.ues.smData"
	filter := bson.M{"ueId": ueId}
	putData := toBsonM(smPolicyData)
	putData["ueId"] = ueId
	MongoDBLibrary.RestfulAPIPutOne(collName, filter, putData)
}

func GetSmPolicyDataFromMongoDB(ueId string) (smPolicyData *models.SmPolicyData) {
	collName := "policyData.ues.smData"
	filter := bson.M{"ueId": ueId}
	getData := MongoDBLibrary.RestfulAPIGetOne(collName, filter)
	if getData == nil {
		return
	}
	tmp, err := json.Marshal(getData)
	if err != nil {
		return
	}
	smPolicyData = new(models.SmPolicyData)
	err = json.Unmarshal(tmp, smPolicyData)
	if err != nil {
		log.Fatalf("Unmarshal error in GetSmPolicyDataFromMongoDB: %+v", err)
	}
	return
}

func DelSmPolicyDataFromMongoDB(ueId string) {
	collName := "policyData.ues.smData"
	filter := bson.M{"ueId": ueId}
	MongoDBLibrary.RestfulAPIDeleteMany(collName, filter)
}
