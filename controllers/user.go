package controllers

import (
	"HelloApp_Api/constants"
	"HelloApp_Api/models"
	"HelloApp_Api/reqparams"
	"HelloApp_Api/resparams"
	"HelloApp_Api/resutil"
	"encoding/json"
	"fmt"
)

//用户相关接口
type UserController struct {
	BaseController
}

// @Title 用户登录
// @Description 用户各种方式的登陆
// @Param	body	body	reqparams.LoginParam	true	"The LoginParam content"
// @Success 200 {object} resparams.LoginResult
// @Failure -100 :登陆失败
// @router /login [post]
func (c *UserController) Login() {
	param := reqparams.LoginParam{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {
		fmt.Println("param", param)

		//检查登录类型
		switch param.LoginType {
		case 1: //手机号密码登陆
			code, msg, user, success := models.LoginByPassword(c.o, param.Phone, param.Password)
			if success {
				c.Data["json"] = resparams.LoginResult{
					Result: constants.Result{
						Code:    code,
						Message: msg,
						Success: success,
					},
					Data: resutil.CreateResultUser(*user),
				}
			} else {
				c.Data["json"] = constants.Result{
					Code:    code,
					Message: msg,
					Success: success,
				}
			}
			c.ServeJSON()
			break
		case 2: //短信登录
			code, msg, user, success := models.LoginBySmsCode(c.o, param.Phone, param.SmsCode)
			if success {
				c.Data["json"] = resparams.LoginResult{
					Result: constants.Result{
						Code:    code,
						Message: msg,
						Success: success,
					},
					Data: resutil.CreateResultUser(*user),
				}
			} else {
				c.Data["json"] = constants.Result{
					Code:    code,
					Message: msg,
					Success: success,
				}
			}
			c.ServeJSON()
			break
		case 3: //qq登录

			break
		}
	} else { //请求参数错误
		fmt.Println(err)
		c.Data["json"] = constants.Result{
			Code:    constants.RESULT_ERROR,
			Message: "参数错误",
			Success: false,
		}
		c.ServeJSON()
	}
}

// @Title 查找用户
// @Description 通过关键字查找用户
// @Param	body	body	reqparams.SearchUser	true	"The SearchUser content"
// @Success 200 {object} resparams.LoginResult
// @Failure -100 :登陆失败
// @router /searchUser [post]
func (c *UserController) SearchUser() {
	param := reqparams.SearchUser{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {
		code, msg, searchUsers, success := models.SearchUser(c.o, param.Keyword)
		if success {
			c.Data["json"] = resparams.SearchUserResult{
				Result: constants.Result{
					Code:    code,
					Message: msg,
					Success: success,
				},
				Data: resutil.CreateResultSearchUser(*searchUsers),
			}
		} else {
			c.Data["json"] = constants.Result{
				Code:    code,
				Message: msg,
				Success: success,
			}
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = constants.Result{
			Code:    constants.RESULT_ERROR,
			Message: "参数错误",
			Success: false,
		}
		c.ServeJSON()
	}
}
