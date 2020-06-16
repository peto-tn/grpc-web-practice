package application

import (
	"context"

	"blog/app/domain/service"
)

type SystemInteractor interface {
	Ready(ctx context.Context) error
}

type systemInteractor struct {
	systemService service.SystemService
}

func NewSystemInteractor(
	systemService service.SystemService,
) SystemInteractor {
	return &systemInteractor{systemService}
}

func (i *systemInteractor) Ready(ctx context.Context) error {
	return i.systemService.Ready(ctx)
}
