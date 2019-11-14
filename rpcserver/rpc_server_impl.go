package rpcserver

import (
	"HelloApp_Api/models"
	. "HelloApp_Api/rpcprotos"
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type GoRpcServerInterfaceServerImpl struct {
	UnimplementedGoRpcServerInterfaceServer
}

//登录IM
func (*GoRpcServerInterfaceServerImpl) IMLogin(ctx context.Context, req *LoginParam) (*GoBaseResult, error) {
	if req == nil {
		return &GoBaseResult{
			Code:    -100,
			Message: "参数错误",
		}, nil
	} else {
		account := req.Account
		fmt.Println("IM登录，用户名：" + account)
		//password := req.Password
		//检查用户名是否存在
		if models.CheckUsernameExist(orm.NewOrm(), account) {
			return &GoBaseResult{
				Code:    200,
				Message: "登陆成功",
			}, nil
		} else {
			return &GoBaseResult{Code: -100, Message: "用户名或密码错误"}, nil
		}
	}
}

//ServerToClientMsgResult
func (*GoRpcServerInterfaceServerImpl) ServerToClientMessageResult(ctx context.Context, req *ServerToClientMsgResultParam) (*GoBaseResult, error) {
	if req == nil {
		return &GoBaseResult{
			Code:    -100,
			Message: "参数错误",
		}, nil
	} else {
		msgId := req.MsgId
		status := models.MessageStatus_AlreadySend
		if req.Code != 200 {
			status = models.MessageStatus_WaitSend
		}
		success := models.UpdateMessageStatus(orm.NewOrm(), int(msgId), status)
		if success {
			return &GoBaseResult{
				Code:    200,
				Message: "更新状态成功",
			}, nil
		} else {
			return &GoBaseResult{
				Code:    -100,
				Message: "更新状态失败",
			}, nil
		}
	}
}
