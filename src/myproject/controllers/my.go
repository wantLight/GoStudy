package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"strconv"
	"time"
)

// MainController 自动拥有了所有 beego.Controller 的方法
type MyController struct {
	beego.Controller
}

type MyHello struct {
	// 这个首字母要大写不然不序列化
	Name string
	Age  int
}

type UserInfo struct {
	Username   string    `json:"name"`
	UserId     int       `json:"id"`
	CreateTime time.Time `json:"time"`
	MyTime     string
}

// 我们可以通过重写的方式来实现这些方法
func (c *MyController) GetMy() {
	// 用户可以通过 this.Ctx.Request 获取信息
	id := c.Input().Get("id")
	intid, _ := strconv.Atoi(id)
	fmt.Printf("到了！", intid)

	// 开始查询
	o := orm.NewOrm()
	var users []UserInfo
	_, err := o.Raw("SELECT * FROM sys_user").QueryRows(&users)
	if err == nil {
		fmt.Println("err!!!: ", err)
	}
	fmt.Println("------------------------------")
	// 这里不能用for range 循环修改
	for i := 0; i < len(users); i++ {
		// 判定时间不为空
		if !users[i].CreateTime.IsZero() {
			users[i].MyTime = users[i].CreateTime.Format("2006-01-02 15:04:05")
		}

	}
	//mystruct := &MyHello{"啦啦啦啦", 666}

	fmt.Println("------------------------------11111111111111")

	c.Data["json"] = users
	c.ServeJSON()

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}

func (c *MyController) UpdateMy() {
	o := orm.NewOrm()
	var s models.Store
	s.Title = "芜湖，go起飞！"
	s.Created = time.Now()
	s.Views = 2
	s.TopicTime = time.Now()

	myid, err := o.Insert(&s)
	if err == nil {
		fmt.Println(err.Error())
	}
	fmt.Print(myid, "-----------------------")
}
