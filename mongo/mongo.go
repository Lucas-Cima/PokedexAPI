package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client *mongo.Client
)

//Connects to the DB
func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://Lucas:pokemon@pokedex.l4iml.mongodb.net/Pokedex?retryWrites=true&w=majority",
	))
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Printf("could not ping to mongo db service: %v\n", err)
		return
	}
}
