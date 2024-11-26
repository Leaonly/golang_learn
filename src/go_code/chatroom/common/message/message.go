package message

import "time"

//消息类型常量
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	SendPrivateMesType      = "SendPrivateMes"
	GetGroupMesType         = "GetGroupMes"
	LogoutMesType           = "LogoutMes"
	LogoutResMesType        = "LogoutResMes"
)

//定义用户状态常量
const (
	UserOnline = iota //自增
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//定义具体消息
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code    int    `json:"code"` //返回码 500表示未注册，200表示登录成功
	UsersId []int  //保存用户id的切片
	Error   string `json:"error"` //返回的错误信息
}

type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  //400表示已占用，200表示成功
	Error string `json:"error"` //返回错误信息
}

//配合服务器端推送用户状态的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

//发送的聊天消息结构体
type SmsMes struct {
	Content  string    `json:"content"` //内容
	TargetId int       //私聊用户的id
	MesDate  time.Time //群消息发送时间
	User               //继承的User结构体
}

//群消息查询结构体
type GroupResMes struct {
	GroupId int
	Content string
}

type LogoutMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type LogoutResMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
