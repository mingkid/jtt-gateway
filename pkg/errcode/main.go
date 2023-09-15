package errcode

import "fmt"

type Error struct {
	// 错误码
	code int
	// 错误消息
	msg string
	// 详细信息
	details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return Error{code: code, msg: msg}
}

func (e Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Msg() string {
	return e.msg
}

func (e Error) Details() []string {
	return e.details
}

func (e Error) Msgf(args ...any) Error {
	e.msg = fmt.Sprintf(e.msg, args...)
	return e
}

func (e Error) SetMsg(msg string) Error {
	e.msg = msg
	return e
}

// 存放公共的异常码

var (
	ParamsError       = NewError(1000000, "参数错误")
	UnauthorizedError = NewError(1000001, "未授权")
	NotFoundError     = NewError(1000002, "未找到")
)
