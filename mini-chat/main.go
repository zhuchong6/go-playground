package main

import (
	"minichat/router"
	"minichat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()

	r := router.Router()

	_ = r.Run(":8081")
}
