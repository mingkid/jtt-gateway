package main

import (
	_ "github.com/mingkid/jtt808-gateway/db"
)

func main() {
	service := NewServer("", 7676)
	service.Serve()
}
