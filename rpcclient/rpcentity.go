package rpcclient

type RpcMsgEntity struct {
	MsgType      int16  `json:"msgType"`      //Command
	FromUserId   int    `json:"fromUserId"`   //消息发送方Id
	FromUsername string `json:"fromUsername"` //消息发送方Username
	ToUserId     int    `json:"toUserId"`     //消息接收方Id
	ToUsername   string `json:"toUsername"`   //消息接收方Username
	Content      string `json:"content"`      //消息内容
}
