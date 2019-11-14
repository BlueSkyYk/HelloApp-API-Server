package controllers

import (
	"HelloApp_Api/constants"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BaseController struct {
	o orm.Ormer
	beego.Controller
}

func (c *BaseController) Prepare() {
	//创建Ormer
	c.o = orm.NewOrm()
}

/**
响应成功
 */
func (c *BaseController) responseSuccess(code int, msg string) {
	c.responseNotData(code, true, msg)
}

/**
响应失败
 */
func (c *BaseController) responseFail(code int, msg string) {
	c.responseNotData(code, false, msg);
}

func (c *BaseController) responseNotData(code int, success bool, msg string) {
	c.Data["json"] = constants.Result{
		Code:    code,
		Success: success,
		Message: msg,
	}
}
