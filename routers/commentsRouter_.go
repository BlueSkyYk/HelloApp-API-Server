package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["HelloApp_Api/controllers:SmsVerCodeController"] = append(beego.GlobalControllerRouter["HelloApp_Api/controllers:SmsVerCodeController"],
        beego.ControllerComments{
            Method: "GetSmsVerCode",
            Router: `/getSmsVerCode`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["HelloApp_Api/controllers:UserController"] = append(beego.GlobalControllerRouter["HelloApp_Api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["HelloApp_Api/controllers:UserController"] = append(beego.GlobalControllerRouter["HelloApp_Api/controllers:UserController"],
        beego.ControllerComments{
            Method: "SearchUser",
            Router: `/searchUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["HelloApp_Api/controllers:UserFriendController"] = append(beego.GlobalControllerRouter["HelloApp_Api/controllers:UserFriendController"],
        beego.ControllerComments{
            Method: "AddFriendRequest",
            Router: `/addFriendRequest`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["HelloApp_Api/controllers:UserFriendController"] = append(beego.GlobalControllerRouter["HelloApp_Api/controllers:UserFriendController"],
        beego.ControllerComments{
            Method: "AddFriendResult",
            Router: `/addFriendResult`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["HelloApp_Api/controllers:UserFriendController"] = append(beego.GlobalControllerRouter["HelloApp_Api/controllers:UserFriendController"],
        beego.ControllerComments{
            Method: "GetUserFriendList",
            Router: `/getUserFriendList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
