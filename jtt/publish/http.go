package publish

import (
	"io"
	"net/http"
	"net/url"

	"github.com/mingkid/jtt-gateway/pkg/errcode"
)

// HTTP 订阅服务
type HTTP struct {
	host        string
	LocationAPI string
}

// Locate 推送定位数据
func (h *HTTP) Locate(opt LocationOpt) error {
	buffer, err := opt.Buffer()
	url, err := h.getAPI(h.LocationAPI)
	resp, err := http.Post(url, "application/json;charset=UTF-8", buffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	return err
}

func (h *HTTP) getAPI(api string) (string, error) {
	if api == "" {
		return "", errcode.PlatformHostNotNull
	}

	host, err := url.Parse(h.host)
	if err != nil {
		return "", errcode.PlatformHostFormatError
	}
	return host.ResolveReference(&url.URL{Path: api}).String(), nil
}
