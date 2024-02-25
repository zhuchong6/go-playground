package main

import (
	"minichat/router"
	"minichat/utils"
)

// @title mini-cha
// @version 1.0
// @description this project is used for learning
func main() {
	utils.InitConfig()
	utils.InitMysql()

	r := router.Router()

	_ = r.Run(":8081")
}
