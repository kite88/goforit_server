package controllers

import (
	"encoding/json"
	"goforit/services/AuthService"
	"goforit/services/LogService"
	"strconv"
	"strings"
)

type UserController struct {
	BaseController
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

// @router /login [post]
func (u *UserController) Login() {
	var data loginRequest
	param := u.Ctx.Input.RequestBody
	json.Unmarshal(param, &data)
	if strings.Trim(data.Username, " ") == "" || strings.Trim(data.Password, " ") == "" {
		u.NormalException("用户名或者密码不能为空")
	}
	user, count := AuthService.Login(data.Username, data.Password)
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
		LogService.CreateLogs(user.ID, 1, "登录成功")
		u.Normal(result, "登录成功")
	}
	u.Normal(NoneObject{}, "请求成功")
}

// @router /logout [get]
func (u *UserController) Logout() {
	userId := u.NeedLogin(true)
	AuthService.Logout(userId)
	LogService.CreateLogs(userId, 2, "退出登录")
	u.Normal(NoneObject{}, "退出成功")
}
