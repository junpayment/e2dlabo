package main

import (
	"e2dlabo/real_world/wire/web/wire"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	h, err := wire.Initialize(`user:password@tcp(0.0.0.0:3306)/real_world`)
	if err != nil {
		log.Fatalln(err)
	}
	r.GET(`/me`, h.Me)
	err = r.Run(":3000")
}
