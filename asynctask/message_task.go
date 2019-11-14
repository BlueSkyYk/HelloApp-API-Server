package asynctask

import (
	"HelloApp_Api/models"
	"HelloApp_Api/rpcclient"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
)

/**
	开启消息发送任务
 */
func StartMessageSendTask() {
	o := orm.NewOrm()
	task := toolbox.NewTask("MessageSendTask", "*/2 * * * * *", func() error {
		//获取Message列表
		msgList := models.GetMessageList(o, models.MessageStatus_WaitSend)
		//遍历发送Message
		if len(msgList) > 0 {
			for _, msg := range msgList {
				var fromUser *models.User
				//获取发送方用户信息
				if msg.FromUserId > 0 {
					fromUser = models.FindUserById(o, msg.ToUserId)
				}
				//获取接收方用户信息
				toUser := models.FindUserById(o, msg.ToUserId)

				//消息内容
				content := msg.Content

				//判断如果是添加好友请求，需要获取用户昵称
				if msg.Type == models.MessageType_ADD_FRIEND_REQ {
					//发送方用户昵称
					fromNickname := ""
					if fromUser != nil {
						fromNickname = fromUser.Nickname
					}
					contentMap := map[string]string{
						"verInfo":        msg.Content,
						"friendNickname": fromNickname,
					}
					bytes, e := json.Marshal(contentMap)
					if e == nil {
						content = string(bytes)
					}
				}

				//如果消息接收方不存在，则删除消息
				if toUser == nil {
					o.Delete(&msg)
					break
				}
				//创建消息Model
				messageModel := rpcclient.MessageModel{
					MsgId:      msg.Id,
					Type:       msg.Type,
					FromUserId: msg.FromUserId,
					ToUserId:   msg.ToUserId,
					Content:    content,
				}
				//调用Im服务端，发送消息，此处监听消息发送失败
				if success := rpcclient.SendMessageToUser(toUser.Username, messageModel); !success {
					fmt.Print("发送消息失败。。。。。。")
					fmt.Println(msg)
				} else {
					//更新为已发送，待确定状态
					msg.Status = models.MessageStatus_SendWaitConfirm
					o.Update(msg)
				}
			}
		}

		return nil
	})
	toolbox.AddTask("CheckMessageTask", task)
	toolbox.StartTask()
	fmt.Println("定时任务启动")
}
