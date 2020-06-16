package interfaces

import (
	"context"

	"blog/app/application"
	api "blog/app/proto/api"
)

type articleServer struct {
	app application.ArticleInteractor
}

func NewArticleServer(app application.ArticleInteractor) api.ArticleServer {
	return &articleServer{app}
}

func (s *articleServer) Get(ctx context.Context, req *api.GetArticleRequest) (*api.GetArticleResponse, error) {
	article, err := s.app.Get(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	return &api.GetArticleResponse{
		ArticleData: toAPIArticleData(article),
	}, nil
}

func (s *articleServer) List(ctx context.Context, req *api.ListArticleRequest) (*api.ListArticleResponse, error) {
	articles, err := s.app.List(ctx)
	if err != nil {
		return nil, err
	}
	return &api.ListArticleResponse{
		ArticleData: toAPIArticleDatas(articles),
	}, nil
}
