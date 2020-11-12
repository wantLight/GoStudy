package routers

import (
	"github.com/astaxie/beego"
	"html/template"
	"myproject/controllers"
	"net/http"
)

func init() {
	beego.ErrorHandler("404", page_not_found)

	// 一个固定的路由，一个控制器，然后根据用户请求方法不同请求控制器中对应的方法，典型的 RESTful 方式
	beego.Router("/", &controllers.MainController{})

	/**
	使用第三个参数，第三个参数就是用来设置对应 method 到函数名，定义如下
	*表示任意的 method 都执行该函数
	使用 httpmethod:funcname 格式来展示
	多个不同的格式使用 ; 分割
	多个 method 对应同一个 funcname，method 之间通过 , 来分割
	*/
	beego.Router("/my", &controllers.MyController{}, "get:GetMy")
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
