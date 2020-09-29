package controllers

type AccountController struct {
	BaseController
}

// @router /modifyPassword [put]
func (a *AccountController) ModifyPassword() {

	a.Normal(NoneObject{},"123")
}
