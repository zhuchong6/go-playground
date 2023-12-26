package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BindData struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

type NestBindData struct {
	Bind BindData
	Sex  string `form:"sex"`
}


func BindFormData(r *gin.Engine) {
	r.GET("/bindFormDataToStruct", func(ctx *gin.Context) {
		var data BindData
		if err := ctx.Bind(&data); err != nil {
			fmt.Println("bind err")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"name": data.Name,
			"age":  data.Age,
		})
	})
}

func BindFormDataToNest(r *gin.Engine) {
	r.GET("/bindFormDataToNestStruct", func(ctx *gin.Context) {
		var data NestBindData
		if err := ctx.Bind(&data); err != nil {
			fmt.Println("bind err")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"info": data.Bind,
			"sex":  data.Sex,
		})
	})
}

func NoBindFormData(r *gin.Engine) {
	r.GET("/noBindFormDataToStruct", func(ctx *gin.Context) {
		name := ctx.Query("name")
		age := ctx.Query("age")

		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
}

