package db

import (
	"cainiaofundbackend/config"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Database
)

func init() {
	mongoConf := fmt.Sprintf("mongodb://%s:%s@%s/%s?connect=direct",
		config.Config.Mongo.User,
		config.Config.Mongo.Password,
		config.Config.Mongo.Addr,
		config.Config.Mongo.AuthDB)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConf))
	if err != nil {
		logrus.Fatalf("err:%v", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logrus.Fatalf("err:%v", err)
	}
	MongoClient = client.Database(config.Config.Mongo.DB)
}

func GetUserCol() *mongo.Collection {
	return MongoClient.Collection("user")
}

func GetFundCol() *mongo.Collection {
	return MongoClient.Collection("fund")
}

func GetFundTrendCol() *mongo.Collection {
	return MongoClient.Collection("fundTrend")
}

func GetRemindRecordCol() *mongo.Collection {
	return MongoClient.Collection("remindRecord")
}

func GetFundRelationCol() *mongo.Collection {
	return MongoClient.Collection("fundRelation")
}
