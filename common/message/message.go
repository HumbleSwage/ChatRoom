package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes" //主动推送
	SmsMesType              = "SmsMes"
)



//这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus 
)



type Message struct{
	Type string `json:"type"`//message type 
	Data string `json:"data"`//message data 
}
// 登陆的消息协议
type LoginMes struct {
	UserId int `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResMes struct {
	Code int `json:"code"`//500 means not registe,200 means login success
	UserIds []int							//保存用户id的切片
	Error string `json:"error"`//return error information
}
// 注册的消息协议
type RegisterMes struct {
	User User `json:"user"`//类型就是结构体
}

type RegisterResMes struct {
	Code int `json:"code"` //返回状态码 400表示该用户已经占用 200注册成功
	Error string `json:"error"`
}

//为了配合服务器端推送用户状态的变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:userId`
	Status int `json:status`
}

//增加一个SmsMes //发送
type SmsMes struct {
	Content string `json:"content"`
	User //匿名结构体，继承
}
//SmsReMes 








