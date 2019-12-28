package user

import (
	"ArticleManage/domain/value"
	"ArticleManage/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func Register(res http.ResponseWriter, req *http.Request) {
	var user value.User
	var result string
	json.NewDecoder(req.Body).Decode(&user)
	if err := service.Add(user); err != nil {
		result = "注册失败"
	}
	result = "注册成功"
	marshal, _ := json.Marshal(result)
	res.Write(marshal)
}
func ChangePassword(res http.ResponseWriter, req *http.Request) {
	type newUser struct {
		Name        string
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	var newu newUser
	var result string
	json.NewDecoder(req.Body).Decode(&newu)
	if id, err := service.GetUser(value.User{
		Name:     newu.Name,
		Password: newu.OldPassword,
	}); err != nil {
		result = "用户名或密码错误" + fmt.Sprintf("%v", err)
	} else {
		if err := service.Update(id, value.User{
			Name:     newu.Name,
			Password: newu.NewPassword,
		}); err != nil {
			result = "修改用户密码失败" + fmt.Sprintf("%v", err)
		} else {
			result = "修改成功"
		}
	}
	marshal, _ := json.Marshal(result)
	res.Write(marshal)
}
func Login(res http.ResponseWriter, req *http.Request) {
	var user value.User
	var result string
	json.NewDecoder(req.Body).Decode(&user)
	if _, err := service.GetUser(user); err != nil {
		result = "用户名或密码错误" + fmt.Sprintf("%v", err)
	}
	result = "登录成功"
	marshal, _ := json.Marshal(result)
	res.Write(marshal)
}
