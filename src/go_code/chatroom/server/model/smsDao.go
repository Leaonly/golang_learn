package model

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"

	"github.com/gomodule/redigo/redis"
)

// 在服务器启动后，初始化一个全局SmsDao实例
var (
	MySmsDao *SmsDao
)

// 定义一个SmsDao结构体，完成各种方法操作
type SmsDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建一个SmsDao实例
func NewSmsDao(pool *redis.Pool) (smsDao *SmsDao) {

	smsDao = &SmsDao{
		pool: pool,
	}
	return
}

// 群消息入库Dao
func (wow *SmsDao) GroupMesDao(sms *message.SmsMes) (err error) {

	//从SmsDao连接池取出一个连接
	conn := wow.pool.Get()
	defer conn.Close()

	//将整个SmsMes结构体序列化为JSON字符串
	data, err := json.Marshal(sms)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//格式化消息时间
	mesTime := sms.MesDate.Format("2006-01-02T15:04:05Z07:00")
	//入库
	_, err = conn.Do("Hset", "group1mes", mesTime, string(data))
	if err != nil {
		fmt.Println("群聊消息入库错误 err= ", err)
		return
	}
	return
}

// 群消息查询
func (wow *SmsDao) GetGroupMesDao() (result map[string]string, err error) {

	//从SmsDao连接池取出一个连接
	conn := wow.pool.Get()
	defer conn.Close()

	//去redis查询群消息
	values, err := redis.Values(conn.Do("HGetall", "group1mes"))
	if err != nil {
		//错误
		if err == redis.ErrNil {
			err = ErrGroupHget
		}
		return
	}

	// 将values转换为map[string]string
	result, err = redis.StringMap(values, err)
	if err != nil {
		fmt.Println("StringMap err=", err)
		return
	}
	return
}
