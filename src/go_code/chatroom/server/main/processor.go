package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/processes"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

// 创建一个结构体
type Processor struct {
	Conn net.Conn
}

// 根据客户端发送消息种类不同，来调用对应函数处理
func (wow *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		//创建一个UserProcess结构体实例
		up := &processes.UserProcess{
			Conn: wow.Conn,
		}
		//不能直接使用 结构体名字.方法名 来调用方法。方法必须通过结构体的实例来调用
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &processes.UserProcess{
			Conn: wow.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//处理群发消息逻辑
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	case message.SendPrivateMesType:
		//处理私聊消息逻辑
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendPrivateMes(mes)
	case message.GetGroupMesType:
		//处理查询群聊消息逻辑
		smsProcess := &processes.SmsProcess{}
		smsProcess.GetGroupMes(mes, wow.Conn)
	case message.LogoutMesType:
		//处理登出
		up := &processes.UserProcess{
			Conn: wow.Conn,
		}
		err = up.ServerProcessLogout(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

// 根据客户端发送消息种类不同，来调用对应函数处理
func (wow *Processor) Process2() (err error) {

	//读客户端发送的信息
	for {
		//创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: wow.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出....")
				break
			} else {
				fmt.Println("ReadPkg err=", err)
				return err
			}
		}
		fmt.Println("mes=", mes)

		err = wow.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
	return
}
