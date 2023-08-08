package main

import (
	"fmt"

	"github.com/mingkid/jtt808-gateway/conf"
	"github.com/mingkid/jtt808-gateway/dal"
	_ "github.com/mingkid/jtt808-gateway/db"
	"github.com/mingkid/jtt808-gateway/server/jtt808"
	"github.com/mingkid/jtt808-gateway/server/web"
)

func main() {
	jt808svr := jtt808.NewServer("", conf.JTT808.Port, dal.DefaultSessionCache)
	go jt808svr.Serve()

	web.Serve(fmt.Sprintf(":%d", conf.Web.Port))
}
