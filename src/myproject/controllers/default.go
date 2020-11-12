package controllers

import (
	"github.com/astaxie/beego"
)

// MainController 自动拥有了所有 beego.Controller 的方法
type MainController struct {
	beego.Controller
}

// 我们可以通过重写的方式来实现这些方法
func (c *MainController) Get() {
	//  this.Ctx.WriteString("hello")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
