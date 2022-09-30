package main

import(
	"fmt"
	"net"
	"time"
	"go_code/ChatRoom/server/model"
)

func init(){
	//注意初始化的相关顺序
	//服务器端一开始就是初始化连接池
	initPool("localhost:6379",16,0,300 * time.Second)
	//获取userDao，pool在redis.go中设置为全局变量
	model.MyUserDao = model.NewUserDao(pool)

}


func process(conn net.Conn){
	//its necessary to close conn at here
	defer conn.Close()
	processor := &Processor{
		Conn : conn,
	}
	err := processor.Process()
	if err != nil {
		fmt.Println("some err happpen between client and server,err=",err)
		return
	}
	
}

func main(){
	fmt.Println("server run on localhost:8889")
	listen,err := net.Listen("tcp","localhost:8889")
	//its necessary to close listen at here
	defer listen.Close()
	if err != nil {
		fmt.Println("net listen err=",err)
		return
	}
	//if listen success,wait for client connect
	for {
		fmt.Println("wait for client connect.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err=",err)
			return
		}
		go process(conn)
	}
}