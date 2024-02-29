package main

import (
	"mini-admin/config"
	"mini-admin/router"
)

// @title mini-admin
// @version 1.0
// @description admin project
// @contact.name zhuchong6.github.io
func main() {
	config.Init()

	r := router.Router()

	r.Run(":8081")
}
