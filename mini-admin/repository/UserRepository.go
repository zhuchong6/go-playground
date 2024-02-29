package repository

import (
	"mini-admin/config"
	"mini-admin/model"
)

func FindUserByName(name string) model.User {
	user := model.User{
		Name: name,
	}
	config.Conn.First(&user)
	return user
}
