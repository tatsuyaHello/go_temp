package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tatsuyaHello-template/db"
	"github.com/tatsuyaHello-template/model"
)

func GetUser(email string) (*model.User, bool) {
	var user model.User
	if err := db.DBConn.Where("email = ?", email).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return nil, false
	}
	return &user, true
}

func GetHashByEmail(email string) (*string, error) {
	var user model.User
	err := db.DBConn.Where("email = ?", email).First(&user).Error
	return user.Password, err
}

func CreateUser(requestUser *model.RequestUser) (*model.User, error) {
	var user model.User
	// テーブル名を指定しないと、requestUsersというテーブルをみに行ってしまう
	err := db.DBConn.Table("users").Create(&requestUser).Error
	if err != nil {
		return nil, err
	}
	err = db.DBConn.Where("email = ? AND password = ?", requestUser.Email, requestUser.Password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
