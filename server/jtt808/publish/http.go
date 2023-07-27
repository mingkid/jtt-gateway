package publish

import (
	"io"
	"net/http"
)

// HTTP 订阅服务
type HTTP struct {
	host        string
	LocationAPI string
}

// Locate 推送定位数据
func (h *HTTP) Locate(opt LocationOpt) error {
	buffer, err := opt.Buffer()
	resp, err := http.Post(h.LocationAPI, "application/json;charset=UTF-8", buffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	return err
}

// New 创建订阅服务
func New(host string) *HTTP {
	return &HTTP{
		host: host,
	}
}
