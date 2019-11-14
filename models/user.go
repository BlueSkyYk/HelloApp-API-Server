package models

import (
	"HelloApp_Api/constants"
	"HelloApp_Api/utils"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	//性别
	Gender_Unknown = 0 //未知
	Gender_Male    = 1 //男性
	Gender_Female  = 2 //女性
)

const TABLE_NAME_USER = "user"

//实体类
type User struct {
	Id          int                                      //id
	Username    string             `orm:"null"`          //用户名
	Nickname    string             `orm:"null"`          //用户昵称
	RealName    string             `orm:"null"`          //真实姓名
	Gender      int8               `orm:"default(0)"`    //性别
	AuthInfos   []*UserAuth        `orm:"reverse(many)"` //验证信息一对多
	Friends     []*UserFriend      `orm:"reverse(many)"` //好友信息一对多
	FriendGroup []*UserFriendGroup `orm:"reverse(many)"` //好友分组一对多
	CreateTime  time.Time                                //创建时间
}

//设置表名
func (u *User) TableName() string {
	return TABLE_NAME_USER;
}

//注册新用户
func CreateNewUser(o orm.Ormer) *User {
	//创建用户
	user := User{
		Gender:     Gender_Unknown,
		CreateTime: time.Now(),
		Username:   "Hello_" + utils.GetRandomString(8),
	}
	//保存用户
	_, err1 := o.Insert(&user)

	//创建默认好友分组
	userFriendGroup, err2 := CreateNewFriendGroup(o, &user, "我的好友", true)

	if err1 == nil && err2 == nil {
		user.FriendGroup = []*UserFriendGroup{userFriendGroup}
		return &user
	} else {
		return nil
	}
}

//查找用户
func SearchUser(o orm.Ormer, keyWord string) (int, string, *[]User, bool) {
	//检查
	if keyWord == "" {
		return constants.RESULT_ERROR, "关键字不能为空", nil, false;
	}
	searchUsers := []User{}
	//检查是否是手机号
	if utils.CheckPhone(keyWord) {
		user, err := checkUserAuthExist(o, keyWord)
		if err == nil {
			searchUsers = append(searchUsers, *user)
		}
	}
	//添加用户名搜索
	users := []User{}
	qs := o.QueryTable(TABLE_NAME_USER)
	_, err := qs.Filter("username__icontains", keyWord).All(&users)
	if err == nil && len(users) > 0 {
		searchUsers = append(searchUsers, users...)
	}
	//添加昵称搜索
	users = []User{}
	_, err = qs.Filter("nickname__icontains", keyWord).All(&users)
	if err == nil && len(users) > 0 {
		searchUsers = append(searchUsers, users...)
	}
	return constants.RESULT_OK, "返回成功", &searchUsers, true
}

//根据UserId获取用户信息
func FindUserById(o orm.Ormer, userId int) *User {
	user := User{
		Id: userId,
	}
	err := o.Read(&user)
	if err == nil {
		return &user
	}
	return nil
}

//检查用户名是否存在
func CheckUsernameExist(o orm.Ormer, username string) bool {
	qs := o.QueryTable(TABLE_NAME_USER)
	return qs.Filter("username", username).Exist()
}
