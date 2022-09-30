package processes


import (
	"fmt"
	"os"
	"go_code/ChatRoom/common/utils"
	"net"
	"encoding/json"
	"go_code/ChatRoom/common/message"
)


//显示登陆成功后的页面
func ShowMeau(UserId int){
	for {
		fmt.Printf("------%d,login success--------\n",UserId)
		fmt.Println("------1、show all users--------")
		fmt.Println("------2、send message----------")
		fmt.Println("------3、show message list ----")
		fmt.Println("------4、exit system-----------")
		fmt.Println("please input number 1 ~ 4:")
		var key int
		var content string
		//因为我们总会使用到SmsProcess实例，因此将其定义switch外边
		smsProcess := &SmsProcess{}
		fmt.Scanf("%d\n",&key)
		switch key{
		case 1:
			// fmt.Println("显示在线用户列表")
			outputOnlineUser()
		case 2:
			// fmt.Println("发送消息")
			fmt.Print("请输入你想对大家说什么:")
			fmt.Scanf("%s\n",&content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("消息列表")
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("%d maybe have proplem , please choose a right one",key)
		}
	}
}
//keep connect with server
func serverProcessMes(conn net.Conn){
	tf := &utils.Transfer{
		Conn : conn,
	}
	for {
		mes , err := tf.ReadPkg()
		if err != nil {
			fmt.Println("serverProcessMes tf.ReadPkg() err=",err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 1、取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
			// 2、得把人加入到客户端维护的那个map中去
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType://有人群发消息
			outputGroupMes(&mes)
		default :
			fmt.Println("服务器端返回了一个未知的消息类型")
		}
	}
}













