package models

import "aries/utils"

// User 用户
type User struct {
	ID       uint
	Username string
	Email    string
	Pwd      string
}

// GetUniqueUser 根据用户名和密码获取用户
func (user User) GetUniqueUser() User {
	pwd, _ := utils.EncryptPwd("password")
	return User{
		ID:       uint(1001),
		Username: "admin",
		Email:    "helloworld@a.com",
		Pwd:      pwd,
	}
}
