package processes

import(
	"fmt"
	"go_code/ChatRoom/common/message"
	"go_code/ChatRoom/client/model"

)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)

var CurUser model.CurUser //我们在客户登陆成功后完成对CurUser初始化

//在客户端显示当前在线的用户
func outputOnlineUser(){
	fmt.Println("当前在线用户列表如下：")
	//将onlineUser遍历
	for id , _ := range onlineUsers {
		fmt.Println("用户id：",id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes){
	//适当的优化
	user , ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User {
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
}

