package ormx

import (
	"HelloApp_Api/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

/**
初始化ORM
 */
func InitORM() {
	//获取数据库初始化参数
	address := beego.AppConfig.String("address")
	port, _ := beego.AppConfig.Int("port")
	user := beego.AppConfig.String("user")
	psw := beego.AppConfig.String("psw")
	database_dev := beego.AppConfig.String("database_dev")
	database_release := beego.AppConfig.String("database_release")

	//初始化Orm
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//ormx.RegisterDataBase("default", "mysql", "root:Ykk123456!@tcp(rm-wz9044464o655n4451o.mysql.rds.aliyuncs.com:3306)/hello_app_dev?charset=utf8")
	orm.RegisterDataBase("default", "mysql", user+":"+psw+"@tcp("+address+":"+strconv.Itoa(port)+")/"+database_dev+"?charset=utf8")
	orm.RegisterDataBase("release", "mysql", user+":"+psw+"@tcp("+address+":"+strconv.Itoa(port)+")/"+database_release+"?charset=utf8")

	//注册Model
	registerModel()

	//建表策略(使用策略，是否从新建表，)
	orm.RunSyncdb("default", false, true)
}

//注册Model
func registerModel() {
	orm.RegisterModel(
		new(models.User),
		new(models.UserAuth),
		new(models.UserFriend),
		new(models.UserFriendGroup),
		new(models.Message),
	)
}
