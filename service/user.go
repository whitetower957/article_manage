package service

import (
	"ArticleManage/domain/value"
	"ArticleManage/storage/user_article_manage"
	"errors"
	"log"
)

func Add(user value.User) error {
	u :=user_article_manage.NewUserImpl()
	return u.Add(user)
}

func Del(id int) error {
	u :=user_article_manage.NewUserImpl()
	return u.Del(value.NewUserId(id))
}
func Update(id int, user value.User) error {
	u :=user_article_manage.NewUserImpl()
	return u.Update(value.NewUserId(id), user)
}
func GetUser(user value.User) (int, error) {
	u :=user_article_manage.NewUserImpl()
	id, exist, err := u.GetUser(user)
	if !exist {
		return -1, errors.New("用户名或密码错误")
	}
	if err != nil {
		log.Println(err)
	}
	return int(id), err
}
