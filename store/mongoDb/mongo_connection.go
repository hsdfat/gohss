package mongoDb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/omec-project/MongoDBLibrary"
	"github.com/omec-project/MongoDBLibrary/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBClient struct {
	URL    string
	client *mongo.Client
	DBName string
	pools  map[string]map[string]int
}

func (c *MongoDBClient) SetMongoDB(setdbName string, url string) error {
	log.Println("Init database", setdbName, url)
	if c.client != nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	defer cancel()
	if err != nil {
		//defer cancel()
		return err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	filter := bson.D{{}}
	fmt.Println(client.ListDatabases(context.TODO(), filter))
	c.client = client
	c.DBName = setdbName
	return nil
}

func (c *MongoDBClient) InitializeChunkPool(poolName string, min int, max int, retries int, chunkSize int) {
	logger.MongoDBLog.Println("ENTERING InitializeChunkPool")
	var poolData = map[string]int{}
	poolData["min"] = min
	poolData["max"] = max
	poolData["retries"] = retries
	poolData["chunkSize"] = chunkSize

	c.pools[poolName] = poolData
	logger.MongoDBLog.Println("Pools: ", c.pools)
}

func (c *MongoDBClient) getDataFromDB(collName string, filter bson.M) (map[string]interface{}, error) {
	data := MongoDBLibrary.RestfulAPIGetOne(collName, filter)
	if data == nil {
		return nil, fmt.Errorf("no data found")
	}

	// Delete "_id" entry which is auto-inserted by MongoDB
	delete(data, "_id")
	return data, nil
}

func (c *MongoDBClient) deleteDataFromDB(collName string, filter bson.M) {
	MongoDBLibrary.RestfulAPIDeleteOne(collName, filter)
}

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
