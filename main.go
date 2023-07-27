package main

import (
	_ "github.com/mingkid/jtt808-gateway/db"
	"github.com/mingkid/jtt808-gateway/domain"
	"github.com/mingkid/jtt808-gateway/server/jtt808"
	"github.com/mingkid/jtt808-gateway/server/web"
)

func main() {
	jt808svr := jtt808.NewServer("", 7676, domain.NewMemorySessionCache())
	go jt808svr.Serve()

	web.Serve(":8000")
}
