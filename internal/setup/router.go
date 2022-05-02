package setup

import (
	"net/http"

	inMemoryHandler "go-getir/internal/in-memory/handler"
	recordHandler "go-getir/internal/record/handler"
)

func SetupRoutes(service *Service) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/record", recordHandler.NewHandler(service.RecordService))
	mux.Handle("/in-memory", inMemoryHandler.NewHandler(service.InMemoryService))

	return mux
}
