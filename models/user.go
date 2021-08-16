package models

// User 用户
type User struct {
	ID       uint
	Username string
	Email    string
	Pwd      string
}

// GetUniqueUser 根据用户名和密码获取用户
func (user User) GetUniqueUser() User {
	return User{
		ID:       uint(1001),
		Username: "username",
		Email:    "helloworld@a.com",
		Pwd:      "pwd",
	}
}
