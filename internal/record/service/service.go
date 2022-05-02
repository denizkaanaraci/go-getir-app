package service

import (
	"context"
	"log"
	"net/http"
	"time"

	"go-getir/internal/record/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	Fetch(r *model.Request) (*model.Response, error)
}

type service struct {
	mc *mongo.Collection
}

func NewService(mc *mongo.Collection) *service {
	return &service{mc: mc}
}

func (s *service) Fetch(modelRequest *model.Request) (*model.Response, error) {

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	resp := new(model.Response)
	resp.Records = make([]model.Record, 0)

	start, err := time.Parse("2006-01-02", modelRequest.StartDate)
	if err != nil {
		resp.Msg = err.Error()
		return resp, err
	}
	end, err := time.Parse("2006-01-02", modelRequest.EndDate)
	if err != nil {
		resp.Msg = err.Error()
		return resp, err
	}
	filter := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gt": start,
					"$lt": end,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gt": modelRequest.MinCount,
					"$lt": modelRequest.MaxCount,
				},
			},
		},
	}

	cursor, err := s.mc.Aggregate(ctx, filter)

	if err != nil {
		log.Printf("MongoDB aggregate failed. Error: %s", err.Error())
		resp.Msg = err.Error()
		resp.Code = http.StatusInternalServerError
		return resp, err
	}

	if err = cursor.All(ctx, &resp.Records); err != nil {
		resp.Msg = err.Error()
		resp.Code = http.StatusInternalServerError
		log.Printf("MongoDB cursor failed. Error %s", err.Error())
		return resp, err
	}

	resp.Msg = "Success"
	return resp, nil
}
