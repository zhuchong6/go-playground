package service

import (
	"github.com/gin-gonic/gin"
	"minichat/model"
)

// GetUserServiceList
// @Tags 用户模块
// @Summary 用户列表
// @Success 200 {string} json{"code","userList"}
// @Router /user/list [get]
func GetUserServiceList(c *gin.Context) {
	list := model.FindUserList()
	c.JSON(200, gin.H{
		"userList": list,
	})
}

// CreateUser
// @Tags 用户模块
// @Summary 新增用户
// @param name query string false "用户名"
// @param pwd query string false "密码"
// @param repwd query string false "重复密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/create [get]
func CreateUser(c *gin.Context) {
	user := model.UserBasic{}
	user.Name = c.Query("name")
	pwd := c.Query("pwd")
	repwd := c.Query("repwd")
	if pwd != repwd {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.Password = pwd
	model.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功",
	})
}
