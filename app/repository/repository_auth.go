package repository

import (
	"project-name/app/middlewares"
	"project-name/app/models"
	"project-name/app/reqres"
	"project-name/config"
)

func Login(emailorphone string) (data models.User, token string, err error) {
	err = config.DB.Debug().Where("email = ? OR phone = ?", emailorphone, emailorphone).First(&data).Error
	if err != nil {
		return
	}

	token, err = middlewares.AuthMakeToken(data)
	if err != nil {
		return
	}

	return
}

func Register(data reqres.UserRequest) (response models.User, err error) {
	password := middlewares.BcryptPassword(data.Password)

	response = models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: password,
		Gender:   data.Gender,
		Phone:    data.Phone,
		Image:    data.Image,
		Address:  data.Address,
		IsVerify: data.IsVerify,
		RoleID:   3,
		Status:   0,
	}

	err = config.DB.Create(&response).Error

	return
}
