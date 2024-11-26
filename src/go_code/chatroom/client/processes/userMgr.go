package processes

import (
	"fmt"
	"go_code/chatroom/client/model"
	"go_code/chatroom/common/message"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //在用户登陆成功后，完成初始化

// 客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历onlineUsers
	fmt.Println("当前在线用户列表:")
	for id := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// 处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//优化，判断map有没id再update
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}

// 处理离线服务端返回的消息
func updateUserStatusOffline(logoutResMes *message.LogoutResMes) {

	//优化，判断map有没id再update
	user, ok := onlineUsers[logoutResMes.UserId]
	if !ok {
		delete(onlineUsers, logoutResMes.UserId)
	}
	user.UserStatus = logoutResMes.Status
	delete(onlineUsers, logoutResMes.UserId)
	fmt.Println(logoutResMes.UserId, "已离线")

}
