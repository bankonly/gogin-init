package configs

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBClient   *mongo.Client
	DBInstance *mongo.Database
)

func DBContext() context.Context {
	return context.Background()
}

func ConnectDatabase() {
	client, err := mongo.Connect(DBContext(), options.Client().ApplyURI(Constant.MONGO_URI))
	if err != nil {
		fmt.Print(err)
		panic("Connect database failed")
	}

	DBClient = client
	DBInstance = client.Database(Constant.DATABASE_NAME)
}
