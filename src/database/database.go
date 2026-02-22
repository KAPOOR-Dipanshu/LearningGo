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

// GetAllEmployees retrieves all employees from the database.
func GetAllEmployees(ctx context.Context) ([]models.Employee, error) {
	var results []models.Employee
	cursor, err := GetCollection().Find(ctx, bson.M{})
	if err != nil {
		log.Printf("error finding employees: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &results); err != nil {
		log.Printf("error decoding employees: %v", err)
		return nil, err
	}
	return results, nil
}

// GetEmployeeByID retrieves an employee by their custom id.
func GetEmployeeByID(ctx context.Context, id int) (*models.Employee, error) {
	var result models.Employee
	filter := bson.M{"id": id}
	err := GetCollection().FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Printf("error finding employee by id %d: %v", id, err)
		return nil, err
	}
	return &result, nil
}

// InsertEmployee inserts a new employee and auto-assigns a unique id.
func InsertEmployee(ctx context.Context, emp models.Employee) (int32, error) {
	currentCount, err := GetCollection().CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("error counting employees: %v", err)
		return 0, err
	}
	emp.ID = int32(currentCount) + 1
	_, err = GetCollection().InsertOne(ctx, emp)
	if err != nil {
		log.Printf("error inserting employee: %v", err)
		return 0, err
	}
	return emp.ID, nil
}

// UpdateEmployee updates an employee's information by id.
func UpdateEmployee(ctx context.Context, id int, emp models.Employee) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": emp}
	_, err := GetCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("error updating employee id %d: %v", id, err)
	}
	return err
}

// DeleteEmployee removes an employee by id.
func DeleteEmployee(ctx context.Context, id int) error {
	filter := bson.M{"id": id}
	_, err := GetCollection().DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("error deleting employee id %d: %v", id, err)
	}
	return err
}
