package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GlobalMiddleware(r *gin.Engine) {
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())
}

func RouteMiddleware(r *gin.Engine) {
	r.GET("/benchmark", gin.Logger(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]int{
			"a": 1,
			"b": 2,
		})
	})
}

func GroupMiddleware(r *gin.Engine) {
	authGroup := r.Group("/auth")
	authGroup.Use(gin.BasicAuth(gin.Accounts{
		"name": "z",
		"pwd":  "123",
	}))
	{
		authGroup.GET("/login", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "login success")
		})
		authGroup.GET("/hello", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "hello success")
		})
	}
}

func CustomerMiddleware(r *gin.Engine) {
	r.GET("/customerMiddleware", customerMiddlewareLogger(), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "customerMiddleware is ok")
	})
}

// customer a middleware that return gin.HandlerFunc type
func customerMiddlewareLogger() gin.HandlerFunc {
	customerMiddleware := func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request
		log.Println("before request")

		c.Next()

		// after request
		log.Println("after request")
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
	return customerMiddleware
}
