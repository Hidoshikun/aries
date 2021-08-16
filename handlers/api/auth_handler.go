package api

import (
	"aries/config/setting"
	"aries/forms"
	"aries/log"
	"aries/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

// Login
// @Summary 登录
// @Tags 授权
// @version 1.0
// @Accept application/json
// @Param loginForm body forms.LoginForm true "登录表单"
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/auth/login [post]
func (a *AuthHandler) Login(ctx *gin.Context) {
	loginForm := forms.LoginForm{}
	result := utils.Result{ // 定义 api 返回信息结构
		Code: utils.Success,
		Msg:  "登录成功",
		Data: nil,
	}

	if err := ctx.ShouldBindJSON(&loginForm); err != nil { // 表单校验失败
		result.Code = utils.RequestError     // 请求数据有误
		result.Msg = utils.GetFormError(err) // 获取表单错误信息
		ctx.JSON(http.StatusOK, result)      // 返回 json
		return
	}

	captchaConfig := &utils.CaptchaConfig{
		Id:          loginForm.CaptchaId,
		VerifyValue: loginForm.CaptchaVal,
	}
	if !utils.CaptchaVerify(captchaConfig) { // 校验验证码
		result.Code = utils.RequestError // 请求数据有误
		result.Msg = "验证码错误"
		ctx.JSON(http.StatusOK, result) // 返回 json
		return
	}

	user := loginForm.BindToModel() // 绑定表单数据到实体类
	u := user.GetUniqueUser()    // 根据用户名获取用户
	if u.Username == "" {           // 用户不存在
		result.Code = utils.RequestError
		result.Msg = "不存在该用户"
		ctx.JSON(http.StatusOK, result)
		return
	}
	if !utils.VerifyPwd(u.Pwd, user.Pwd) { // 密码错误
		result.Code = utils.RequestError
		result.Msg = "密码错误"
		ctx.JSON(http.StatusOK, result) // 返回 json
		return
	}

	j := utils.NewJWT() // 创建 JWT 实例
	token, err := j.CreateToken(utils.CustomClaims{ // 生成 JWT token
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.
				Duration(setting.Config.Server.TokenExpireTime)).Unix(), // 设置过期时间
			IssuedAt: time.Now().Unix(),
		},
	})
	if err != nil { // 异常处理
		log.Logger.Sugar().Error("error: ", err.Error())
		result.Code = utils.ServerError
		result.Msg = "服务器端错误"
		ctx.JSON(http.StatusOK, result) // 返回 json
		return
	}

	result.Data = utils.Token{ // 封装 Token 信息
		Token:    token,
		Username: u.Username,
	}

	ctx.JSON(http.StatusOK, result) // 返回 json
}

// CreateCaptcha
// @Summary 创建验证码
// @Tags 授权
// @version 1.0
// @Accept application/json
// @Success 100 object utils.Result 成功
// @Failure 103/104 object utils.Result 失败
// @Router /api/v1/auth/captcha [get]
func (a *AuthHandler) CreateCaptcha(ctx *gin.Context) {
	captcha := utils.CaptchaConfig{} // 创建验证码配置结构
	result := utils.Result{ // 返回数据结构
		Code: utils.Success,
		Msg:  "验证码创建成功",
		Data: nil,
	}

	base64, err := utils.GenerateCaptcha(&captcha) // 创建验证码
	if err != nil {                                // 异常处理
		result.Code = utils.ServerError
		result.Msg = "服务器端错误"
		ctx.JSON(http.StatusOK, result)
		return
	}

	result.Data = gin.H{ // 封装 data
		"captcha_id":  captcha.Id,
		"captcha_url": base64,
	}

	ctx.JSON(http.StatusOK, result) // 返回 json 数据
}
