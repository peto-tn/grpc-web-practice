package domain

import (
	"context"

	"blog/app/domain/repository"
	"blog/app/domain/service"
)

type systemService struct {
	systemRepo repository.SystemRepository
}

func NewSystemService(systemRepo repository.SystemRepository) service.SystemService {
	return &systemService{systemRepo}
}

func (s *systemService) Ready(ctx context.Context) error {
	return s.systemRepo.Ready(ctx)
}
