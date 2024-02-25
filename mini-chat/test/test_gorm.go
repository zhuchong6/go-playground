package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"minichat/model"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3308)/demo?charset=utf8mb4&parseTime=true&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&model.UserBasic{})

	userBasic := &model.UserBasic{}
	userBasic.Name = "zcc2"
	// Create
	db.Create(userBasic)

	// Read
	var user model.UserBasic
	fmt.Println(db.First(&user, 1)) // 根据整型主键查找
	// db.First(&product, "id = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)
}
