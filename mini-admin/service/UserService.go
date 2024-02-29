package service

import (
	"mini-admin/repository"

	"github.com/gin-gonic/gin"
)

// GetUserByName
// @Tags user
// @Summary GetUserByName
// @Title GetUserByName
// @Version 1.0
// @BasePath: /
// @Produce json
// @param name query string false "user's name"
// @Success 200 {string} json{"code":"200", "msg":"success", "data":"idnex" }
// @Router /users/getUserByName [get]
func GetUserByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	user := repository.FindUserByName(name)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": user,
	})
}
