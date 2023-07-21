package main

import (
	_ "github.com/mingkid/jtt808-gateway/db"
	"github.com/mingkid/jtt808-gateway/domain"
	"github.com/mingkid/jtt808-gateway/server/web"
)

func main() {
	service := NewServer("", 7676, domain.NewMemorySessionCache())
	go service.Serve()

	web.Serve(":8000")
}
