package repository

import (
	"context"

	"blog/app/domain/entity"
)

type ArticleRepository interface {
	GetById(ctx context.Context, articleId int64) (*entity.Article, error)
	List(ctx context.Context) ([]*entity.Article, error)
}
