package setup

import (
	inMemoryService "go-getir/internal/in-memory/service"
	recordService "go-getir/internal/record/service"
)

type Service struct {
	RecordService   recordService.Service
	InMemoryService inMemoryService.Service
}
