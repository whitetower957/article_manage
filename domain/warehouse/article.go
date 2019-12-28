package warehouse

import "ArticleManage/domain/value"

type ArticleOperation interface {
	Add(article value.Article) error
	Del(id value.ArticleId) error
	Update(article value.Article) error
	//根据文章ID获取文章所有信息
	GetArticle(id value.ArticleId) (article value.Article, exist bool, err error)
	//获取用户所有文章
	GetUserArticles(id value.UserId) (articles []value.Article, err error)
}
