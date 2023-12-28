package main

import (
	"github.com/gin-gonic/gin"
	"zhuchong6.com/gin-prictise/middleware"
	"zhuchong6.com/gin-prictise/route"
	"zhuchong6.com/gin-prictise/templates"
)

func main() {
	r := gin.Default()

	// curl "http://localhost:8080/"
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello gin")
	})

	// curl "http://localhost:8080/ascjsonDemo"
	// return ascjson like {"lang":"go \u003e\u003e\u003e","tag":"\u003cbr\u003e"}
	route.AscjsonDemo(r)

	// curl "http://localhost:8080/bindFormDataToStruct?name=zzz&age=222"
	// will return {"age":222,"name":"zzz"}
	route.BindFormData(r)
	// curl "http://localhost:8080/bindFormDataToNestStruct?name=zzz&age=222&sex=man"
	// will return {"info":{"Name":"zzz","Age":222},"sex":"man"}
	route.BindFormDataToNest(r)
	// curl "http://localhost:8080/bindFormDataToStruct?name=aaa&age=222"
	// will return {"age":222,"name":"aaa"}
	route.NoBindFormData(r)
	//
	route.BindDataFromUri(r)

	// curl --request POST --url http://localhost:8080/uploadOne --header 'content-type: multipart/form-data' --form file=@/Users/knight/Desktop/ic.png
	route.UploadOneFile(r)
	// curl --request POST --url http://localhost:8080/uploadMulti --header 'content-type: multipart/form-data' --form 'files=@/Users/knight/Desktop/ic.png,/Users/knight/Desktop/截屏2023-11-14 16.41.35.png'
	route.UploadMultiFile(r)

	//
	route.UserGroup(r)

	route.Redirect(r)

	// curl --request GET --url http://localhost:8080/hello
	templates.HelloHtml(r)

	middleware.GlobalMiddleware(r)
	middleware.RouteMiddleware(r)
	middleware.GroupMiddleware(r)
	middleware.CustomerMiddleware(r)

	r.Run()
}
