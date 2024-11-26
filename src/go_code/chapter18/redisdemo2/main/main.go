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
	_, err = conn.Do("Hset", "user01", "name", "john")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	_, err = conn.Do("Hset", "user01", "age", "18")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//3.读取数据
	//因为返回r是interface{}，需要转换对应类型
	r1, err := redis.String(conn.Do("Hget", "user01", "name"))
	if err != nil {
		fmt.Println("Hget err=", err)
		return
	}
	r2, err := redis.String(conn.Do("Hget", "user01", "age"))
	if err != nil {
		fmt.Println("Hget err=", err)
		return
	}

	fmt.Println("1redis do success...", r1)
	fmt.Println("2redis do success...", r2)
}
