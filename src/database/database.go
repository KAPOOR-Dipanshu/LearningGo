package database

import (
	"context"
	"go-api-app/src/constants"
	"go-api-app/src/models"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
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

func GetAllEmployees() ([]models.Employee, error) {
	var results []models.Employee
	cursor, err := GetCollection().Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

func GetEmployeeByID(id int) (*models.Employee, error) {
	var result models.Employee
	filter := bson.M{"id": id}
	err := GetCollection().FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func InsertEmployee(emp models.Employee) (int32, error) {
	currentCount, err := GetCollection().CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}
	emp.ID = int32(currentCount) + 1
	_, err = GetCollection().InsertOne(context.Background(), emp)
	if err != nil {
		return 0, err
	}
	return emp.ID, nil
}

func UpdateEmployee(id int, emp models.Employee) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": emp}
	_, err := GetCollection().UpdateOne(context.Background(), filter, update)
	return err
}

func DeleteEmployee(id int) error {
	filter := bson.M{"id": id}
	_, err := GetCollection().DeleteOne(context.Background(), filter)
	return err
}
