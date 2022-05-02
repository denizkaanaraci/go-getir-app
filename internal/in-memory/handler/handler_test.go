package record

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	inMemoryService "go-getir/internal/in-memory/service"
)

func TestGet(t *testing.T) {
	t.Run("Get key from in-memory storage", func(t *testing.T) {
		service := inMemoryService.NewService()
		handler := NewHandler(service)

		service.Set("foo-get", "bar-get")
		target := "/in-memory?key=foo-get"

		r := httptest.NewRequest(http.MethodGet, target, nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("Expected %d, got %d", http.StatusOK, w.Code)
		}

		if w.Body.String() != `{"key":"foo-get","value":"bar-get"}` {
			t.Errorf("Expected %s, got %s", `{"key":"foo-get","value":"bar-get"}`, w.Body.String())
		}

	})
}

func TestSet(t *testing.T) {
	t.Run("Set key", func(t *testing.T) {
		service := inMemoryService.NewService()
		handler := NewHandler(service)

		payload := []byte(`{"key":"foo-set","value":"bar-set"}`)

		r := httptest.NewRequest(http.MethodPost, "/in-memory", bytes.NewBuffer(payload))
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, r)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected %d, got %d", http.StatusCreated, w.Code)
		}
		value, _ := service.Get("foo-set")
		if value != "bar-set" {
			t.Errorf("Expected %s, got %s", "bar-set", value)
		}

	})
}
