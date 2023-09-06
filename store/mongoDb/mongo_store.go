package mongoDb

import (
	"fmt"
	"strings"

	"gohss/modules/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *MongoDBClient) Init() error {
	return c.SetMongoDB(c.DBName, c.URL)
}

func (c *MongoDBClient) GetAuthenSubscriptionData(id string) (subs *models.AuthenticationSubscription, err error) {
	if !strings.HasPrefix(id, "imsi-") {
		id = fmt.Sprintf("imsi-%s", id)
	}
	filter := bson.M{"ueId": id}
	subs = &models.AuthenticationSubscription{}
	result := c.GetOne(AuthenticationSubscriptionColName, filter)
	err = result.Decode(subs)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (c *MongoDBClient) GetSMSubscriptionData(id string) (subs *models.SessionManagementSubscriptionData, err error) {
	if !strings.HasPrefix(id, "imsi-") {
		id = fmt.Sprintf("imsi-%s", id)
	}
	filter := bson.M{"ueId": id}
	subs = &models.SessionManagementSubscriptionData{}
	result := c.GetOne(SessionManagementSubscriptionDataColName, filter)
	err = result.Decode(subs)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (c *MongoDBClient) GetAMSubscriptionData(id string) (subs *models.AccessAndMobilitySubscriptionData, err error) {
	if !strings.HasPrefix(id, "imsi-") {
		id = fmt.Sprintf("imsi-%s", id)
	}
	filter := bson.M{"ueId": id}
	subs = &models.AccessAndMobilitySubscriptionData{}
	result := c.GetOne(AccessAndMobilitySubscriptionColName, filter)
	err = result.Decode(subs)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (c *MongoDBClient) UpdateAuthenSubscriptionData(id string, subs *models.AuthenticationSubscription) error {
	if !strings.HasPrefix(id, "imsi-") {
		id = fmt.Sprintf("imsi-%s", id)
	}
	filter := bson.M{"ueId": id}
	putData := toBsonM(subs)
	putData["ueId"] = id

	return c.PutOne(AuthenticationSubscriptionColName, filter, putData)
}
