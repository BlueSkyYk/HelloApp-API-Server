package controllers

import (
	"HelloApp_Api/constants"
	"HelloApp_Api/models"
	"HelloApp_Api/reqparams"
	"HelloApp_Api/resparams"
	"HelloApp_Api/resutil"
	"encoding/json"
)

type UserFriendController struct {
	BaseController
}

// @Title 添加好友请求
// @Description 通过用户Id添加好友
// @Param	body	body	reqparams.AddFriendRequestParam	true	"The AddFriendRequest content"
// @Success 200 {object} constants.BaseResult
// @Failure -100 :请求失败
// @router /addFriendRequest [post]
func (c *UserFriendController) AddFriendRequest() {
	param := reqparams.AddFriendRequestParam{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {
		//添加消息
		success, msg := models.AddMessage(c.o,
			param.UserId,
			param.FriendId,
			models.MessageStatus_WaitSend,
			param.ExtraInfo)
		if success {
			c.Data["json"] = constants.Result{
				Success: true,
				Code:    constants.RESULT_OK,
				Message: "发送成功",
			}
		} else {
			c.Data["json"] = constants.Result{
				Success: false,
				Code:    constants.RESULT_ERROR,
				Message: msg,
			}
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = constants.Result{
			Success: false,
			Code:    constants.PARAM_ERROR,
			Message: "参数错误",
		}
		c.ServeJSON()
	}
}

// @Title 添加好友结果
// @Description 好友返回添加好友结果
// @Param	body	body	reqparams.AddFriendResultParam	true	"The AddFriendResult content"
// @Success 200 {object} constants.BaseResult
// @Failure -100 :请求失败
// @router /addFriendResult [post]
func (c *UserFriendController) AddFriendResult() {
	param := reqparams.AddFriendResultParam{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {

	} else {
		c.Data["json"] = constants.Result{
			Success: false,
			Code:    constants.PARAM_ERROR,
			Message: "参数错误",
		}
		c.ServeJSON()
	}
}

// @Title 获取好友列表
// @Description 返回好友列表
// @Param	body	body	reqparams.GetFriendListParam	true	"The GetFriendListResult content"
// @Success 200 {object} resparams.GetFriendListResult
// @Failure -100 :请求失败
// @router /getUserFriendList [post]
func (c *UserFriendController) GetUserFriendList() {
	param := reqparams.GetFriendListParam{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {
		code, msg, groups, success := models.GetUserFriendList(c.o, param.UserId)
		if success {
			groupList := resutil.CreateResultFriendGroup(groups)
			c.Data["json"] = resparams.GetFriendListResult{
				Result: constants.Result{
					Success: success,
					Code:    code,
					Message: msg,
				},
				Data: groupList,
			}
			c.ServeJSON()
		} else {
			c.Data["json"] = constants.Result{
				Success: success,
				Code:    code,
				Message: msg,
			}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = constants.Result{
			Success: false,
			Code:    constants.PARAM_ERROR,
			Message: "参数错误",
		}
		c.ServeJSON()
	}
}
