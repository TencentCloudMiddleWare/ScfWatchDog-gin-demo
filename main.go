package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type reqtype struct {
	Sid string `json:"sid"`
}
type mytoken struct {
	Name      string `json:"name"`
	Timestemp int64  `json:"timestemp"`
	Sid       string `json:"sid"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest")
		ip := c.GetHeader("X-Forwarded-For")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("%s pong ip is %s", name, ip),
		})
	})
	r.POST("/token/:name", func(c *gin.Context) {
		name := c.Param("name")
		reqp := reqtype{}
		c.ShouldBind(&reqp)
		c.JSON(200, mytoken{
			Name:      name,
			Timestemp: time.Now().Unix(),
			Sid:       reqp.Sid,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
