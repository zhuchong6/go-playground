package service

import "github.com/gin-gonic/gin"

// IndexService
// @Tags index
// @Summary indexPage
// @Title indexPage
// @Version 1.0
// @BasePath: /
// @Produce json
// @Success 200 {string} json{"code":"200", "msg":"success", "data":"idnex" }
// @Router /index [get]
func IndexService(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "index",
	})
}
