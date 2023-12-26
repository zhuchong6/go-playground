package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(r *gin.Engine) {
	r.GET("/redirect", func(ctx *gin.Context) {
		// 301 redirect to /
		ctx.Redirect(http.StatusMovedPermanently, "/")
	})
}
