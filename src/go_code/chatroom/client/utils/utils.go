package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

// 将方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这是传输时使用的缓存
}

func (wow *Transfer) ReadPkg() (mes message.Message, err error) {
	// 将读取数据包封装称一个函数readPkg(),返回message, err
	// 读取对端发送的消息长度，然后再读具体消息时就可以按照长度读取
	_, err = wow.Conn.Read(wow.Buf[:4])
	if err != nil {
		err = errors.New("read pkg header error")
		return
	}
	//根据buf[:4]转成unit32类型
	pkgLen := binary.BigEndian.Uint32(wow.Buf[0:4])

	//根据pkgLen消息长度 读取conn消息内容，即从连接中读取数据，并将读取的数据存储到切片 b 中。返回值 n 表示实际读取的字节数
	n, err := wow.Conn.Read(wow.Buf[:pkgLen])
	//验证Read返回的字节长度是否嗯pkgLen长度一样及是否有错误
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	//把pkgLen 反序列化成 -> message.Message
	err = json.Unmarshal(wow.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	return
}

func (wow *Transfer) WritePkg(data []byte) (err error) {

	//发送一个长度给对方验证，为了防止网络传输中的丢包，如果丢包则弃包重发
	pkgLen := uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(wow.Buf[0:4], pkgLen)
	//发送长度
	n, err := wow.Conn.Write(wow.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err= ", err)
		return
	}

	//发送data本身
	n, err = wow.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(buf) err= ", err)
		return
	}
	return
}
