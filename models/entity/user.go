package entity

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Id        uint   `gorm:"primary_key"`
	Age       int    `json:"age"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	LoginName string `json:"login_name"`
}

func (UserEntity) TableName() string {
	return "User"
}
