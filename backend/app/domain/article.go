package domain

import (
	"context"

	"blog/app/domain/entity"
	"blog/app/domain/repository"
	"blog/app/domain/service"
)

type articleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) service.ArticleService {
	return &articleService{articleRepo}
}

func (s *articleService) Get(ctx context.Context, id int64) (*entity.Article, error) {
	return s.articleRepo.GetById(ctx, id)
}

func (s *articleService) List(ctx context.Context) ([]*entity.Article, error) {
	return s.articleRepo.List(ctx)
}
