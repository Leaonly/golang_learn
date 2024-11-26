package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
)

type SmsProcess struct {
}

// 发送群聊的消息
func (wow *SmsProcess) SendGroupMes(content string) (err error) {
	//1.创建Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content //内容
	//短消息的用户id 等于 登录时初始化CurUser当前用户的id
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}
	mes.Data = string(data)

	//4.对mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}

	//5.发送mes给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes WritePkg fail =", err.Error())
		return
	}
	return
}

// 发送群聊的消息
func (wow *SmsProcess) SendPrivateMes(content string, id int) (err error) {
	//1.创建Mes
	var mes message.Message
	mes.Type = message.SendPrivateMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content //内容
	//短消息的用户id 等于 登录时初始化CurUser当前用户的id
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	smsMes.TargetId = id

	//3.序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendPrivateMes json.Marshal fail =", err.Error())
		return
	}
	mes.Data = string(data)

	//4.对mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendPrivateMes json.Marshal fail =", err.Error())
		return
	}

	//5.发送mes给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendPrivateMes WritePkg fail =", err.Error())
		return
	}
	return
}

// 查询群聊的消息
func (wow *SmsProcess) GetGroupMes() (err error) {
	//1.创建Mes
	var mes message.Message
	mes.Type = message.GetGroupMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	//短消息的用户id 等于 登录时初始化CurUser当前用户的id
	smsMes.UserId = CurUser.UserId

	//3.序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendPrivateMes json.Marshal fail =", err.Error())
		return
	}
	mes.Data = string(data)

	//4.对mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendPrivateMes json.Marshal fail =", err.Error())
		return
	}

	//5.发送mes给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendPrivateMes WritePkg fail =", err.Error())
		return
	}
	return
}
