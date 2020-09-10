package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"goforit/services/AuthService"
	"goforit/utils"
	"time"
)

type BaseController struct {
	beego.Controller
	Uid uint64
}

type NoneObject struct{}
type NoneArray []struct{}

type baseRequest struct {
	PageIndex string `json:"pageIndex"`
	PageSize  string `json:"pageSize"`
}

type data struct {
	Code   int64       `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

var (
	successCode int64 = 0
	failCode    int64 = -1
	authCode    int64 = -10
)

func (t *BaseController) Normal(result interface{}, msg string) {
	t.Data["json"] = data{
		Code:   successCode,
		Msg:    msg,
		Result: result,
	}
	t.ServeJSON()
	t.StopRun()
}

func (t *BaseController) NormalException(msg string) {
	t.Data["json"] = data{
		Code:   failCode,
		Msg:    msg,
		Result: NoneObject{},
	}
	t.ServeJSON()
	t.StopRun()
}

func (t *BaseController) CreateToken(part string) string {
	str := fmt.Sprintf("%s-%d", part, time.Now().UnixNano())
	return utils.Md5(str)
}

func (t *BaseController) NeedLogin(b bool) uint64 {
	var tokenStr string
	var header = t.Ctx.Request.Header
	token := header["Token"]
	if token != nil {
		tokenStr = token[0]
	}
	uid := AuthService.GetUserIdByToken(tokenStr)
	t.Uid = uid
	if b == true && t.Uid == 0 {
		t.Data["json"] = data{
			Code:   authCode,
			Msg:    "请登录",
			Result: NoneObject{},
		}
		t.ServeJSON()
		t.StopRun()
	}
	return t.Uid
}
