// param
package main

import (
	"com/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})
	r.GET("/index.js", func(c *gin.Context) {
		c.File("index.js")
	})
	r.GET("/services", func(c *gin.Context) {
		c.JSON(200, dao.GetServices())
	})
	r.GET("/envs", func(c *gin.Context) {
		c.JSON(200, dao.GetEnvs())
	})
	r.GET("/services/:serviceId/serviceconfig", func(c *gin.Context) {
		serviceId := c.Params.ByName("serviceId")
		c.JSON(200, dao.GetServiceConfig(serviceId))
	})
	r.PUT("/services/:serviceId/envs/:envId/serviceconfig", func(c *gin.Context) {
		serviceId := c.Params.ByName("serviceId")
		envId := c.Params.ByName("envId")
		contentBytes, _ := ioutil.ReadAll(c.Request.Body)
		var content interface{}
		json.Unmarshal(contentBytes, &content)
		dao.UpdateServiceConfig(serviceId, envId, content)
		c.String(200, "")
	})

	r.GET("/collection/:collection", func(c *gin.Context) {
		collection := c.Params.ByName("collection")
		c.JSON(200, dao.GetList(collection))
	})

	r.POST("/collection/:collection", func(c *gin.Context) {
		collection := c.Params.ByName("collection")
		contentBytes, _ := ioutil.ReadAll(c.Request.Body)
		var content interface{}
		json.Unmarshal(contentBytes, &content)
		c.JSON(200, dao.New(collection, content))
	})
	r.PUT("/collection/:collection/:id", func(c *gin.Context) {
		collection := c.Params.ByName("collection")
		id := c.Params.ByName("id")
		contentBytes, _ := ioutil.ReadAll(c.Request.Body)
		var content interface{}
		json.Unmarshal(contentBytes, &content)
		c.JSON(200, dao.Update(collection, id, content))
	})

	r.DELETE("/collection/:collection/:id", func(c *gin.Context) {
		collection := c.Params.ByName("collection")
		id := c.Params.ByName("id")
		c.JSON(200, dao.Delete(collection, id))
	})

	defer catchErr()
	r.Run(":80")

}

func catchErr() {
	if err := recover(); err != nil {
		log.Println("Err:", err)
	}
}
