package processes

import (
	"fmt"
)

//UserMgr在服务器端有且仅有一个
//因为很多的地方，都会使用到，因此我们将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr {
		onlineUsers : make(map[int]*UserProcess,1024),
	} 
}

//完成对onlineUser添加
func (um *UserMgr)AddOlineUser(up *UserProcess){
	um.onlineUsers[up.UserId] = up
}

//删除
func (um *UserMgr)DelOnlineUser(up *UserProcess){
	delete(um.onlineUsers,up.UserId)
}

//返回当前所有在线的用户
func (um *UserMgr)getAllUsers() map[int]*UserProcess {
	return um.onlineUsers
}

//根据id 返回对应的值
func (um *UserMgr)GetOnlineUserById(UserId int)(up *UserProcess,err error){
	//从map中取出一个值，带检测的方式
	up , ok := um.onlineUsers[UserId]
	if !ok { // 说明，你要查找的这个用户，当前不在线
		err = fmt.Errorf("用户%d不存在",UserId)
		return 
	}
	return 
}








