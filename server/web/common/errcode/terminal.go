package errcode

// 存放关于Terminal的异常码

var (
	TermSNNotNull = NewError(200000, "序号不能为空")
)
