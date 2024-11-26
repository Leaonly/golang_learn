package processes

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"os"
	"time"
)

var content string

// 处理返回的NotifyUserStatusMes
func outputMes(mes *message.Message) {

	//1.反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("", err.Error())
		return
	}
	if mes.Type == message.SendPrivateMesType {
		//显示信息
		info := fmt.Sprintf("用户id:\t%d私聊说:\t%s", smsMes.UserId, smsMes.Content)
		println(info)
		fmt.Println()
	} else if mes.Type == message.SmsMesType {
		//显示信息
		info := fmt.Sprintf("用户id:\t%d说:\t%s", smsMes.UserId, smsMes.Content)
		println(info)
		fmt.Println()
	}

}

// 处理返回的NotifyUserStatusMes
func outputGroupMes(mes *message.Message) {

	//1.反序列化
	var smsMes map[string]string
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("", err.Error())
		return
	}
	for key, value := range smsMes {
		//反序列化消息内容
		var msgContent message.SmsMes
		err := json.Unmarshal([]byte(value), &msgContent)
		if err != nil {
			fmt.Println("解析内容失败：", err.Error())
			continue
		}
		//解析消息发送时间
		// 定义时间布局
		layout := "2006-01-02T15:04:05Z07:00"
		mesDT, err := time.Parse(layout, key)
		if err != nil {
			fmt.Println("解析时间失败：", err.Error())
			continue
		}
		mesDTf := mesDT.Format("2006-01-02 15:04:05")
		fmt.Println(mesDTf)
		fmt.Printf("用户id: %d, 说: %s\n", msgContent.UserId, msgContent.Content)
	}
}

func inputGroupMes() {
	//SmsProcess实例放外面方便多次调用
	smsProcess := &SmsProcess{}

	// 创建一个新的Scanner对象，从标准输入读取
	scanner := bufio.NewScanner(os.Stdin)

	// 提示用户输入
	fmt.Print("请输入内容: ")

	// 读取一行输入
	if scanner.Scan() {
		content = scanner.Text() // 获取输入的文本
		if content == "" {
			fmt.Println("输入内容为空，请重新输入。")
			return
		}
		//群发消息发送给服务端的方法
		smsProcess.SendGroupMes(content)
	} else {
		// 处理读取错误
		if err := scanner.Err(); err != nil {
			fmt.Println("读取输入时发生错误:", err)
		} else {
			fmt.Println("没有读取到任何内容")
		}
	}
}

func inputPrivateChatMes() {
	//SmsProcess实例放外面方便多次调用
	smsProcess := &SmsProcess{}

	// 创建一个新的Scanner对象，从标准输入读取
	scanner := bufio.NewScanner(os.Stdin)

	// 提示用户输入
	var id int
	fmt.Println("请输入用户id:")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入内容: ")

	// 读取一行输入
	if scanner.Scan() {
		content = scanner.Text() // 获取输入的文本
		if content == "" {
			fmt.Println("输入内容为空，请重新输入。")
			return
		}
		//群发消息发送给服务端的方法
		smsProcess.SendPrivateMes(content, id)
	} else {
		// 处理读取错误
		if err := scanner.Err(); err != nil {
			fmt.Println("读取输入时发生错误:", err)
		} else {
			fmt.Println("没有读取到任何内容")
		}
	}
}
