package main

import (
	"net/http"

	"zhuchong6.com/mini-gin/mgin"
)

func main() {
	r := mgin.New()

	r.GET("/", func(c *mgin.Context) {
		c.HTML(http.StatusOK, "<h1>hello html</h1>")
	})

	r.GET("/hello", func(c *mgin.Context) {
		c.String(http.StatusOK, "hello %s, you're at \"%s\"\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *mgin.Context) {
		c.JSON(http.StatusOK, mgin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8080")
}
