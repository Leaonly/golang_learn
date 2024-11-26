package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// 创建一个客户端实例
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.9.20.11:6379",      // Redis服务器地址
		Password: "fNuu7Yhs5zQXa16WfMxQ", // 没有密码则为空字符串
		DB:       0,                      // 使用默认的数据库
	})

	// 测试连接是否成功
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong) // 输出: PONG
}
