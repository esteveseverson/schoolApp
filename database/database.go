package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

type Counter struct {
	ID  string `bson:"_id"`
	Seq int    `bson:"seq"`
}

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("schoolApp")
	log.Println("Connected to MongoDB!")
}

func GetNextSequence(collection *mongo.Collection, sequenceName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var counter Counter
	filter := bson.M{"_id": sequenceName}
	update := bson.M{"$inc": bson.M{"seq": 1}}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)

	err := collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&counter)
	if err != nil {
		return 0, err
	}

	return counter.Seq, nil
}
