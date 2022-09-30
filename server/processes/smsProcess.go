package processes

import (
	"fmt"
	"net"
	"go_code/ChatRoom/common/message"
	"go_code/ChatRoom/server/utils"
	"encoding/json"
)

type SmsProcess struct {
	//暂时不需要字段
}


func (sp *SmsProcess)SendGroupMes(mes *message.Message){
	//将mes中的内容取出来
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("SendGroupMes Marshal err=",err)
	}
	data,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	//遍历服务器端的OnlineUsers map[int]*UsersProcess 
	for id , up := range userMgr.onlineUsers {
		//这里自己还需要过滤自己
		if id == smsMes.UserId {
			continue
		}
		sp.SendMesToEachOlineUsers(data,up.Conn)

	}
	//将消息转发出去
}

func (sp *SmsProcess)SendMesToEachOlineUsers(data []byte,conn net.Conn){
	//创建一个transfer实例，发送data
	tf := &utils.Transfer{
		Conn : conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败了，err=",err)
	}









}