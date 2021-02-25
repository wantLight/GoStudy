package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	// HandlerFunc调用之后的结果，这个类型默认就实现了ServeHTTP这个接口
	//http.HandleFunc("/", sayhelloName) //设置访问的路由
	//// 默认nil即为http.DefaultServeMux，通过DefaultServeMux.ServeHTTP函数来进行调度，遍历之前存储的map路由信息，和用户访问的URL进行匹配，以查询对应注册的处理函数
	//err := http.ListenAndServe(":9090", nil) //设置监听的端口
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	x := 1
	p := x         // p, of type *int, points to x
	fmt.Println(p) // "1"
	p = 2          // equivalent to x = 2
	fmt.Println(x) // "2"
}
