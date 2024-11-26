package processes

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"go_code/chatroom/server/utils"
	"net"
	"time"
)

// 用结构体管理方法
type SmsProcess struct {
	//暂时不用字段

}

// 处理群发消息函数
func (wow *SmsProcess) SendGroupMes(mes *message.Message) (err error) {

	//1.先从mes取出mes.Data，并反序列化为smsMes
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//当前发送消息时间
	sendTime := time.Now()
	smsMes.MesDate = sendTime
	//群消息入库做离线查看
	err = model.MySmsDao.GroupMesDao(&smsMes)
	if err != nil {
		fmt.Println("GroupMesDao fail err=", err)
		return
	}

	//序列化准备发送
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//遍历服务端onlineUsers的map，获取各用户id和连接发送
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == smsMes.UserId {
			continue
		}
		wow.SendMesOnlineUser(data, up.Conn)
	}
	return
}

// 处理私聊消息函数
func (wow *SmsProcess) SendPrivateMes(mes *message.Message) (err error) {

	//1.先从mes取出mes.Data，并反序列化为smsMes
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	// 标志变量，用于标记是否找到目标用户
	found := false
	//遍历服务端onlineUsers的map，获取各用户id和连接
	for id, up := range userMgr.onlineUsers {
		//发给目标
		if id == smsMes.TargetId {
			wow.SendMesOnlineUser(data, up.Conn)
			found = true
			break
		}
	}
	if !found {
		err = errors.New("用户不在线")
	}
	return
}

func (wow *SmsProcess) SendMesOnlineUser(data []byte, conn net.Conn) {

	//6.发送data， 封装到writePkg函数
	//因为使用分层模式(mvc)，先创建一个Transfer结构体实例再读取
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息 fail err=", err)
		return
	}
}

// 处理查询群发消息函数
func (wow *SmsProcess) GetGroupMes(mes *message.Message, conn net.Conn) (err error) {

	//2.先声明一个resMes，返回给客户端
	var resMes message.Message
	resMes.Type = message.GetGroupMesType

	res, err := model.MySmsDao.GetGroupMesDao()
	if err != nil {
		fmt.Println("MySmsDao fial", err)
		return
	}
	//3.对resMes.data进行序列化
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println("json.Marshal fial", err)
		return
	}

	//6.将data赋值给resMes
	resMes.Data = string(data)

	//5.将Mes 序列化封装
	data1, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fial", err)
		return
	}
	//6.发送data， 封装到writePkg函数
	//因为使用分层模式(mvc)，先创建一个Transfer结构体实例再读取
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data1)

	return
}
