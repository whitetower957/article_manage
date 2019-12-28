package user_article_manage

import (
	"ArticleManage/domain/value"
	"ArticleManage/storage/model"
	"ArticleManage/storage/util/db"
)

type UserImpl struct {
}

func NewUserImpl() UserImpl {
	return UserImpl{}
}
func (u *UserImpl) Add(user value.User) error {
	usr := model.NewUserModel(0, user)
	_, err := db.DbConn.Insert(&usr)
	return err
}

func (u *UserImpl) Del(id value.UserId) error {
	usr := model.NewUserModel(int(id), value.User{})
	_, err := db.DbConn.ID(int(id)).Delete(&usr)
	return err
}

func (u *UserImpl) Update(id value.UserId, user value.User) error {
	usr := model.NewUserModel(0, user)
	_, err := db.DbConn.Where("id = ?",id).Update(&usr)
	return err
}

func (u *UserImpl) GetUser(user value.User) (id value.UserId, exist bool, err error) {
	usr := model.NewUserModel(int(id), user)
	exist, err = db.DbConn.Where("name = ? and password = ?",
		user.Name,user.Password).Get(&usr)
	return value.NewUserId(usr.Id), exist, err
}
