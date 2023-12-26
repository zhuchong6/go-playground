package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserGroup(r *gin.Engine) {
	defaultHandler := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"path": ctx.FullPath(),
		})
	}

	// group :v1
	v1 := r.Group("/user")
	{
		v1.GET("/u1", defaultHandler)
		v1.GET("/u2", defaultHandler)
	}

	// group :v1
	v2 := r.Group("/wallet")
	{
		v2.GET("/w1", defaultHandler)
		v2.GET("/w2", defaultHandler)
	}
}
