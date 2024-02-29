package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id int `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Age int `json:"age"`
	Sex string `json:"sex"`
	Address string `json:"address"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}