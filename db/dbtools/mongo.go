package dbtools

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

func GetMany(ctx context.Context, col *mongo.Collection, v interface{}, filter interface{}, opts ...*options.FindOptions) error {
	r := reflect.TypeOf(v)
	if r.Kind() != reflect.Ptr {
		return fmt.Errorf("results argument must be a pointer, but was a %v", r.Kind())
	}

	if r.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("results argument must be a pointer to a slice, but was a %v", r.Elem().Kind())
	}

	cur, err := col.Find(ctx, filter, opts...)
	if err != nil {
		logrus.Errorf("get many fail; err:%v", err)
		return err
	}
	defer func() {
		err = cur.Close(ctx)
		if err != nil {
			logrus.Errorf("mongo cur close fail; err:%v", err)
		}
	}()

	err = cur.All(ctx, v)
	if err != nil {
		logrus.Errorf("cur all decode fail; err:%v", err)
		return err
	}

	return nil
}
