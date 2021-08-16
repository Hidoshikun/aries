package forms

import (
	"aries/models"
)

// UserInfoForm 用户信息表单
type UserInfoForm struct {
	Username  string `json:"username" binding:"required,min=3,max=30" label:"用户名"`
	Email     string `json:"email" binding:"required,max=30,email" label:"邮箱"`
}

// BindToModel 绑定用户信息表单到用户实体
func (form UserInfoForm) BindToModel() models.User {
	return models.User{
		Username:  form.Username,
		Email:     form.Email,
	}
}
