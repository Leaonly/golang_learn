package message

//定义用户的结构体
type User struct {
	//为了序列化和反序列化，需要用户信息json字符串key和结构体的字段tag对应！
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态
	Sex        string `json:"sex"`        //性别
}
