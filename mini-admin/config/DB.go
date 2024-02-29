package config

import (
	"fmt"

	log2 "log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func Init() {
	// config log to print log
	log := logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阀值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //颜色
		},
	)

	dsn := "root:root@tcp(127.0.0.1:3308)/demo?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Conn = db
}
