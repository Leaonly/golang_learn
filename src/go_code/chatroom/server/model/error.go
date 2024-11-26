package model

import (
	"errors"
)

//根据业务逻辑，自定义错误

var (
	ErrUserNotExists = errors.New("用户或密码错误，请重试")
	ErrUserExists    = errors.New("用户已经存在")
	ErrUserPwd       = errors.New("密码不正确")
	ErrGroupHget     = errors.New("群消息查询错误")
)
