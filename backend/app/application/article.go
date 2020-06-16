package application

import (
	"context"

	"blog/app/domain/entity"
	"blog/app/domain/service"
	"blog/app/library"
)

type ArticleInteractor interface {
	Get(ctx context.Context, id int64) (*entity.Article, error)
	List(ctx context.Context) ([]*entity.Article, error)
}

type articleInteractor struct {
	transaction    library.Transaction
	articleService service.ArticleService
}

func NewArticleInteractor(
	transaction library.Transaction,
	articleService service.ArticleService,
) ArticleInteractor {
	return &articleInteractor{transaction, articleService}
}

func (i *articleInteractor) Get(ctx context.Context, id int64) (*entity.Article, error) {
	return i.articleService.Get(ctx, id)
}

func (i *articleInteractor) List(ctx context.Context) ([]*entity.Article, error) {
	return i.articleService.List(ctx)
}
