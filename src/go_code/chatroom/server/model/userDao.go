package model

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"

	"github.com/gomodule/redigo/redis"
)

// 在服务器启动后，初始化一个全局UserDao实例
var (
	MyUserDao *UserDao
)

// 定义一个UserDao结构体，完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 1.根据用户id 返回一个User实例+err
func (wow *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	//通过给定id去redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误
		if err == redis.ErrNil { //表示在users哈希重没有找到id
			err = ErrUserNotExists
		}
		return
	}
	user = &User{}
	//这里需要把res反序列化称User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 完成登录的校验Login
// 1.Login完成对用户的验证
// 2.如果用户验证正确，则返回一个user实例
// 3.如果用户验证错误，则返回对应错误
func (wow *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//从UserDao连接池取出一个连接
	conn := wow.pool.Get()
	defer conn.Close()
	user, err = wow.getUserById(conn, userId)
	if err != nil {
		return
	}

	//用户id获取到，继续验证密码
	if user.UserPwd != userPwd {
		err = ErrUserPwd
		return
	}
	return
}

func (wow *UserDao) Register(user *message.User) (err error) {

	//从UserDao连接池取出一个连接
	conn := wow.pool.Get()
	defer conn.Close()
	_, err = wow.getUserById(conn, user.UserId)
	if err == nil {
		err = ErrUserExists
		return
	}

	//这时说明id在redis还没有，则开始注册
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//入库
	_, err = conn.Do("Hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("注册用户错误 err= ", err)
		return
	}
	return
}
