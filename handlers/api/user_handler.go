package api

import (
	"aries/models"
	"aries/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

// GetAllUsers
// @Summary 获取所有用户
// @Tags 用户
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/all_users [get]
func (u *UserHandler) GetAllUsers(ctx *gin.Context) {
	user := models.User{}.GetUniqueUser()
	userList := []models.User{user}

	ctx.JSON(http.StatusOK, utils.Result{
		Code: utils.Success,
		Msg:  "查询成功",
		Data: userList,
	})
}
