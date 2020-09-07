package controllers

import (
	"encoding/json"
	"goforit/services/AuthService"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

type loginRequest struct {
	Username string
	Password string
}

type loginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// @router /login [post]
func (u *UserController) Login() {
	var param loginRequest
	data := u.Ctx.Input.RequestBody
	json.Unmarshal(data, &param)
	if strings.Trim(param.Username, " ") == "" || strings.Trim(param.Password, " ") == "" {
		u.NormalException("用户名或者密码不能为空")
	}
	user, count := AuthService.Login(param.Username, param.Password)
	if count == 0 {
		u.NormalException("用户不存在或者密码有误")
	}
	if count == 1 {
		token := AuthService.Token(user.ID, u.CreateToken(strconv.Itoa(int(user.ID))))
		result := loginResponse{
			ID:       user.ID,
			Username: user.Username,
			Token:    token,
		}
		u.Normal(result, "登录成功")
	}
	u.Normal(nil, "请求成功")
}

// @router /logout [get]
func (u *UserController) Logout() {
	userId := u.NeedLogin(true)
	AuthService.Logout(userId)
	u.Normal(nil,"退出成功")
}