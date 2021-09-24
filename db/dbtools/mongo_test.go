package dbtools

import (
	"cainiaofundbackend/db"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestGetMany(t *testing.T) {
	userList := make([]db.User, 0)
	err := GetMany(context.TODO(), db.GetUserCol(), &userList, bson.M{})
	logrus.Printf("test get mang err:%v", err)
}
