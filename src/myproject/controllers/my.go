package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myproject/models"
	"strconv"
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

// 我们可以通过重写的方式来实现这些方法
func (c *MyController) GetMy() {
	// 用户可以通过 this.Ctx.Request 获取信息
	id := c.Input().Get("id")
	intid, _ := strconv.Atoi(id)

	fmt.Printf("到了！", intid)

	o := orm.NewOrm()
	var s models.Store
	s.Title = "测试呀呀呀"

	id, err := o.Insert(&s)
	if err == nil {
		fmt.Println(id)
	}

	mystruct := &MyHello{"啦啦啦啦", 666}

	fmt.Print(mystruct)

	c.Data["json"] = mystruct
	c.ServeJSON()

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}
