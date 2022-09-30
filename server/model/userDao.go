package model

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"go_code/ChatRoom/common/message"
	
)

//在服务器启动后，就启动一个userDao实例
//把它做成全局的变量，在需要和redis交换的时候，直接拿取即可
var (
	MyUserDao *UserDao
)

//使用一个工厂模式创建一个UserDao的实例
func NewUserDao(pool *redis.Pool)(userDao *UserDao){
	userDao = &UserDao{
		pool : pool,
	}
	return
}

type UserDao struct {
	pool *redis.Pool
}

// 完成用户的
//1、根据用户的id返回一个User的实例+err
func (ud *UserDao)getUserById(conn redis.Conn,id int)(user User, err error){
	fmt.Println(id)
	res , err := redis.String(conn.Do("Hget","users",id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOEXISTS
		}
		return
	}
	//2、将res反序列化，将res反序列化为users实例
	err = json.Unmarshal([]byte(res),&user)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	return
}

//2、完成登陆校验
//1.userDao 完成对用户的验证；
//2.如果用户的username和pwd都正确，则返回一个user；
//3.如果用户username和pwd有错误，则返回对应的错误信息；
func (ud *UserDao)Login(UserId int, Password string)(user User, err error){
	//先从userDao的连接池取出一根连接
	conn := ud.pool.Get()
	defer conn.Close()
	user ,err = ud.getUserById(conn,UserId)
	if err != nil {
		return
	}
	//证明了用户获取到了，接下来进行密码和用户名的比对
	if user.Password != Password {
		err = ERROR_USER_PWD
		return
	}
	return

}

//3、完成用户注册
func (ud *UserDao)Register(user *message.User)(err error){
	// 先从UserDao的连接池里面中取出一根连接
	conn := ud.pool.Get()
	defer conn.Close()
	_,err = ud.getUserById(conn,user.UserId)
	if err == nil {//如果查询不报错，说明此时redis数据库中已经有该变量了
		err = ERROR_USER_EXISTS
		return
	}
	//说明该用户在redis还没有有该用户，可以入库
	data,err := json.Marshal(user)
	if err != nil {
		return
	}
	//入库
	_ , err = conn.Do("Hset","users",user.UserId,string(data))
	if err != nil {
		fmt.Println("保存用户出错！err=",err)
		return
	}
	return
}












