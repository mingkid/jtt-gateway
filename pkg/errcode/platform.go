package errcode

var (
	PlatformIdentityNotNull = NewError(201001, "标识不能为空")
	PlatformHostNotNull     = NewError(201002, "平台域名不能为空")
	PlatformHostFormatError = NewError(201003, "业务平台域名格式有误")
)
