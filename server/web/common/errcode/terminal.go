package errcode

// 存放关于Terminal的异常码

var (
	SNCanNotNull = NewError(2000000, "序号不能为空")
)
