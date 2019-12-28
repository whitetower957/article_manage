package value

import (
	"time"
)

type ArticleId int

func NewArticleId(id int) ArticleId {
	return ArticleId(id)
}

type Article struct {
	CreateUserId int       `json:"create_user_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreateTime   time.Time `json:"create_time"`
}

func NewArticle(createUserId int, title, content string, createTime time.Time) Article {
	return Article{
		CreateUserId: createUserId,
		Title:        title,
		Content:      content,
		CreateTime:   createTime,
	}
}
