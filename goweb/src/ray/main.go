// ray project main.go
package main

import (
	"github.com/gin-gonic/gin"
)

// https://github.com/gin-gonic/gin/blob/master/README.md
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":80")

}
