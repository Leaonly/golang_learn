// package main

// import (
// 	"fmt"
// 	"github.com/gomodule/redigo/redis"
// )

// // 定义一个全局的pool
// var pool *redis.Pool

// //当启动程序时，就初始化连接池
// func init(){
// 	pool = &redis.Pool{
// 		MaxIdle: 8,
// 		MaxActive: 0,
// 		IdleTimeout: 100,
// 		Dial: func() (redis.Conn, error) {
// 			return redis.Dial("tcp", "10.9.20.11:6379",redis.DialPassword("fNuu7Yhs5zQXa16WfMxQ"))
// 		}
// 	}
// }

// func main(){
// 	fmt.Println("")
// 	conn := pool.Get()
// 	defer conn.Close()
// }

package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,                  // 最大空闲连接数
		MaxActive:   0,                  // 最大活动连接数，0 表示无限制
		IdleTimeout: 240 * time.Second,  // 空闲连接超时时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server,
				redis.DialPassword(password),
			)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func main() {
	// 创建连接池
	pool := newPool("10.9.20.11:6379", "fNuu7Yhs5zQXa16WfMxQ")

	// 从连接池中获取连接
	conn := pool.Get()
	defer conn.Close()

	// 写入数据
	_, err := conn.Do("mset", "age", "22", "name", "john", "job", "dev")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	// 读取数据
	values, err := redis.Values(conn.Do("mget", "name", "age", "job"))
	if err != nil {
		fmt.Println("mget err=", err)
		return
	}

	// 解析值
	var name, age, job string
	if _, err := redis.Scan(values, &name, &age, &job); err != nil {
		fmt.Println("scan err=", err)
		return
	}

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("job:", job)
}