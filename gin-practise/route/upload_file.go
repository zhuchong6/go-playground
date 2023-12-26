package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadOneFile(r *gin.Engine) {
	r.POST("/uploadOne", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		ctx.String(http.StatusOK, "%s uploaded!", file.Filename)
	})
}

func UploadMultiFile(r *gin.Engine) {
	r.POST("/uploadMulti", func(ctx *gin.Context) {
		f, _ := ctx.MultipartForm()
		files := f.File["files"]
		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}
		ctx.String(http.StatusOK, "%d files uploaded!", len(files))
	})
}
