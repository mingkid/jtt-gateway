package main

import (
	"github.com/mingkid/jtt808-gateway/dal"
	_ "github.com/mingkid/jtt808-gateway/db"
	"github.com/mingkid/jtt808-gateway/server/jtt808"
	"github.com/mingkid/jtt808-gateway/server/web"
)

func main() {
	jt808svr := jtt808.NewServer("", 7676, dal.DefaultSessionCache)
	go jt808svr.Serve()

	web.Serve(":8000")
}
