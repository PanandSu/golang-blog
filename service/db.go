package service

import (
	"golang-blog/models/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/blog"), &gorm.Config{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&entity.UserEntity{})
	if err != nil {
		panic(err)
	}
}
