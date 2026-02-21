package database

import (
	"context"
	"go-api-app/src/constants"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
	once       sync.Once
)

func InitMongo() {
	once.Do(func() {
		MONGO_URI := constants.GetConstant("MONGO_URI")
		DATABASE := constants.GetConstant("DATABASE_NAME")
		COLLECTION := constants.GetConstant("COLLECTION_NAME")

		clientOptions := options.Client().ApplyURI(MONGO_URI)
		var err error
		client, err = mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Connected to mongoDB!!!")
		}

		collection = client.Database(DATABASE).Collection(COLLECTION)
	})
}

func GetCollection() *mongo.Collection {
	if collection == nil {
		InitMongo()
	}
	return collection
}

func GetAllData() ([]map[string]any, error) {
    var results []map[string]any
    cursor, err := GetCollection().Find(context.Background(), map[string]any{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    if err = cursor.All(context.Background(), &results); err != nil {
        return nil, err
    }
    return results, nil
}

func GetDataByID(id string) (map[string]any, error) {
    var result map[string]any
    filter := map[string]any{"_id": id}
    err := GetCollection().FindOne(context.Background(), filter).Decode(&result)
    if err != nil {
        return nil, err
    }
    return result, nil
}