package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myproject/models"
	_ "myproject/routers"
)

//引入数据模型
func init() {
	// 注册数据库
	models.RegisterDB()
}

// 使用 Bee run 开启服务
func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
	// 运行时
	beego.Run()
}
