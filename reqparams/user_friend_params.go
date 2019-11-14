package reqparams

/**
	添加好友请求参数
 */
type AddFriendRequestParam struct {
	UserId    int    `json:"userId"`    //用户Id
	FriendId  int    `json:"friendId"`  //好友用户Id
	ExtraInfo string `json:"extraInfo"` //附加信息
}

/**
	添加好友好友返回参数
 */
type AddFriendResultParam struct {
	UserId         int    `json:"userId"`         //用户Id（发起请求的用户Id）
	FriendId       int    `json:"friendId"`       //好友Id（好友用户Id）
	Agree          bool   `json:"agree"`          //是否同意
	ExtraInfo      string `json:"extraInfo"`      //附加数据
}

/**
	获取用户好友列表返回参数
 */
type GetFriendListParam struct {
	UserId int `json:"userId"` //用户Id
}
