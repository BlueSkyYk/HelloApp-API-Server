package controllers

import (
	"HelloApp_Api/constants"
	"HelloApp_Api/reqparams"
	"HelloApp_Api/utils"
	"encoding/json"
	"fmt"
)

type SmsVerCodeController struct {
	BaseController
}

// @Title 获取短信验证码
// @Description 过去各种类型的短信验证码
// @Param	body	body	reqparams.SmsVerCodeParam	true	"The SmsVerCodeParam content"
// @Success 200 {object} resparams.SmsVerCodeResult
// @Failure -100 :获取短信验证码失败
// @router /getSmsVerCode [post]
func (c *SmsVerCodeController) GetSmsVerCode() {
	param := reqparams.SmsVerCodeParam{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {
		var phone = param.Phone
		//检查手机号码
		if phone == "" || !utils.CheckPhone(phone) {
			c.Data["json"] = constants.Result{
				Code:    constants.PHONE_ERROR,
				Message: "手机号码错误",
				Success: false,
			}
			c.ServeJSON()
			return
		}
		//获取短信验证码
		if err := utils.GetSmsCode(phone); err == nil {
			c.Data["json"] = constants.Result{
				Code:    constants.RESULT_OK,
				Message: "获取短信验证码成功",
				Success: true,
			}
		} else {
			c.Data["json"] = constants.Result{
				Code:    constants.RESULT_ERROR,
				Message: err.Error(),
				Success: false,
			}
		}
		c.ServeJSON()
	} else { //请求参数错误
		fmt.Println(err)
		c.Data["json"] = constants.Result{
			Code:    0,
			Message: "参数错误",
			Success: false,
		}
		c.ServeJSON()
	}
}
