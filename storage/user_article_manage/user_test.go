package user_article_manage

import (
	"ArticleManage/domain/value"
	"ArticleManage/storage/config/constant"
	"ArticleManage/storage/model"
	"ArticleManage/storage/util/db"
	"fmt"
	"testing"
)

func Init() {
	db.InitDb(constant.Dbname, constant.DbDataSource)
}
func TestUserImpl_Add(t *testing.T) {
	Init()
	t.Run("添加用户", func(t *testing.T) {
		user := NewUserImpl()
		var u = value.User{
			Name:     "ming",
			Password: "ming",
		}
		err := user.Add(u)
		fmt.Println(err)
	})
}
func TestUserImpl_Update(t *testing.T) {
	Init()
	t.Run("修改密码", func(t *testing.T) {
		usr:=model.UserModel{
			Id:       0,
			Name:     "xiye",
			Password: "nnnn",
		}
		_,err:=db.DbConn.Where("id = ? and name = ?",6,"xiye").Update(&usr)
		fmt.Println(err)
	})
}
