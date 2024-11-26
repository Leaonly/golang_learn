package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"go_code/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
	//表示该conn是哪个用户的字段
	UserId int
}

// 通知所有在线用户的方法
func (wow *UserProcess) NotifyOthersOnlineUser(userId int, status int) {
	//遍历userMgr的在线用户 onlineUsers切片 ，发送id给 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {

		if status == 0 {
			//过滤自己
			if id == userId {
				continue
			}
			//通知在线
			up.NotifyMeOnline(userId)
		} else if status == 1 {
			//过滤自己
			if id == wow.UserId {
				continue
			}
			//通知离线
			up.NotifyMeOffline(userId)
		}

	}
}

// 在线用户的通知子方法(在线)
func (wow *UserProcess) NotifyMeOnline(userId int) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给data
	mes.Data = string(data)

	//对mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送,创建Transfer实例
	tf := &utils.Transfer{
		Conn: wow.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err= ", err)
		return
	}
}

// 在线用户的通知子方法(离线)
func (wow *UserProcess) NotifyMeOffline(userId int) {
	//组装LogoutResMes
	var mes message.Message
	mes.Type = message.LogoutResMesType

	var LogoutResMes message.LogoutResMes
	LogoutResMes.UserId = userId
	LogoutResMes.Status = message.UserOffline

	//将LogoutResMes序列化
	data, err := json.Marshal(LogoutResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将序列化后的LogoutResMes赋值给data
	mes.Data = string(data)

	//对mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送,创建Transfer实例
	tf := &utils.Transfer{
		Conn: wow.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOffline err= ", err)
		return
	}
}

// 处理注册请求函数
func (wow *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mes 取出 mes.Data，并反序列化为 RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//2.先声明一个resMes，返回给客户端
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//需要到redis数据库完成验证
	//4.使用model.MyUserDao去redis验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {

		if err == model.ErrUserExists {
			registerResMes.Code = 400
			registerResMes.Error = model.ErrUserExists.Error()
			println("server userprocess Register err =", registerResMes.Error)
		} else {
			registerResMes.Code = 500
			registerResMes.Error = "服务器内部错误"
			println("server userprocess Register err =", registerResMes.Error)
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功!")
	}

	println("server userprocess Res code =", registerResMes.Code)

	//5.将registerResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fial", err)
		return
	}

	//6.将data赋值给resMes
	resMes.Data = string(data)

	//7.对resMes进行序列化准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fial", err)
		return
	}

	//8.发送data给客户端， 封装到writePkg函数
	//因为使用分层模式(mvc)，先创建一个Transfer结构体实例再读取
	tf := &utils.Transfer{
		Conn: wow.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 处理登录请求函数
func (wow *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1.先从mes取出mes.Data，并反序列化为LoginMes
	var LoginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &LoginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个LoginResMes，并赋值
	var loginResMes message.LoginResMes

	//需要到redis数据库完成验证
	//1.使用model.MyUserDao去redis验证
	user, err := model.MyUserDao.Login(LoginMes.UserId, LoginMes.UserPwd)
	if err != nil {

		if err == model.ErrUserNotExists {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ErrUserPwd {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		//将登陆成功的用户消息userId 赋予给wow指针的UserId
		wow.UserId = LoginMes.UserId
		//因为登录成功，把用户放到UserMgr中, 管理在线列表用
		userMgr.AddOnlineUser(wow)
		//通知其他在线用户该用户上线
		wow.NotifyOthersOnlineUser(LoginMes.UserId, message.UserOnline)
		//将当前在线用户的id放入到loginResMes.UsersId
		//遍历 userMgr.onlineUsers
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登录成功!")
	}

	//3.将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fial", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes进行序列化准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fial", err)
		return
	}

	//6.发送data， 封装到writePkg函数
	//因为使用分层模式(mvc)，先创建一个Transfer结构体实例再读取
	tf := &utils.Transfer{
		Conn: wow.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 处理退出请求函数
func (wow *UserProcess) ServerProcessLogout(mes *message.Message) (err error) {
	//1.先从mes取出mes.Data，并反序列化为LogoutMes
	var logoutMes message.LogoutMes
	err = json.Unmarshal([]byte(mes.Data), &logoutMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//将退出成功的用户消息userId 赋予给wow指针的UserId
	wow.UserId = logoutMes.UserId
	//把用户从UserMgr的map移除
	userMgr.DelOnlineUser(logoutMes.UserId)

	//通知其他在线用户该用户离线
	wow.NotifyOthersOnlineUser(logoutMes.UserId, message.UserOffline)
	fmt.Println(logoutMes.UserId, "用户已退出")
	// 优雅地关闭连接
	wow.Conn.Close()
	return nil
}
