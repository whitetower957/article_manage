package article

import (
	"ArticleManage/domain/value"
	"ArticleManage/service"
	"encoding/json"
	"net/http"
)

func AddArticle(res http.ResponseWriter, req *http.Request) {
	var article value.Article
	var result string
	json.NewDecoder(req.Body).Decode(&article)
	if err := service.AddArticle(article); err != nil {
		result = "添加文章失败"
	}
	result = "添加文章成功"
	marshal, _ := json.Marshal(result)
	res.Write(marshal)
}

func GetUserArticle(res http.ResponseWriter, req *http.Request) {
	type Userr struct {
		Id value.UserId
	}
	var u Userr
	var result string
	var arts []value.Article
	var err error
	var marshal []byte
	json.NewDecoder(req.Body).Decode(&u)
	if arts, err = service.GetUserArticle(u.Id); err != nil {
		result = "获取用户文章失败"
		marshal, _ = json.Marshal(result)
	}
	marshal, _ = json.Marshal(arts)
	res.Write(marshal)
}
