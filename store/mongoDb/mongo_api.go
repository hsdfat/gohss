package mongoDb

import (
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func (c *MongoDBClient) QueryAuthSubsData(collName string, ueId string) (*models.AuthenticationSubscription, error) {
// 	filter := bson.M{"ueId": ueId}

// 	result := MongoDBLibrary.RestfulAPIGetOneCollection(collName, filter)
// 	authenSubs := &models.AuthenticationSubscription{}
// 	err := result.Decode(authenSubs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return authenSubs, nil
// }

func (c *MongoDBClient) GetOne(collName string, filter bson.M) *mongo.SingleResult {

	collection := c.client.Database(c.DBName).Collection(collName)

	return collection.FindOne(context.TODO(), filter)
}

func (c *MongoDBClient) PutOne(collName string, filter bson.M, putData map[string]interface{}) error {

	collection := c.client.Database(c.DBName).Collection(collName)
	var checkItem map[string]interface{}
	collection.FindOne(context.TODO(), filter).Decode(&checkItem)

	if checkItem == nil {
		_, err := collection.InsertOne(context.TODO(), putData)
		return err
	} else {
		_, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": putData})
		return err
	}
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
