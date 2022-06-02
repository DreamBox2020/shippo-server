package ecode

import (
	"shippo-server/utils/email"
	"strconv"
)

var (
	_codes = map[int]struct{}{}
)

func New(code int, message string) Code {
	if _, ok := _codes[code]; ok {
		panic("code已经存在")
	}
	_codes[code] = struct{}{}
	return Code{code, message}
}

type Codes interface {
	Error() string
	Code() int
	Message() string
	Equal(error) bool
}

type Code struct {
	code    int
	message string
}

func (e Code) Error() string {
	return "Error " + strconv.FormatInt(int64(e.code), 10) + ": " + e.message
}

func (e Code) Code() int {
	return e.code
}

func (e Code) Message() string {
	return e.message
}

func (e Code) Equal(err error) bool { return EqualError(e, err) }

func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := e.(Codes)
	if ok {
		return ec
	}
	// 未知的错误，发送邮件给管理员。
	email.SendWarningEmail(e.Error())
	return ServerErr
}

func Equal(a, b Codes) bool {
	if a == nil {
		a = OK
	}
	if b == nil {
		b = OK
	}
	return a.Code() == b.Code()
}

func EqualError(code Codes, err error) bool {
	return Cause(err).Code() == code.Code()
}
