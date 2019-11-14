package main

import (
	"HelloApp_Api/asynctask"
	"HelloApp_Api/constants"
	"HelloApp_Api/ormx"
	_ "HelloApp_Api/routers"
	"HelloApp_Api/rpcclient"
	"HelloApp_Api/rpcserver"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//开启Rpc服务，使用新线程开启
	go rpcserver.StartRpcServer(constants.GO_RPC_PORT)

	//开启异步任务
	asynctask.StartMessageSendTask()

	//Swagger配置
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func test() {
	for true {
		fmt.Println("hhhhhhh")
		success := rpcclient.SendMessageToUser("hhh", rpcclient.MessageModel{
			FromUserId: 2, Content: "hhhh", Type: 2, MsgId: 2,
		})
		result := "失败"
		if success {
			result = "成功"
		}
		fmt.Println("发送数据：" + result)
	}
}

func init() {
	//初始化ORM
	ormx.InitORM()
}
