package resutil

import (
	"HelloApp_Api/models"
	. "HelloApp_Api/resparams"
)

/**
	创建返回的User
 */
func CreateResultUser(user models.User) ResultUser {
	phone := ""
	qq := ""

	resultUser := new(ResultUser)

	//遍历查找手机号和qq
	if user.AuthInfos != nil {
		for e := range user.AuthInfos {
			ua := user.AuthInfos[e]
			switch ua.AuthType {
			case models.AuthType_Phone:
				phone = ua.Account
				break;
			case models.AuthType_QQ:
				qq = ua.Account
				break;
			}
		}
	}
	resultUser.UserId = user.Id
	resultUser.Username = user.Username
	resultUser.RealName = user.RealName
	resultUser.Gender = user.Gender
	resultUser.Phone = phone
	resultUser.QQ = qq

	return *resultUser
}

/**
	创建返回的 SearchUser
 */
func CreateResultSearchUser(users []models.User) []ResultSearchUser {
	searchUsers := []ResultSearchUser{}
	for _, user := range users {
		searchUser := ResultSearchUser{
			UserId:   user.Id,
			Username: user.Username,
			Nickname: user.Nickname,
			Gender:   user.Gender,
		}
		searchUsers = append(searchUsers, searchUser)
	}
	return searchUsers
}

/**
	创建返回的好友列表分组
 */
func CreateResultFriendGroup(groups []*models.UserFriendGroup) []ResultFriendGroup {
	resultGroupList := []ResultFriendGroup{}

	if groups != nil {
		//遍历分组
		for _, group := range groups {
			friendList := []ResultFriend{}
			resultGroup := ResultFriendGroup{
				GroupName:  group.Name,
				GroupId:    group.Id,
				FriendList: friendList,
			}
			resultGroupList = append(resultGroupList, resultGroup)

			if group.Friends != nil {
				for _, friend := range group.Friends {
					user := CreateResultUser(*friend.User)
					resultFriend := ResultFriend{
						ResultUser: user,
					}
					friendList = append(friendList, resultFriend)
				}
			}
		}
	}

	return resultGroupList
}
