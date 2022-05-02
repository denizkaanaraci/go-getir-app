package service

import (
	"errors"
	"sync"
)

type Service interface {
	Get(k string) (string, error)
	Set(k string, v string)
}

type service struct {
	storage map[string]string
	mu      *sync.Mutex
}

func NewService() *service {
	return &service{
		storage: make(map[string]string),
		mu:      &sync.Mutex{},
	}
}

func (s *service) Get(k string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if val, exists := s.storage[k]; exists {
		return val, nil
	}
	return "", errors.New("Key not found")
}

func (s *service) Set(k string, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.storage[k] = v
}
