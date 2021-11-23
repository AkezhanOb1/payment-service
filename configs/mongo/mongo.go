package mongo

import (
	"context"
	"os"

	"github.com/AkezhanOb1/payment/configs/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Payments is a collection where we store donations
var Payments *mongo.Collection

func InitMongo() {
	var err error
	var connectionStr = os.Getenv("MongoClientURI")
	var clientOptions = options.Client().ApplyURI(connectionStr)
	MongoConn, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Logger.Fatal("can not connect to mongoDB", err)
	}

	err = MongoConn.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Logger.Fatal("can not ping the mongoDB", err)
	}

	Payments = MongoConn.Database("bloomzed").Collection("payments")

	log.Logger.Info("successfully connected to the mongoDB")
}
