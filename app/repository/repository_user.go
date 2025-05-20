package repository

import (
	"project-name/app/middlewares"
	"project-name/app/models"
	"project-name/app/reqres"
	"project-name/app/utils"
	"project-name/config"
	"strconv"
	"time"
)

func CreateUser(tglLahir time.Time, data reqres.UserRequest) (response models.User, err error) {
	password := middlewares.BcryptPassword(data.Password)

	response = models.User{
		Name:       data.Name,
		Email:      data.Email,
		Password:   password,
		Gender:     data.Gender,
		Phone:      data.Phone,
		TglLahir:   tglLahir,
		Image:      data.Image,
		Address:    data.Address,
		IsVerify:   true,
		RoleID:     data.RoleID,
		Status:     0,
		Prov:       data.Prov,
		Kab:        data.Kab,
		Kec:        data.Kec,
		Kel:        data.Kel,
		PostalCode: data.PostalCode,
	}

	err = config.DB.Create(&response).Error

	return
}

func BuildUserResponse(data models.User) (response reqres.UserResponse) {
	response = reqres.UserResponse{
		CustomGormModel: data.CustomGormModel,
		Name:            data.Name,
		Email:           data.Email,
		Image:           config.LoadConfig().BaseUrl + "/" + config.LoadConfig().DirPath + data.Image,
		Gender:          data.Gender,
		TglLahir:        data.TglLahir,
		Phone:           data.Phone,
		Address:         data.Address,
		RoleID:          data.RoleID,
		IsVerify:        data.IsVerify,
		Status:          data.Status,
		PostalCode:      data.PostalCode,
		Kel:             data.Kel,
	}

	return
}

func GetUsers(roleID int, param reqres.ReqPaging) (data reqres.ResPaging) {
	var out []models.User
	where := "deleted_at IS NULL"

	if roleID != 0 {
		where += " AND role_id = " + strconv.Itoa(roleID)
	}
	if param.Search != "" {
		where += " AND (name ILIKE '%" + param.Search + "%' OR email ILIKE '%" + param.Search + "%' OR phone ILIKE '%" + param.Search + "%')"
	}
	if param.Custom != "" {
		where += " AND status = " + param.Custom.(string)
	}

	var modelTotal []models.User
	var totalResult int64
	config.DB.Model(&modelTotal).Where(where).Count(&totalResult)

	var totalFiltered int64
	config.DB.Model(&modelTotal).Where(where).Count(&totalFiltered)

	config.DB.Where(where).Offset(param.Offset).Order(param.Sort + " " + param.Order).Limit(param.Limit).Find(&out)

	var responses []reqres.UserResponse
	for _, response := range out {
		responses = append(responses, BuildUserResponse(response))
	}

	data = utils.PopulateResPaging(&param, responses, totalResult, totalFiltered)

	return
}

func GetAllUsers() (data []reqres.UserResponse, err error) {
	var out []models.User

	err = config.DB.Find(&out).Error

	for _, response := range out {
		data = append(data, BuildUserResponse(response))
	}

	return
}

func GetUserByID(id int) (data reqres.UserResponse, err error) {
	var out models.User

	err = config.DB.First(&out, id).Error

	data = BuildUserResponse(out)

	return
}

func GetUserByIDPlain(id int) (data models.User, err error) {
	err = config.DB.First(&data, id).Error

	return
}

func GetUserByEmail(email string) (data models.User, err error) {
	err = config.DB.Where("email = ?", email).First(&data).Error

	return
}

func GetUserByPhone(phone string) (data models.User, err error) {
	err = config.DB.Where("phone = ?", phone).First(&data).Error

	return
}

func UpdateUser(data models.User) (response models.User, err error) {

	err = config.DB.Save(&data).Scan(&response).Error

	return
}

func DeleteUser(data models.User) (response models.User, err error) {

	err = config.DB.Unscoped().Delete(&data).Error

	return
}
