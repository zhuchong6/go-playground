package service

import (
	"fmt"
	"minichat/model"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserServiceList
// @Tags User
// @Summary query user‘s list
// @Success 200 {string} json{"code","userList"}
// @Router /user/list [get]
func GetUserServiceList(c *gin.Context) {
	list := model.FindUserList()
	c.JSON(200, gin.H{
		"userList": list,
	})
}

// CreateUser
// @Tags User
// @Summary add a new user
// @param name query string false "name"
// @param pwd query string false "password"
// @param repwd query string false "repeat password"
// @Success 200 {string} json{"code","message"}
// @Router /user/create [get]
func CreateUser(c *gin.Context) {
	user := model.UserBasic{}
	user.Name = c.Query("name")
	pwd := c.Query("pwd")
	repwd := c.Query("repwd")
	if pwd != repwd {
		c.JSON(-1, gin.H{
			"message": "Entered passwords differ",
		})
		return
	}
	userFromDb := model.FindUserByName(user.Name)
	if userFromDb.Name != "" {
		c.JSON(200, gin.H{
			"message": "用户已经存在",
		})
		return
	}
	user.Password = pwd
	model.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// DeleteUser
// @Tags User
// @Summary delete user by user's id
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := model.UserBasic{}

	id, _ := strconv.Atoi(c.Query("id"))

	user.ID = uint(id)
	model.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// UpdateUser
// @Tags User
// @Summary modify user's info
// @param id formData string false "id"
// @param name formData string false "name"
// @param pwd formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := model.UserBasic{}

	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("pwd")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	validateStruct, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "param err",
			"info":    validateStruct,
		})
		return
	}
	model.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "success",
	})
}
