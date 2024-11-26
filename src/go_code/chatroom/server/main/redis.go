package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// 定义一个全局pool
var pool *redis.Pool

func initPool(address string, redisPwd string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	//当启动程序时，就初始化连接池
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address, redis.DialPassword(redisPwd))
		},
	}

}
