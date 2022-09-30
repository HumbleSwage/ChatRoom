package utils

import(
	"fmt"
	"net"
	"go_code/ChatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
)

type Transfer struct {
	//这个地方字段的确定取决于所关联的方法需要哪些变量
	Conn net.Conn
	Buf [8064]byte //transfor's buffer
}

func (t *Transfer)ReadPkg()(mes message.Message,err error){
	//use loop to get the information
	n, err := t.Conn.Read(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("t.Conn.Read err=",err)
		return
	}
	//covert t.Buf[:4] to uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[:4])
	//reading len is pkglen
	n, err = t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg heaher error")
		return
	}
	//UnMarshal t.Buf to message.Message
	err = json.Unmarshal(t.Buf[:pkgLen],&mes)
	if err != nil {
		err = errors.New("read pkg body error")
		return
	}
	return
}


func (t *Transfer)WritePkg(data []byte)(err error){
	//send the len of message to client 
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(t.Buf[0:4],pkgLen) // int ----> []byte
	n,err := t.Conn.Write(t.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("writePkg conn.Write(len) fail",err)
		return
	}
	//send the data itself to client
	n,err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("writePkg conn.Write(len) fail",err)
		return
	}
	return
}