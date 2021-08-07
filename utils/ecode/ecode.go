package ecode

import (
	"github.com/pkg/errors"
	"strconv"
)

var (
	_messages = map[int]string{}
	_codes    = map[int]struct{}{}
)

func Register(cm map[int]string) {
	_messages = cm
}

func New(e int) Code {
	if e <= 0 {
		panic("ecode 必须是大于0的数字")
	}
	return add(e)
}

func add(e int) Code {
	if _, ok := _codes[e]; ok {
		panic("ecode已经存在")
	}
	_codes[e] = struct{}{}
	return Int(e)
}

type Codes interface {
	Error() string
	Code() int
	Message() string
	Details() []interface{}
	Equal(error) bool
}

type Code int

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e Code) Code() int { return int(e) }

func (e Code) Message() string {
	if msg, ok := _messages[e.Code()]; ok {
		return msg
	}
	return e.Error()
}

func (e Code) Details() []interface{} { return nil }

func (e Code) Equal(err error) bool { return EqualError(e, err) }

func Int(i int) Code { return Code(i) }

func String(e string) Code {
	if e == "" {
		return OK
	}
	i, err := strconv.Atoi(e)
	if err != nil {
		return ServerErr
	}
	return Code(i)
}

func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
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
