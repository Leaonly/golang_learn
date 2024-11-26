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
	_, err = conn.Do("mset", "age", "22", "name", "john", "job", "dev")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//3.读取数据
	//因为返回r是interface{}，需要转换对应类型
	values, err := redis.Values(conn.Do("mget", "name", "age", "job"))
	if err != nil {
		fmt.Println("Hget err=", err)
		return
	}

	// 解析values, mget 命令返回的是一个字符串数组（即 []interface{} 类型）。因此，你需要使用 redis.Values 来处理 mget 命令的返回值
	var name, age, job string
	if _, err := redis.Scan(values, &name, &age, &job); err != nil {
		fmt.Println("scan err=", err)
		return
	}

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("job:", job)
}
