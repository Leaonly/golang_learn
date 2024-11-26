package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

// 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("------------登录成功，进入聊天系统------------")
	fmt.Println("\t 1 显示在线用户列表")
	fmt.Println("\t 2 群发消息")
	fmt.Println("\t 3 私聊消息")
	fmt.Println("\t 4 消息列表")
	fmt.Println("\t 5 退出系统")
	fmt.Println("\t 请选择(1-5): ")
	var key int
	fmt.Scanf("%d\n", &key)

	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		// 输入消息群发
		inputGroupMes()
	case 3:
		// 输入私聊消息
		inputPrivateChatMes()
	case 4:
		//查询群消息
		sp := &SmsProcess{}
		sp.GetGroupMes()
	case 5:
		up := &UserProcess{}
		up.Logout(CurUser.UserId)
		fmt.Println("已退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重新输入。")
		fmt.Println("")
	}

}

// 和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不听读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取...")
		//连接获取消息
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//如果读取到，next逻辑
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//1.取出NotifyUserStatusMes 的id和状态
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2.把这个用户信息状态保存到客户端map[int]User中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			//群发消息
			outputMes(&mes)
		case message.SendPrivateMesType:
			//私聊消息
			outputMes(&mes)
		case message.GetGroupMesType:
			//查询群消息
			outputGroupMes(&mes)
		case message.LogoutResMesType:
			var logoutResMes message.LogoutResMes
			json.Unmarshal([]byte(mes.Data), &logoutResMes)
			updateUserStatusOffline(&logoutResMes)
		default:
			fmt.Println("未知服务端消息类型")
		}
	}
}
