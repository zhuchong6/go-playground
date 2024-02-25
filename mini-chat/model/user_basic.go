package model

import (
	"fmt"
	"gorm.io/gorm"
	"minichat/utils"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     *time.Time
	HeartBeatTime *time.Time
	LoginOutTime  *time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLoginOut    bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func FindUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.Conn.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.Conn.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.Conn.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.Conn.Model(&user).Updates(&UserBasic{
		Name:     user.Name,
		Password: user.Password,
	})
}
