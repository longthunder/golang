// ray project main.go
package main

import (
	"github.com/gin-gonic/gin"
)

// https://github.com/gin-gonic/gin/blob/master/README.md
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/index.tpl")
	r.GET("/", func(c *gin.Context) {
		obj := gin.H{"title": "Main website"}
		c.HTML(200, "index.tpl", obj)
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":80")

}
