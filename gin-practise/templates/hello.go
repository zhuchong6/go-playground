package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	Name string
	Age  int
}

func HelloHtml(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")
	stu1 := &student{Name: "zzz", Age: 11}
	stu2 := &student{Name: "aaa", Age: 22}

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "hello.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
}
