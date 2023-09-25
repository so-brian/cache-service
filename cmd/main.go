package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/so-brian/cache-service/internal/pkg/memory"
)

type StringRequest struct {
	Key    string    `uri:"key" json:"key" binding:"required"`
	Value  string    `json:"value"`
	Expire time.Time `json:"expire"`
}

func main() {
	provider := memory.NewKeyValueMemoryProvider()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/string/:key", func(c *gin.Context) {
		var req StringRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		value, flag := provider.Get(req.Key)
		if !flag {
			c.JSON(404, gin.H{"msg": "not found"})
			return
		}

		c.JSON(200, gin.H{"key": req.Key, "value": value})
	})

	r.POST("/string", func(c *gin.Context) {
		var req StringRequest
		if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		provider.Set(req.Key, req.Value, &req.Expire)
		c.JSON(201, gin.H{"key": req.Key, "value": req.Value})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
