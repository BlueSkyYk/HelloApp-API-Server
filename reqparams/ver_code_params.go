package reqparams

const (
	//用户登录
	CodeType_Login = 1
)

//获取短信验证码
type SmsVerCodeParam struct {
	Phone string `json:"phone"`
}
