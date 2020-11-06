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
			dial, err := redis.Dial("tcp", "172.17.18.26:6379",
				redis.DialDatabase(2), redis.DialPassword("Szsti@1109"))
			if err != nil {
				fmt.Println("Connect to redis failed ,cause by >>>", err)
			}
			return dial, err
		},
	}
}

func main() {

	conn := Pool.Get()
	//写入值{"test-Key":"test-Value"}
	_, err := conn.Do("SET", "test-Key", "test-Value", "EX", "5")
	if err != nil {
		fmt.Println("redis set value failed >>>", err)
	}
	//检查是否存在key值
	_, err = redis.Bool(conn.Do("EXISTS", "test-Key"))
	if err != nil {
		fmt.Println("illegal exception")
	}

	//read value
	v, err := redis.String(conn.Do("GET", "test-Key"))
	if err != nil {
		fmt.Println("redis get value failed >>>", err)
	}
	fmt.Println("get value: ", v)

	//del kv
	_, err = conn.Do("DEL", "test-Key")
	if err != nil {
		fmt.Println("redis delelte value failed >>>", err)
	}

	res, err := conn.Do("HSET", "student", "name", "jack")
	fmt.Println(res, err)
	res1, err := redis.String(conn.Do("HGET", "student", "name"))
	fmt.Printf("res:%s,error:%v", res1, err)

}
