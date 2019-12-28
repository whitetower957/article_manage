package model

import "ArticleManage/domain/value"

type UserModel struct {
	Id       int    `xorm:"id"`
	Name     string `xorm:"name"`
	Password string `xorm:"password"`
}

func (u UserModel) TableName() string {
	return "user"
}
func NewUserModel(id int, user value.User) UserModel {
	return UserModel{
		Id:       id,
		Name:     user.Name,
		Password: user.Password,
	}
}
