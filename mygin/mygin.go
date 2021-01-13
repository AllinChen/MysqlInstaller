package mygin

import (
	"time"

	"github.com/gin-gonic/gin"
)

//StartGin 展开Gin服务
func StartGin() {
	r := gin.Default()

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
		time.Sleep(5 * time.Second)
		c.String(200, "asdasdas")
	})
	r.Run(":8080")
}
