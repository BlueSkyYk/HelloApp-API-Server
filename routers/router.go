// @APIVersion 1.0.0
// @Title HelloApp API接口
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact 1362616851@qq.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"HelloApp_Api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/helloapp/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{}),
		),
		beego.NSNamespace("/sms",
			beego.NSInclude(
				&controllers.SmsVerCodeController{}),
		),
		beego.NSNamespace("/userFriend",
			beego.NSInclude(
				&controllers.UserFriendController{}),
		),
	)
	beego.AddNamespace(ns)
}
