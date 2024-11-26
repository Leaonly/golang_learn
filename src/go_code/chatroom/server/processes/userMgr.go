package processes

import (
	"fmt"
)

// 因为UserMgr 实例在服务器端有且只有一个，要定义全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 初始化UserMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 对onlineUsers添加的方法
func (wow *UserMgr) AddOnlineUser(up *UserProcess) {
	wow.onlineUsers[up.UserId] = up
}

// 删除
func (wow *UserMgr) DelOnlineUser(userId int) {
	delete(wow.onlineUsers, userId)
}

// 返回当前在线所有用户
func (wow *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return wow.onlineUsers
}

// 根据id返回对应的值
func (wow *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//从map取出对应值
	up, ok := wow.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
