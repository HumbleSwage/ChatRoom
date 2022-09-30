package processes

import (
	"fmt"
	"net"
	"go_code/ChatRoom/common/message"
	"go_code/ChatRoom/common/utils"
	"encoding/json"
	"go_code/ChatRoom/server/model"
)

type UserProcess struct {
	//字段
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户
	UserId int
}
//这里我们编写通知所有在线用户的方法
//这个UserID要通知其他用户，我上线
func (up *UserProcess)NotifyOthersOnlineUser(UserId int){
	//遍历 onlineUsers，然后一个一个的发送NotifyUserStatusMes
	for id,up := range userMgr.onlineUsers{
		if id == UserId {
			continue
		}
		//开始通知【单独写一个方法】
		up.notifyUserStatusMes(UserId)
	}
}

func (up *UserProcess)notifyUserStatusMes(UserId int){
	//组装我们的NotifyUserStatus
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = UserId
	notifyUserStatusMes.Status = message.UserOnline
	//将notifyUserStatusMes序列化
	data ,err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notifyUserStatusMes json.Marshal err=",err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//对mes再次序化，准备发送
	data,_  = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	tf := utils.Transfer{
		Conn : up.Conn, 
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=",err)
		return
	}
}


//serverProcessLogin：Process Login server
func (up *UserProcess)ServerProcessLogin(mes *message.Message)(err error){
	//get data from message and Unmarshal to LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("login message Unmarshal fail ,err=",err)
		return
	}
	//declare a resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//declare a LoginResMes
	var loginResMes message.LoginResMes
	// if loginMes.Username == "dz" && loginMes.Password == "123"{
	// 	//合法
	// 	loginResMes.Code = 200
	// } else{
	// 	loginResMes.Code = 500 //500 means this user is not exits
	// 	loginResMes.Error = "this user is not exit,please register firstly"
	// 	//不合法
	// }
	// 需要去redis数据库完成验证
	// 1、使用model.MyUserDao到redis中去完成验证
	user,err := model.MyUserDao.Login(loginMes.UserId,loginMes.Password)

	if err != nil {
		//先测试成功，然后再根据具体情况返回具体的错误信息
		if err == model.ERROR_USER_NOEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务内部错误。。。"
		}
	} else {
		loginResMes.Code = 200
		//将ID放进去
		up.UserId = loginMes.UserId
		//将当前在线用户的id放入LoginResMes.UserIds
		//用户登陆成功，把该登陆成功的放入userMgr中
		userMgr.AddOlineUser(up)
		//通知其他在线用户我上线了
		up.NotifyOthersOnlineUser(loginMes.UserId)
		fmt.Println(user,"登陆成功")
		//遍历userMgr.onlineUsers
		for id,_ := range userMgr.onlineUsers{
			loginResMes.UserIds = append(loginResMes.UserIds,id)
		}
	}
	data,err :=  json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}
	//tap the data to resMes
	resMes.Data  = string(data)
	//Marshal resMes,and prepare send it
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes Marshal fail,err=",err)
	}
	//wirtePkg
	//create a transfer which tap the utils func
	tf := utils.Transfer{
		Conn : up.Conn, 
	}
	err = tf.WritePkg(data)
	return
}

//serverProcessRegister
func (up *UserProcess)ServerProcessRegister(mes *message.Message)(err error){
	//get data from message and Unmarshal to LoginMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil {
		fmt.Println("mes message Unmarshal fail ,err=",err)
		return
	}
	//declare a resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//declare a LoginResMes
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS{
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误。。。"
		}
	} else {
		registerResMes.Code = 200
	}

	data,err :=  json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}
	//tap the data to resMes
	resMes.Data  = string(data)
	//Marshal resMes,and prepare send it
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes Marshal fail,err=",err)
	}
	tf := utils.Transfer{
		Conn : up.Conn, 
	}
	err = tf.WritePkg(data)
	return

}
