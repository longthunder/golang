// param
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	bind(r, 200, 207)
	bind(r, 300, 307)
	bind(r, 400, 426)
	bind(r, 500, 510)
	defer catchErr()
	r.Run(":80")

}

func bind(r *gin.Engine, i, j int) {
	for index := i; index <= j; index++ {
		ii := index
		r.GET(fmt.Sprintf("/%d", ii), func(c *gin.Context) {
			log.Println(ii)
			c.String(ii, "hi", nil)
		})
	}
}

func catchErr() {
	if err := recover(); err != nil {
		log.Println("Err:", err)
	}
}
