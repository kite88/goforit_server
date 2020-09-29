package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["goforit/controllers:AccountController"] = append(beego.GlobalControllerRouter["goforit/controllers:AccountController"],
        beego.ControllerComments{
            Method: "ModifyPassword",
            Router: "/modifyPassword",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:MemoController"] = append(beego.GlobalControllerRouter["goforit/controllers:MemoController"],
        beego.ControllerComments{
            Method: "GetMemo",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:MemoController"] = append(beego.GlobalControllerRouter["goforit/controllers:MemoController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:MemoController"] = append(beego.GlobalControllerRouter["goforit/controllers:MemoController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:MemoController"] = append(beego.GlobalControllerRouter["goforit/controllers:MemoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:MemoController"] = append(beego.GlobalControllerRouter["goforit/controllers:MemoController"],
        beego.ControllerComments{
            Method: "GetClassify",
            Router: "/classify",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:MemoController"] = append(beego.GlobalControllerRouter["goforit/controllers:MemoController"],
        beego.ControllerComments{
            Method: "CreateClassify",
            Router: "/classify",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:MemoController"] = append(beego.GlobalControllerRouter["goforit/controllers:MemoController"],
        beego.ControllerComments{
            Method: "EditClassify",
            Router: "/classify",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:MemoController"] = append(beego.GlobalControllerRouter["goforit/controllers:MemoController"],
        beego.ControllerComments{
            Method: "DeleteClassify",
            Router: "/classify",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:UserController"] = append(beego.GlobalControllerRouter["goforit/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goforit/controllers:UserController"] = append(beego.GlobalControllerRouter["goforit/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
