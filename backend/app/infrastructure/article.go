package infrastructure

import (
	"context"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql" // driver

	"blog/app/domain/entity"
	"blog/app/domain/repository"
	cerror "blog/app/error"
	"blog/app/infrastructure/model"
	"blog/app/infrastructure/persistence/db"
)

type articleRepository struct {
	conn *db.Connection
}

func NewArticleRepository(conn *db.Connection) repository.ArticleRepository {
	return &articleRepository{conn}
}

func (r *articleRepository) GetById(ctx context.Context, articleId int64) (*entity.Article, error) {
	db := r.conn.GetConnection(ctx)
	mArticle := model.Article{}
	res := db.Where("article_id = ?", articleId).Take(&mArticle)
	if res.RecordNotFound() {
		return nil, cerror.NewNotFound(fmt.Sprintf("article_id=%s", articleId))
	}
	return mArticle.ToEntity(), res.Error
}

func (r *articleRepository) List(ctx context.Context) ([]*entity.Article, error) {
	db := r.conn.GetConnection(ctx)
	mArticles := []*model.Article{}
	err := db.Find(&mArticles).Error
	return model.ToEntitiesArticle(mArticles), err
}
