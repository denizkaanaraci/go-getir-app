package record

import (
	"encoding/json"
	helper "go-getir/internal/pkg/http-helper"
	"go-getir/internal/record/model"
	"go-getir/internal/record/service"
	"log"
	"net/http"
)

type handler struct {
	recordService service.Service
}

func NewHandler(recordService service.Service) *handler {
	return &handler{
		recordService: recordService,
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		var modelRequest model.Request

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&modelRequest)

		if err != nil {
			log.Println(err)
			helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		result, _ := h.recordService.Fetch(&modelRequest)
		helper.RespondWithJSON(w, http.StatusOK, result)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}
