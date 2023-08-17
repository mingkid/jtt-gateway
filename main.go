package main

import (
	"fmt"

	"github.com/mingkid/jtt808-gateway/domain/conn"

	"github.com/mingkid/jtt808-gateway/conf"
	_ "github.com/mingkid/jtt808-gateway/db"
	"github.com/mingkid/jtt808-gateway/server/jtt808"
	"github.com/mingkid/jtt808-gateway/server/web"
)

func main() {
	jt808svr := jtt808.New("", conf.JTT808.Port, conn.DefaultConnPool())
	go jt808svr.Serve()

	web.Serve(fmt.Sprintf(":%d", conf.Web.Port))
}
