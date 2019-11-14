package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//消息类型
const (
	MessageType_SYSTEM         = 0 //系统消息
	MessageType_ADD_FRIEND_REQ = 1 //添加好友请求消息
	MessageType_ADD_FRIEND_RES = 2 //添加好友响应消息
	MessageType_CHAT           = 3 //聊天消息
)

//消息状态
const (
	MessageStatus_Expire          int8 = 0 //消息过期
	MessageStatus_WaitSend        int8 = 1 //消息待发送
	MessageStatus_SendWaitConfirm int8 = 2 //消息已发送，待却确定
	MessageStatus_AlreadySend     int8 = 3 //消息已发送
)

const TABLE_NAME_MESSAGE = "message"

/**
	消息模型
 */
type Message struct {
	Id         int                    //主键Id
	FromUserId int                    //消息发送方的用户Id
	ToUserId   int                    //消息接收方用户Id
	Type       int8                   //消息类型
	Content    string `orm:"null"`    //消息内容
	Status     int8                   //消息状态
	CreateTime time.Time              //消息创建时间
	SendTime   time.Time `orm:"null"` //消息发送时间
}

//设置表名
func (m *Message) TableName() string {
	return TABLE_NAME_MESSAGE;
}

//添加Message
func AddMessage(o orm.Ormer, fromUserId int, toUserId int, msgType int8, content string) (bool, string) {
	msg := Message{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		Type:       msgType,
		Content:    content,
		CreateTime: time.Now(),
		Status:     MessageStatus_WaitSend,
	}
	_, err := o.Insert(&msg)
	if err == nil {
		return true, "提交成功"
	} else {
		return false, "提交失败"
	}
}

//获取消息列表
func GetMessageList(o orm.Ormer, msgStatus int8) []Message {
	msgList := []Message{}
	qu := o.QueryTable(TABLE_NAME_MESSAGE)
	qu.Filter("status", msgStatus).All(&msgList)
	return msgList
}

//更新Message状态
func UpdateMessageStatus(o orm.Ormer, msgId int, status int8) bool {
	message := Message{
		Id:     msgId,
		Status: status,
	}
	var err error = nil
	if status == MessageStatus_AlreadySend { //更新发送时间
		message.SendTime = time.Now()
		_, err = o.Update(&message, "status", "send_time")
	} else {
		_, err = o.Update(&message, "status")
	}
	if err == nil {
		return true
	}
	return false
}
