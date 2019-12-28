package user_article_manage

import (
	"ArticleManage/domain/value"
	"ArticleManage/storage/model"
	"ArticleManage/storage/util/db"
)

type ArticleImpl struct {
}

func NewArticleImpl() ArticleImpl {
	return ArticleImpl{}
}
func (a *ArticleImpl) Add(article value.Article) error {
	art := model.NewArticleModel(0, article)
	_, err := db.DbConn.Insert(&art)
	return err
}

func (a *ArticleImpl) Del(id value.ArticleId) error {
	art := model.NewArticleModel(int(id), value.Article{})
	_, err := db.DbConn.ID(int(id)).Delete(&art)
	return err
}

func (a *ArticleImpl) Update(id value.ArticleId, article value.Article) error {
	art := model.NewArticleModel(int(id), article)
	_, err := db.DbConn.ID(int(id)).Update(&art)
	return err
}

func (a *ArticleImpl) GetArticle(id value.ArticleId) (article value.Article, exist bool, err error) {
	art := model.NewArticleModel(int(id), value.Article{})
	exist, err = db.DbConn.Get(&art)
	article.CreateUserId = art.CreateUserId
	article.Title = art.Title
	article.Content = art.Content
	article.CreateTime = art.CreateTime
	return article, exist, err
}

func (a *ArticleImpl) GetUserArticles(id value.UserId) (articles []value.Article, err error) {
	err = db.DbConn.Where("create_user_id = ?", int(id)).Find(&articles)
	return articles, err
}
