package mongo

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestNewMongoClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//client := NewMongoClient("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true")
	client := NewMongoClient("mongodb://localhost:27017/getir-case-studyy")

	defer client.Disconnect(ctx)

	if client == nil {
		t.Error("Mongo connection failed")
	}

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		t.Error("Cannot ping Mongo!", err)
	}
}
