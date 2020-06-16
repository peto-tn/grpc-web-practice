package repository

import (
	"context"
)

type SystemRepository interface {
	Ready(ctx context.Context) error
}
