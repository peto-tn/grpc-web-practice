package service

import (
	"context"
)

type SystemService interface {
	Ready(ctx context.Context) error
}
