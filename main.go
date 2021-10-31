package main

import (
	"github.com/gin-gonic/gin"
)

func health(ctx *gin.Context) {
	ctx.Writer.WriteString("health")
}

func main() {
	g := gin.Default()
	g.GET("/abc/health", health)

	g.Run(":8081")
}
