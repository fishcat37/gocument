package dao

import (
	"fmt"
	"gocument/app/api/global"
	"gocument/app/api/internal/model"
)

func FindUser(user *model.User) (bool, error) {
	result := global.MysqlDB.Model(model.User{}).Where("username = ? AND password = ?", user.Username, user.Password).First(user)
	if result.RowsAffected == 0 {
		return false, nil
	} else {
		return true, result.Error
	}
}

func FindUserByName(user *model.User) (bool, error) {
	result := global.MysqlDB.Model(model.User{}).Where("username = ?", user.Username).First(user)
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
	result := global.MysqlDB.Model(model.User{}).Where("username = ?", user.Username).Omit("password").Updates(user)
	//if result.RowsAffected == 0 {
	//	return fmt.Errorf("用户不存在")
	//}
	return result.Error
}

func UpdatePassword(user *model.User, newPassword string) error {
	result := global.MysqlDB.Model(model.User{}).Where("username = ? AND password = ?", user.Username, user.Password).Update("password", newPassword)
	if result.RowsAffected == 0 {
		return fmt.Errorf("密码错误")
	}
	return result.Error
}
