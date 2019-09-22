package server

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	setRouter(r)
	r.Run(":80")
}

func setRouter(r *gin.Engine) {
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello")
	})
}
