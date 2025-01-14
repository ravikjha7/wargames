package helpers

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetMongoClient - Return mongodb connection to work with
func getMongoClient(MONGO_URI string) (*mongo.Client, error) {

	/* Used to create a singleton object of MongoDB client.
	Initialized and exposed through  GetMongoClient().*/
	var clientInstance *mongo.Client

	//Used during creation of singleton client object in GetMongoClient().
	var clientInstanceError error

	//Used to execute client creation procedure only once.
	var mongoOnce sync.Once

	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(MONGO_URI)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func InitDatabase() *mongo.Client {
	clientInstance, err := getMongoClient(GetEnv("MONGO_URI"))

	if err != nil {
		log.Fatal(err)
	}

	err = clientInstance.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return clientInstance
}
