package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	log2 "log"
	"os"
	"time"
)

var Conn *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("config app: ", viper.Get("app"))
	fmt.Println("config mysql: ", viper.Get("mysql"))
}

func InitMysql() {
	// 自定义日志打印sql语句
	log := logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阀值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //颜色
		},
	)
	dsn := viper.GetString("mysql.dns")
	fmt.Println("---", dsn)
	Conn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: log})
	//user := model.UserBasic{}
	//Conn.Find(&user)
	//fmt.Println(user)
}
