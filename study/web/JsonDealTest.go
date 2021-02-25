package main

import (
	"encoding/json"
	"fmt"
)

// 能够被赋值的字段必须是可导出字段(即首字母大写）
type Server struct {
	// 当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将你想要的数据对应的字段名大写
	ServerName string `json:"serverName"`
	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	//var s Serverslice
	// 使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)。
	s := new(Serverslice)
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	// func Unmarshal(data []byte, v interface{}) error
	// interface{}可以用来存储任意数据类型的对象
	json.Unmarshal([]byte(str), &s)
	for i := 0; i < len(s.Servers); i++ {
		fmt.Println("ServerIP:" + s.Servers[i].ServerIP)
		fmt.Println("ServerName:" + s.Servers[i].ServerName)
	}

	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))

}
