package main

import (
	"fmt"

	"github.com/mingkid/jtt-gateway/conf"
	_ "github.com/mingkid/jtt-gateway/db"
	"github.com/mingkid/jtt-gateway/server/jtt"
	"github.com/mingkid/jtt-gateway/server/web"
)

func main() {

	go func() {
		web.Serve(fmt.Sprintf(":%d", conf.Web.Port))
	}()
	jtt.Serve(conf.JTT.Port)
}
