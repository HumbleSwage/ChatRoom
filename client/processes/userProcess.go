package processes

import (
	"fmt"
	"encoding/json"
	"encoding/binary"
	"net"
	"go_code/ChatRoom/common/message"
	"go_code/ChatRoom/common/utils"
	_"time"
	"os"
)

type UserProcess struct {
	//暂时不需要字段
}

func (up *UserProcess)Login(UserId int, Password string)(err error){
	//use dial to connect server
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return
	}
	//delaye closing the connect
	defer conn.Close()
	//send message to server after serialize data
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginmes message.LoginMes
	loginmes.UserId = UserId
	loginmes.Password = Password
	data,err := json.Marshal(loginmes)
	if err != nil {
		fmt.Println("loginmes json marshal err=",err)
		return
	}
	mes.Data = string(data) //[]byte ----> string
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("Message json marshal err=",err)
		return
	}
	//warning:must convert int to []byte
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	//send the len of data to server
	n,err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) fail",err)
		return
	}
	//send the data to server
	_,err = conn.Write(data)
	fmt.Println(string(data))
	if n != 4 || err != nil {
		fmt.Println("conn.Write(data) fail",err)
		return
	}
	//recieve the data from server
	tf := &utils.Transfer{
		Conn : conn,
	}
	mes, err = tf.ReadPkg() 
	if err != nil {
		fmt.Println("ReadPkg() err=",err)
		return
	}
	//convert data to LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = UserId
		CurUser.UserStatus = message.UserOnline
		//显示当前在线用户列表，遍历loginResMes.UserId
		fmt.Println("当前在线用户列表如下：")
		for _,v := range loginResMes.UserIds{
			if v == UserId {
				continue
			}
			fmt.Println("用户id:",v)
			//完成客户端的onlineUsers的初始化
			user := &message.User {
				UserId : v,
				UserStatus : message.UserOnline,
			}
			onlineUsers[v] = user
		}


		//keep connect with server to get msg from server on time
		//and show this msg on cmd
		go serverProcessMes(conn)
		//show login success page
		ShowMeau(UserId)
	} else {
		fmt.Println(loginResMes)
	}
	return 
}


func (up *UserProcess)Register(UserId int, UserName string,Password string)(err error){
	//use dial to connect server
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return
	}
	//delaye closing the connect
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User = message.User {
		UserId : UserId,
		UserName : UserName,
		Password : Password,
	}
	data,err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("loginmes json marshal err=",err)
		return
	}
	mes.Data = string(data) //[]byte ----> string
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("Message json marshal err=",err)
		return
	}
	tf := &utils.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		if err != nil {
		fmt.Println("Register WritePkg() err=",err)
		return
		}
	}
	mes, err = tf.ReadPkg()  // mes就是RegisterResMes
	if err != nil {
		fmt.Println("Register ReadPkg() err=",err)
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("Register success,please exit then login")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return 


}