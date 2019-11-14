package resparams

import (
	"HelloApp_Api/constants"
)

//登录返回用户信息
type LoginResult struct {
	constants.Result
	Data ResultUser `json:"data"`
}

//查找用户返回信息
type SearchUserResult struct {
	constants.Result
	Data []ResultSearchUser `json:"data"`
}

//"Id": 1,
//"Username": "Hello_mlin7buc",
//"RealName": "",
//"Gender": 0,
//"AuthInfos": [
//{
//"Id": 1,
//"User": null,
//"AuthType": 1,
//"Account": "18408252860",
//"Secret": "",
//"UpdateTime": "2019-06-11T15:33:32.0758337+08:00"
//}
//],
//"CreateTime": "2019-06-11T15:33:32.0758337+08:00",
//"phone": "18408252860",
//"qq": ""
type ResultUser struct {
	UserId   int    `json:"userId"`   //用户Id
	Username string `json:"username"` //用户名
	Nickname string `json:"nickname"` //昵称
	RealName string `json:"realName"` //真实姓名
	Gender   int8    `json:"gender"`   //性别
	Phone    string `json:"phone"`    //手机号
	QQ       string `json:"qq"`       //qq
}

//查找用户返回用户数据
type ResultSearchUser struct {
	UserId   int    `json:"userId"`   //用户Id
	Username string `json:"username"` //用户名
	Nickname string `json:"nickname"` //昵称
	Gender   int8    `json:"gender"`   //性别
}

