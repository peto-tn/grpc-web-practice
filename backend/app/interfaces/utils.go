package interfaces

import (
	"blog/app/domain/entity"
	data "blog/app/proto/data"
)

func toAPIArticleData(article *entity.Article) *data.ArticleData {
	return &data.ArticleData{
		ArticleId: article.Id,
		Title:     article.Title,
		Body:      article.Body,
		CreatedAt: article.CreatedAt.Unix(),
		UpdatedAt: article.UpdatedAt.Unix(),
	}
}

func toAPIArticleDatas(articles []*entity.Article) []*data.ArticleData {
	datas := []*data.ArticleData{}
	for _, article := range articles {
		datas = append(datas, toAPIArticleData(article))
	}
	return datas
}
