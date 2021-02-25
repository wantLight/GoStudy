package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，你才能对这个表单数据进行操作
		r.ParseForm()
		if len(r.Form["username"][0]) == 0 {
			//为空的处理
			fmt.Println("username为空啦，报错！")
		}
		// 是未选中的复选框和单选按钮，则根本不会在r.Form中产生相应条目
		_, err := strconv.Atoi(r.Form.Get("password"))
		if err != nil {
			//数字转化出错了，那么可能就不是数字
			fmt.Println("password不是数字！")
		}

		if !testCheck([]string{"apple", "pear", "banana"}, r.Form.Get("fruit")) {
			fmt.Println("下拉菜单没有勾选！")
		}

		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ","))
		}

		// v := url.Values{}
	}
}

// 处理/upload 逻辑
/**
1.表单中增加enctype="multipart/form-data"
2.服务端调用r.ParseMultipartForm,把上传的文件存储在内存和临时文件中
3.使用r.FormFile获取文件句柄，然后对文件进行存储等处理。
*/
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		// 上传的文件存储在maxMemory大小的内存里面，如果文件大小超过了maxMemory，那么剩下的部分将存储在系统的临时文件中。
		r.ParseMultipartForm(32 << 20)
		// r.FormFile 获取文件句柄 同时自动解析form
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		// 直接取出来其他参数
		fmt.Println(r.Form.Get("type") + "!!!!!!!!!!!!!!!!!!!!!!!")

		// defer 确保一定回关闭流
		defer file.Close()
		// 这个写入到w的是输出到客户端的 FileHeader是一个结构体
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		// 存储文件
		io.Copy(f, file)
	}
}

/**
检查下拉菜单是否勾选啦
*/
func testCheck(mystic []string, mydata string) bool {
	slice := mystic

	//v := r.Form.Get("fruit")
	for _, item := range slice {
		if item == mydata {
			return true
		}
	}
	return false
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)   //设置访问的路由
	//设置访问的路由
	http.HandleFunc("/upload", upload)

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
