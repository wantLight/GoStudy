package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var Pool redis.Pool

func init() { //init 用于初始化一些参数，先于main执行
	Pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		/**
		Redis数据库：172.17.18.26 6379
		*/
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial("tcp", "172.17.18.26:6379")
			if err != nil {
				fmt.Println("尝试赋值密码。。")
				if _, err := dial.Do("AUTH", ""); err != nil {
					dial.Close()
					fmt.Println(err)
				}
			}
			return dial, err
		},
	}
}

func main() {

	conn := Pool.Get()
	res, err := conn.Do("HSET", "student", "name", "jack")
	fmt.Println(res, err)
	res1, err := redis.String(conn.Do("HGET", "student", "name"))
	fmt.Printf("res:%s,error:%v", res1, err)

}
