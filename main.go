package main

import (
	"ArticleManage/perference/article"
	"ArticleManage/perference/user"
	"ArticleManage/storage/config/constant"
	"ArticleManage/storage/util/db"
	"net/http"
)

func main() {
	db.InitDb(constant.Dbname, constant.DbDataSource)
	http.HandleFunc("/user/register", user.Register)
	http.HandleFunc("/user/changePassword", user.ChangePassword)
	http.HandleFunc("/user/login", user.Login)
	http.HandleFunc("/article/add", article.AddArticle)
	http.HandleFunc("/article/getUserArticle", article.GetUserArticle)
	http.ListenAndServe(":8888", nil)
}
