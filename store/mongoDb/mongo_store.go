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

func (c *MongoDBClient) GetLTESubscriptionData(id string) (interface{}, error) {
	filter := bson.M{"ueId": id}
	result := c.GetOne(LTEsubscriptionColName, filter)
	var subs interface{}
	err := result.Decode(&subs)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

func (c *MongoDBClient) GetSIPServerName(publicId string) (string, error) {
	filter := bson.M{"publicId": publicId}
	result := c.GetOne(PublicIdColName, filter)
	var pubsubs models.PublicId
	err := result.Decode(&pubsubs)
	if err != nil {
		return "", err
	}
	filter = bson.M{"privateId": pubsubs.PrivateId}
	result = c.GetOne(PrivateIdColName, filter)
	var prisubs models.PrivateId
	err = result.Decode(&prisubs)
	if err != nil {
		return "", err
	}
	return prisubs.ServerName, nil
}

func (c *MongoDBClient) GetSIPImsi(publicId string) (string, error) {
	filter := bson.M{"publicId": publicId}
	result := c.GetOne(PublicIdColName, filter)
	var pubsubs models.PublicId
	err := result.Decode(&pubsubs)
	if err != nil {
		return "", err
	}
	filter = bson.M{"privateId": pubsubs.PrivateId}
	result = c.GetOne(PrivateIdColName, filter)
	var prisubs models.PrivateId
	err = result.Decode(&prisubs)
	if err != nil {
		return "", err
	}
	return prisubs.UeId, nil
}

func (c *MongoDBClient) GetSIPPrivateId(publicId string) (string, error) {
	filter := bson.M{"publicId": publicId}
	result := c.GetOne(PublicIdColName, filter)
	var pubsubs models.PublicId
	err := result.Decode(&pubsubs)
	if err != nil {
		return "", err
	}

	return pubsubs.PrivateId, nil
}

func (c *MongoDBClient) UpdateSIPServerName(publicId string, serverName string) error {
	filter := bson.M{"publicId": publicId}
	fmt.Println(publicId)
	result := c.GetOne(PublicIdColName, filter)
	var pubsubs models.PublicId
	err := result.Decode(&pubsubs)
	if err != nil {
		return err
	}
	filter = bson.M{"privateId": pubsubs.PrivateId}
	subs := models.PrivateId{
		ID:         pubsubs.PrivateId,
		ServerName: serverName,
	}
	putData := toBsonM(subs)

	return c.PutOne(PrivateIdColName, filter, putData)
}

func (c *MongoDBClient) GetSIPState(publicId string) (int, error) {
	filter := bson.M{"publicId": publicId}
	result := c.GetOne(PublicIdColName, filter)
	var pubsubs models.PublicId
	err := result.Decode(&pubsubs)
	if err != nil {
		return 0, err
	}
	return pubsubs.State, nil
}

func (c *MongoDBClient) UpdateSIPState(publicId string, state int) error {
	filter := bson.M{"publicId": publicId}
	subs := models.PublicId{
		ID:    publicId,
		State: state,
	}
	putData := toBsonM(subs)
	return c.PutOne(PrivateIdColName, filter, putData)
}
