package resparams

import "HelloApp_Api/constants"

//获取用户好友列表
type GetFriendListResult struct {
	constants.Result
	Data []ResultFriendGroup
}

type ResultFriendGroup struct {
	GroupId    int            `json:"groupId"`
	GroupName  string         `json:"groupName"`
	FriendList []ResultFriend `json:"friendList"`
}

type ResultFriend struct {
	ResultUser
}
