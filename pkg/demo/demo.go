package demo

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Serve() {
	// 通过yaml配置容器内的一些参数,通过env传入,具体可以看deployment/demo/demo.yaml中的env内容
	port := os.Getenv("Serve_Port")
	// 创建一个简单的http服务器对外提供服务
	router := gin.Default()

	router.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello my friend")
	})

	router.Run("0.0.0.0:" + port)
}
