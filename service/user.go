package service

import (
	"github.com/Zhangbokai614/go-template/model"
	"github.com/Zhangbokai614/go-template/utils"
)

func CreateUser(apiUser *model.APIUser) error {
	user := &model.User{
		Name:     apiUser.Name,
		Password: utils.Md5Encode(apiUser.Password),
	}

	if err := dbConn.Model(user).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func UserNameIsExists(name string) bool {
	user := &model.User{
		Name: name,
	}

	if r := dbConn.Model(user).Select("name").Where(user).Find(user).RowsAffected; r > 0 {
		return true
	}

	return false
}

func UserIDIsExists(id uint) bool {
	user := &model.User{
		ID: id,
	}

	if r := dbConn.Model(user).Select("id").Where(user).Find(user).RowsAffected; r > 0 {
		return true
	}

	return false
}

func QueryUserByName(name string) (*model.User, error) {
	result := &model.User{}

	if err := model.GetDBConnection().Select("id", "unit", "password").Where("name = ?", name).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
