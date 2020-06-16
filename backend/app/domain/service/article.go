package service

import (
	"context"

	"blog/app/domain/entity"
)

type ArticleService interface {
	Get(ctx context.Context, id int64) (*entity.Article, error)
	List(ctx context.Context) ([]*entity.Article, error)
}
