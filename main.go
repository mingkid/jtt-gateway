package main

import (
	"fmt"

	"github.com/mingkid/jtt-gateway/http"
	"github.com/mingkid/jtt-gateway/jtt"

	"github.com/mingkid/jtt-gateway/conf"
	_ "github.com/mingkid/jtt-gateway/db"
)

func main() {

	go func() {
		http.Serve(fmt.Sprintf(":%d", conf.Web.Port))
	}()
	jtt.Serve(conf.JTT.Port)
}
