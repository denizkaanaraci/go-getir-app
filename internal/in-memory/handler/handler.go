package record

import (
	"encoding/json"
	"go-getir/internal/in-memory/model"
	"go-getir/internal/in-memory/service"
	helper "go-getir/internal/pkg/http-helper"
	"log"
	"net/http"
)

type handler struct {
	inMemoryService service.Service
}

func NewHandler(inMemoryService service.Service) *handler {
	return &handler{
		inMemoryService: inMemoryService,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		key := r.URL.Query().Get("key")
		value, err := h.inMemoryService.Get(key)

		if err != nil {
			log.Println(err)
			helper.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		helper.RespondWithJSON(w, http.StatusCreated, model.GetResponse{Key: key, Value: value})
	case http.MethodPost:
		var modelRequest model.SetRequest

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&modelRequest)

		if err != nil {
			log.Println(err)
			helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		h.inMemoryService.Set(modelRequest.Key, modelRequest.Value)

		helper.RespondWithJSON(w, http.StatusOK, model.SetResponse{Key: modelRequest.Key, Value: modelRequest.Value})

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}
