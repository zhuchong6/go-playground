package route

import "github.com/gin-gonic/gin"

type BindUrl struct {
	ID   int    `uri:"id" `
	Name string `uri:"name" `
}

func BindDataFromUri(r *gin.Engine) {
	r.GET("/:name/:id", func(ctx *gin.Context) {
		var b BindUrl
		if err := ctx.ShouldBindUri(&b); err != nil {
			ctx.JSON(400, gin.H{"msg": err})
			return
		}
		ctx.JSON(200, gin.H{"name": b.Name, "uuid": b.ID})
	})
}
