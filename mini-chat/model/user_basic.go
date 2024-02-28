package model

import (
	"fmt"
	"minichat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
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
		Phone:    user.Phone,
		Email:    user.Email,
	})
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.Conn.Where("name = ?", name).First(&user)
	return user
}
