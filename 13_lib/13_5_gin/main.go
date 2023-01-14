package main

import (
	"math/rand"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(interface{}(err))
	}

	// 通过Use()来加入middleware
	r.Use(func(context *gin.Context) {
		startTime := time.Now()
		// log path, latency, and response code
		logger.Info("[Zap Log] Incoming request",
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", context.Writer.Status()),
			zap.Duration("latency", time.Now().Sub(startTime)),
		)
		context.Next()
	}, func(context *gin.Context) {
		context.Set("requestId", rand.Int())
		context.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("This is a Home Page"))
		c.Writer.Flush()
	})

	r.GET("/hello", func(c *gin.Context) {
		c.Writer.Write([]byte("Hello Gin"))
		c.Writer.Flush()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/requestid", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exist := c.Get("requestId"); exist {
			h["requestId"] = rid
		}
		c.JSON(http.StatusOK, h)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
