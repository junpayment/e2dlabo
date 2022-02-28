package main

import (
	"e2dlabo/real_world/handlers"
	"e2dlabo/real_world/infrastructures"
	"e2dlabo/real_world/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	db, err := infrastructures.NewDB(`user:password@tcp(0.0.0.0:3306)/real_world`)
	if err != nil {
		log.Fatalln(err)
	}
	h := &handlers.Yeah{
		UserService: &services.UserService{
			DB: db,
		},
	}
	r.GET("/me", h.Me)
	err = r.Run(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
