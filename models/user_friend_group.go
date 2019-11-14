package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

const TABLE_NAME_USER_FRIEND_GROUP = "user_friend_group"

//好友分组
type UserFriendGroup struct {
	Id        int                                 //主键ID
	User      *User `orm:"rel(fk)"`               //用户
	Name      string                              //分组名
	IsDefault bool                                //是否是默认分组
	Friends   []*UserFriend `orm:"reverse(many)"` //好友信息一对多
}

//设置表名
func (ufg *UserFriendGroup) TableName() string {
	return TABLE_NAME_USER_FRIEND_GROUP;
}

//创建好友分组
func CreateNewFriendGroup(o orm.Ormer, user *User, groupName string, isDefault bool) (*UserFriendGroup, error) {
	//检查当前分组名是否已经存在
	if (checkFriendGroupExist(o, user.Id, groupName)) {
		return nil, errors.New("分组名已经存在")
	}
	userFriendGroup := UserFriendGroup{
		Name:      groupName,
		User:      user,
		IsDefault: isDefault,
	}
	_, err := o.Insert(&userFriendGroup)
	if err == nil {
		return &userFriendGroup, nil
	} else {
		return nil, errors.New("创建默认好友分组失败")
	}
}

//获取好友分组
func GetUserFriendGroupList(o orm.Ormer, userId int) ([]*UserFriendGroup, error) {
	qu := o.QueryTable(TABLE_NAME_USER_FRIEND_GROUP)
	groupList := []*UserFriendGroup{}
	_, err := qu.Filter("user_id", userId).All(&groupList)
	if err == nil {
		return groupList, nil
	} else {
		return nil, errors.New("获取好友分组列表失败")
	}
}

//检查用户分组是否存在
func checkFriendGroupExist(o orm.Ormer, userId int, groupName string) bool {
	qu := o.QueryTable(TABLE_NAME_USER_FRIEND_GROUP)
	return qu.Filter("user_id", userId).Filter("name", groupName).Exist()
}
