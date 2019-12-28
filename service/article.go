package service

import (
	"ArticleManage/domain/value"
	"ArticleManage/storage/user_article_manage"
)

func AddArticle(article value.Article) error {
	art:=user_article_manage.NewArticleImpl()
	return art.Add(article)
}
func GetUserArticle(id value.UserId)([]value.Article,error)  {
	art:=user_article_manage.NewArticleImpl()
	return art.GetUserArticles(id)
}
