package message

//定义一个用户体结构
type User struct {
	UserId int `json:userid`
	Password string `json:password`
	UserName string `json:username`
	UserStatus int `json:userStatus`//用户的状态
}


