package setup

import (
	inMemoryService "go-getir/internal/in-memory/service"
	recordService "go-getir/internal/record/service"
	"go-getir/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupServices(cfg *config.Configuration, mc *mongo.Client) *Service {

	return &Service{
		RecordService:   recordService.NewService(mc.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection)),
		InMemoryService: inMemoryService.NewService(),
	}
}
