package value

type UserId int

func NewUserId(userId int) UserId {
	return UserId(userId)
}

type User struct {
	Name     string
	Password string
}

func NewUser(name, password string) User {
	return User{
		Name:     name,
		Password: password,
	}
}
