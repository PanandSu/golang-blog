package service

import "golang-blog/models/entity"

func GetUserByLoginName(loginName string) entity.UserEntity {
	var user entity.UserEntity
	Db.Where(&entity.UserEntity{LoginName: loginName}).First(&user)
	return user
}

func SignIn(user entity.UserEntity) uint {
	user.ID = 123
	Db.Create(&user)
	return user.ID
}

func Update(user entity.UserEntity) uint {
	Db.Model(&entity.UserEntity{}).Where(&entity.UserEntity{Id: user.ID}).Updates(user)
	return user.ID
}
