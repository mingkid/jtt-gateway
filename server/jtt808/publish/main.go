package publish

type Publisher interface {
	Locate(opt LocationOpt) error
}

// New 创建订阅服务
func New(host, locateURL string) Publisher {
	return &HTTP{
		host:        host,
		LocationAPI: locateURL,
	}
}
