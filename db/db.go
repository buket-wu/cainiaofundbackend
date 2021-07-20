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
	UserCol         *mongo.Collection
	FundCol         *mongo.Collection
	FundTrendCol    *mongo.Collection
	RemindRecordCol *mongo.Collection
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
	db := client.Database(config.Config.Mongo.DB)
	UserCol = db.Collection("user")
	FundCol = db.Collection("fund")
	FundTrendCol = db.Collection("fundTrend")
	RemindRecordCol = db.Collection("remindRecord")
}
