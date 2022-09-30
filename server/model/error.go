package model


import(
	"errors"
)

//根据业务需要，来自定义一些错误
var (
	
	ERROR_USER_NOEXISTS = errors.New("user not exists....")
	ERROR_USER_EXISTS = errors.New("user have exists....")
	ERROR_USER_PWD = errors.New("your password maybe not right....")
)