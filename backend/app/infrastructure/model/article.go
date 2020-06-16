package model

import (
	"blog/app/domain/entity"
	"time"

	"github.com/kayac/ddl-maker/dialect"
	"github.com/kayac/ddl-maker/dialect/mysql"
)

type Article struct {
	Id        int64     `db:"id" ddl:"auto"`
	Title     string    `db:"title" ddl:"size=100"`
	Body      string    `db:"body" ddl:"type=longtext"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m Article) Table() string {
	return "articles"
}

func (m Article) PrimaryKey() dialect.PrimaryKey {
	return mysql.AddPrimaryKey("id")
}

func (m Article) Indexes() dialect.Indexes {
	return dialect.Indexes{
		mysql.AddIndex("created_at_idx", "created_at"),
		mysql.AddIndex("updated_at_idx", "updated_at"),
	}
}

func (m *Article) Parse(article *entity.Article) *Article {
	if article == nil {
		return nil
	}
	*m = Article{
		Id:    article.Id,
		Title: article.Title,
	}
	return m
}

func (m *Article) ToEntity() *entity.Article {
	return &entity.Article{
		Id:    m.Id,
		Title: m.Title,
	}
}

func ToEntitiesArticle(mArticles []*Article) []*entity.Article {
	articles := []*entity.Article{}
	for _, mArticle := range mArticles {
		if article := mArticle.ToEntity(); article != nil {
			articles = append(articles, article)
		}
	}
	return articles
}
