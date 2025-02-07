package dao

import (
	"gocument/app/api/global"
	"gocument/app/api/internal/model"
)

func FindUser(user *model.User) (bool, error) {
	result := global.MysqlDB.Model(model.User{}).Where("username = ? OR password = ?", user.Username, user.Password).First(user)
	if result.RowsAffected == 0 {
		return false, nil
	} else {
		return true, result.Error
	}
}

func InsertUser(user *model.User) error {
	result := global.MysqlDB.Model(model.User{}).Create(user)
	return result.Error
}

func UpdateUser(user *model.User) error {
	result := global.MysqlDB.Model(model.User{}).Where("id = ? OR username = ?", user.ID, user.Username).Updates(user)
	return result.Error
}
