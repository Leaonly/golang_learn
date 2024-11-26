package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	//通过go 向redis读写数据
	conn, err := redis.Dial("tcp", "10.9.20.11:6379",
		redis.DialPassword("fNuu7Yhs5zQXa16WfMxQ"),
	)
	if err != nil {
		fmt.Println("resi.Dial err=", err)
		return
	}
	defer conn.Close() //重要！延时关闭
	//2.写入数据
	_, err = conn.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//3.读取数据
	//因为返回r是interface{}，需要转换对应类型
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	fmt.Println("redis do success...", r)
}
