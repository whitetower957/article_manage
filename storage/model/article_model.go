package model

import (
	"ArticleManage/domain/value"
	"time"
)

type ArticleModel struct {
	Id           int       `xorm:"id"`
	CreateUserId int       `xorm:"create_user_id"`
	Title        string    `xorm:"title"`
	Content      string    `xorm:"content"`
	CreateTime   time.Time `xorm:"create_time"`
}

func (a ArticleModel) TableName() string {
	return "article"
}
func NewArticleModel(id int, article value.Article) ArticleModel {
	return ArticleModel{
		Id:           id,
		CreateUserId: article.CreateUserId,
		Title:        article.Title,
		Content:      article.Content,
		CreateTime:   article.CreateTime,
	}
}
