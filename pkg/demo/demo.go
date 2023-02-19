package demo

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Serve() {
	port := os.Getenv("Serve_Port")
	router := gin.Default()

	router.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello my friend")
	})

	err := router.Run("0.0.0.0:" + port)
	if err != nil {
		panic(err)
	}
}
