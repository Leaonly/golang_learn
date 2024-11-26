package processes

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
)

type UserProcess struct {
}

func (wow *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1.连接到服务端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("NET.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal ERR=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal ERR=", err)
		return
	}

	//处理服务端返回
	//创建一个Transfer实例 ，可以做全局
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}
	mes, err = tf.ReadPkg() //mes就是registerResMes
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if registerResMes.Code == 200 {
		fmt.Println("注册成功！")
		//os.Exit(0)
	} else {
		fmt.Println("WARN: ", registerResMes.Error)
		//os.Exit(0)
	}
	return
}

func (wow *UserProcess) Login(userId int, userPwd string) (err error) {

	//1.连接到服务端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("NET.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal ERR=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal ERR=", err)
		return
	}
	//7.这时候data就是我们要发送的消息
	//7.1 把data长度发送给服务器, data长度转换表示长度的为byte切片
	//var pkgLen uint32
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err= ", err)
		return
	}

	fmt.Printf("客户端， 发送消息长度=%d 内容=%s\n", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err= ", err)
		return
	}

	//处理服务端返回
	//创建一个Transfer实例 ，可以做全局
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg() err=", err)
		return
	}
	//将mes的Data部分反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//显示当前在线用户列表，遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下: ")
		for _, v := range loginResMes.UsersId {
			//如果要求不显示自己在线
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成客户端onlineUsers初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		//开启协程，客户端与服务端通信通道
		go serverProcessMes(conn)
		//1.显示登录成功后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

// 登出方法
func (wow *UserProcess) Logout(userid int) (err error) {
	//1.连接到服务端
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("NET.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//1.创建Mes
	var mes message.Message
	mes.Type = message.LogoutMesType

	//2.创建一个离线的notifyUserStatusMes实例
	var logoutMes message.LogoutMes
	logoutMes.Status = message.UserOffline //内容
	logoutMes.UserId = userid

	//3.序列化
	data, err := json.Marshal(logoutMes)
	if err != nil {
		fmt.Println("Logout json.Marshal fail =", err.Error())
		return
	}
	mes.Data = string(data)

	//4.对mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("Logout json.Marshal fail =", err.Error())
		return
	}

	//5.发送mes给服务器
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Logout WritePkg fail =", err.Error())
		return
	}
	// 确保连接在发送完登出消息后关闭
	conn.Close()
	return
}
