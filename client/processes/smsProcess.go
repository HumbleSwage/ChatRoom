package processes



import (
	"fmt"
	"go_code/ChatRoom/common/message"
	"go_code/ChatRoom/common/utils"
	"encoding/json"
)

type SmsProcess struct {

}


//发送群聊消息
func (sp *SmsProcess)SendGroupMes(content string)(err error){
	//1、创建一个message.Message
	var mes message.Message
	mes.Type = message.SmsMesType	
	//2、创建一个SmsMes的实例
	var smsMes message.SmsMes
	smsMes.Content = content //内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	//3、序列化smsMes
	data,err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=",err.Error())
		return
	}
	mes.Data = string(data)
	//4、对mes再次序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err=",err.Error())
		return
	}
	//将mes消息发给服务器
	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes WritePkg err=",err.Error())
		return
	}

	return

}