package model

import (
	"go_code/chatroom/common/message"
	"net"
)

// 客户端很多地方要使用，维护一个连接使用，要做全局
type CurUser struct {
	Conn net.Conn
	message.User
}
