package model

import (
	"net"
	"go_code/ChatRoom/common/message"
)

//因为在客户端，很多地方都会用到CurUser，我们将其作为一个全局的变量
type CurUser struct {
	Conn net.Conn
	message.User
} 