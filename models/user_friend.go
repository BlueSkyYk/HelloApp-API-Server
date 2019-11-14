package models

import (
	"HelloApp_Api/constants"
	"github.com/astaxie/beego/orm"
	"time"
)

const TABLE_NAME_USER_FRIEND = "user_friend"

//实体类
type UserFriend struct {
	Id         int                              //主键ID
	User       *User            `orm:"rel(fk)"` //用户
	FriendUser *User            `orm:"rel(fk)"` //好友
	Group      *UserFriendGroup `orm:"rel(fk)"` //分组
	CreateTime time.Time                        //更新时间
}

//设置表名
func (uf *UserFriend) TableName() string {
	return TABLE_NAME_USER_FRIEND
}

//创建好友关系
func CreateUserFriend(o orm.Ormer, userId int, friendId int) bool {
	//检查用户
	user := FindUserById(o, userId)
	if user == nil {
		return false
	}
	//检查好友用户
	friend := FindUserById(o, friendId)
	if friend == nil {
		return false
	}
	o.Begin()
	//判断正向好友关系
	if !checkFriend(o, userId, friendId) {
		userFriend := UserFriend{
			User:       user,
			FriendUser: friend,
			CreateTime: time.Now(),
		}
		_, err := o.Insert(&userFriend)
		if err != nil {
			o.Rollback()
			return false
		}
	}
	//判断反向好友关系
	if !checkFriend(o, friendId, user.Id) {
		userFriend := UserFriend{
			User:       user,
			FriendUser: friend,
			CreateTime: time.Now(),
		}
		_, err := o.Insert(&userFriend)
		if err != nil {
			o.Rollback()
			return false
		}
	}
	o.Begin()
	return true
}

//获取好友列表
func GetUserFriendList(o orm.Ormer, userId int) (int, string, []*UserFriendGroup, bool) {
	groupList, err := GetUserFriendGroupList(o, userId)
	if err != nil {
		return constants.RESULT_ERROR, err.Error(), nil, false
	}

	for _, group := range groupList {
		//读取好友列表
		o.LoadRelated(group, "Friends")
		if group.Friends != nil && len(group.Friends) > 0 {
			for _, friend := range group.Friends {
				//读取好友用户
				o.LoadRelated(friend, "FriendUser")
			}
		}
	}
	return constants.RESULT_OK, "返回成功", groupList, true
}

//检查好友是否存在
func checkFriend(o orm.Ormer, userId int, friendId int) bool {
	userFriend := UserFriend{}
	qs := o.QueryTable(TABLE_NAME_USER_FRIEND)
	err := qs.Filter("user_id", userId).Filter("friend_user_id", friendId).One(&userFriend)
	if err == nil {
		return true
	} else {
		return false
	}
}
