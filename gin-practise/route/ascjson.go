package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AscjsonDemo(r *gin.Engine) {
	r.GET("/ascjsonDemo", func(ctx *gin.Context) {
		data := map[string]string{
			"lang": "go >>>",
			"tag":  "<br>",
		}
		ctx.AsciiJSON(http.StatusOK, data)
	})
}
