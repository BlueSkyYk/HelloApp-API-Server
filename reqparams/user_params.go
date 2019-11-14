package reqparams

//登录
type LoginParam struct {
	LoginType  int    `json:"loginType"` // 密码登录1、短信登录2、QQ登录3
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	SmsCode    string `json:"smsCode"`
	QQAuthCode string `json:"qqAuthCode"`
}

//查找用户
type SearchUser struct {
	Keyword string `json:"keyword"` //关键字
}
