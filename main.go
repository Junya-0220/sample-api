package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello/2", func(ctx *gin.Context) {
		ctx.String(200, "Hello World")
	})

	log.Print("my first log")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
