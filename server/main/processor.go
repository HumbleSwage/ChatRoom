package main


import(
	"fmt"
	"net"
	"go_code/ChatRoom/common/message"
	"go_code/ChatRoom/server/processes"
	"go_code/ChatRoom/common/utils"
	"io"
)

//创建一个Processor结构体来绑定下来的方法
type Processor struct{
	Conn net.Conn
}

func (p *Processor)Process()(err error){
	//process data to message
	tf := &utils.Transfer {
		Conn : p.Conn,
	}
	for {
		mes ,err :=  tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client exit, server exit at same time")
				return err
			} else {
				fmt.Println("readPkg err=",err)
				return err
			}
		}
		fmt.Println("mes=",mes)
		p.serverProcess(&mes)
	}
	return
}

//use different func to handle different message type
func (p *Processor)serverProcess(mes *message.Message)(err error){
	//看看是否能够接收到客户端发送的群发的消息
	fmt.Println("mes=",mes)
	switch mes.Type {
	case message.LoginMesType :
		userProcess := &processes.UserProcess{
			Conn : p.Conn,
		}
		userProcess.ServerProcessLogin(mes)
	case message.RegisterMesType :
		userProcess := &processes.UserProcess{
			Conn : p.Conn,
		}
		userProcess.ServerProcessRegister(mes)
		//处理注册逻辑
	case message.SmsMesType:
		smsProcess := &processes.SmsProcess {}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Printf("there is no %s",mes.Type)
	}
	return
}











