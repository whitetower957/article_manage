package warehouse

import "ArticleManage/domain/value"

type UserOperation interface {
	Add(user value.User) error
	Del(id value.UserId) error
	Update(user value.User) error
	GetUser(user value.User) (id value.UserId, exist bool, err error)
}
