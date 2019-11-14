package models

import (
	"HelloApp_Api/constants"
	"HelloApp_Api/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

//鉴权方式
const (
	AuthType_Phone = 1 //密码
	AuthType_QQ    = 2 //QQ
)

const TABLE_NAME_USER_AUTH = "user_auth"

//用户鉴权方式
type UserAuth struct {
	Id         int                   //id
	User       *User `orm:"rel(fk)"` //用户Id
	AuthType   int8                   //鉴权类型
	Account    string `orm:"null"`   //账号（账号）
	Secret     string `orm:"null"`   //密钥（密码）
	UpdateTime time.Time             //更新时间
}

func (ua *UserAuth) TableName() string {
	return TABLE_NAME_USER_AUTH
}

//密码登录
func LoginByPassword(o orm.Ormer, phone string, password string) (int, string, *User, bool) {
	//检查手机号码
	if phone == "" || !utils.CheckPhone(phone) {
		return constants.PHONE_ERROR, "手机号码错误", nil, false
	}
	user, err := checkUserAuthExist(o, phone)
	if err != nil {
		return constants.RESULT_ERROR, "登录失败", nil, false
	}
	if user != nil { //用户存在
		for _, auth := range user.AuthInfos {
			if auth.AuthType == AuthType_Phone { //找到手机号认证
				if auth.Secret == "" { //密码为空
					return constants.PASSWORD_ERROR, "未设置密码，请使用短信验证码登录", nil, false
				} else if auth.Secret == utils.GetStrMd5(password) { //密码加密后判断
					return constants.RESULT_OK, "登录成功", user, true
				} else { //密码错误
					return constants.PASSWORD_ERROR, "密码错误", nil, false
				}
			}
		}
		//查找鉴权信息出错
		return constants.RESULT_ERROR, "登录失败", nil, false
	} else {
		return constants.RESULT_ERROR, "用户不存在，请先注册", nil, false
	}
}

//短信验证码登录
func LoginBySmsCode(o orm.Ormer, phone string, smsCode string) (int, string, *User, bool) {
	//检查手机号码
	if phone == "" || !utils.CheckPhone(phone) {
		return constants.PHONE_ERROR, "手机号码错误", nil, false
	}
	//检查短信验证码
	if smsCode == "" {
		return constants.SMS_CODE_ERROR, "验证码错误", nil, false
	}
	//验证短信验证码
	if err := utils.VerSmsCode(phone, smsCode); err != nil {
		return constants.SMS_CODE_ERROR, err.Error(), nil, false
	} else { //短信验证码发送成功
		//检查用户是否存在
		user, err := checkUserAuthExist(o, phone)
		if err != nil {
			return constants.RESULT_ERROR, "登录失败", nil, false
		} else if user != nil { //用户存在
			fmt.Println("登录成功", user)
			return constants.RESULT_OK, "登录成功", user, true
		} else { //用户不存在，创建用户
			user, err := CreateUserAuth(o, phone, "", AuthType_Phone)
			if err == nil {
				return constants.RESULT_OK, "登录成功", user, true
			} else {
				return constants.RESULT_ERROR, "登录失败", nil, false
			}
		}
	}
}

//qq登录
func LoginByQQ() *UserAuth {
	return nil;
}

//创建用户
func CreateUserAuth(o orm.Ormer, account string, secret string, authType int8) (*User, error) {
	userAuth := new(UserAuth)
	qs := o.QueryTable(TABLE_NAME_USER_AUTH)
	//根据Account查询认证信息
	err := qs.Filter("account", account).One(userAuth)
	//认证信息存在
	if err == nil { //判断用户认证信息是否存在
		return userAuth.User, nil
	} else { //用户认证信息不存在，创建用户，并且添加认证信息
		//开启事务
		o.Begin()
		//创建信用
		user := CreateNewUser(o)
		if user == nil {
			o.Rollback()
			return nil, errors.New("创建用户失败")
		}
		//创建用户授权信息
		userAuth := UserAuth{
			Account:    account,
			Secret:     secret,
			AuthType:   authType,
			User:       user,
			UpdateTime: time.Now(),
		}
		//保存认证信息
		_, err2 := o.Insert(&userAuth)
		if err2 == nil {
			o.Commit()
			userAuth.User = nil
			authInfos := [] *UserAuth{&userAuth}
			user.AuthInfos = authInfos
			return user, nil
		} else {
			o.Rollback()
			return nil, errors.New("创建用户失败")
		}
	}
}

//判断用户认证是否存在
func checkUserAuthExist(o orm.Ormer, account string) (*User, error) {
	userAuth := new(UserAuth)
	qs := o.QueryTable(TABLE_NAME_USER_AUTH)
	//根据Account查询认证信息
	err := qs.Filter("account", account).One(userAuth)
	if err != nil || userAuth == nil { //用户不存在
		return nil, nil
	} else {
		user := userAuth.User
		o.Read(user)
		userAuth.User = nil
		authInfos := []*UserAuth{userAuth}
		user.AuthInfos = authInfos
		return user, nil
	}
}
