package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"zhuchong6.com/mini-gin/mgin"
)

func main() {
	r := mgin.New()

	r.GET("/hello", func(c *mgin.Context) {
		c.String(http.StatusOK, "hello %s, you're at \"%s\"\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *mgin.Context) {
		// expect /hello/zz
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *mgin.Context) {
		c.JSON(http.StatusOK, mgin.H{"filepath": c.Param("filepath")})
	})

	r.POST("/login", func(c *mgin.Context) {
		c.JSON(http.StatusOK, mgin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *mgin.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *mgin.Context) {
			// expect /hello/zzz
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHtmlGlob("template/**")
	r.Static("/assets", "./static")

	r.GET("/", func(c *mgin.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.Run(":8080")
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func onlyForV2() mgin.HandleFunc {
	return func(c *mgin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		// c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
