package controllers

import (
	"net/http"
	"project-name/app/repository"
	"project-name/app/reqres"
	"project-name/app/utils"
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body reqres.UserRequest true "Create User Request"
// @Success 200
// @Router /v1/user [post]
// @Security JwtToken
func CreateUser(c echo.Context) error {
	var data reqres.UserRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := data.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	if data.Email != "" {
		email, _ := repository.GetUserByEmail(data.Email)
		if email.Email != "" {
			return c.JSON(400, utils.NewBadRequestError("Email already exists"))
		}
	}

	if data.Phone != "" {
		phone, _ := repository.GetUserByPhone(data.Phone)
		if phone.Phone != "" {
			return c.JSON(400, utils.NewBadRequestError("Phone already exists"))
		}
	}

	var tglLahir time.Time
	var err error
	if data.TglLahir != "" {
		tglLahir, err = time.Parse("2006-01-02 15:04:05", data.TglLahir)
		if err != nil {
			tglLahir, err = time.Parse("2006-01-02", data.TglLahir)
			if err != nil {
				return c.JSON(500, utils.Respond(500, err, "Invalid Tanggal Lahir format"))
			}
		}
	}

	user, err := repository.CreateUser(tglLahir, data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to create user"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    user,
		"message": "Create User Success",
	})
}

// GetUsers godoc
// @Summary Get Users
// @Description Get Users
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200
// @Param role_id query int false "Role ID"
// @Param status query string false "Status"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param search query string false "Search"
// @Param sort query string false "Sort"
// @Param order query string false "Order"
// @Router /v1/user [get]
// @Security JwtToken
func GetUsers(c echo.Context) error {
	roleID, _ := strconv.Atoi(c.QueryParam("role_id"))
	param := utils.PopulatePaging(c, "status")

	data := repository.GetUsers(roleID, param)

	return c.JSON(200, data)
}

// GetAllUsers godoc
// @Summary Get All Users
// @Description Get All Users
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/user/all [get]
// @Security JwtToken
func GetAllUsers(c echo.Context) error {

	users, err := repository.GetAllUsers()
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to get users"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    users,
		"message": "Get All Users Success",
	})
}

// GetUserByID godoc
// @Summary Get User By ID
// @Description Get User By ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200
// @Router /v1/user/{id} [get]
// @Security JwtToken
func GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.GetUserByID(id)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    data,
		"message": "Get User Success",
	})
}

// UpdateUser godoc
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param request body reqres.UserUpdateRequest true "Update User Request"
// @Success 200
// @Router /v1/user/{id} [put]
// @Security JwtToken
func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.GetUserByIDPlain(id)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	var req reqres.UserUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	if req.Name != "" {
		data.Name = req.Name
	}
	if req.Email != "" {
		email, _ := repository.GetUserByEmail(req.Email)
		if email.Email != "" {
			if req.Email == email.Email && data.Email != email.Email {
				return c.JSON(400, utils.NewBadRequestError("Email already exists"))
			}
		}
		data.Email = req.Email
	}
	if req.Gender != "" {
		data.Gender = req.Gender
	}
	if req.TglLahir != "" {
		tglLahir, err := time.Parse("2006-01-02 15:04:05", req.TglLahir)
		if err != nil {
			tglLahir, err = time.Parse("2006-01-02", req.TglLahir)
			if err != nil {
				return c.JSON(500, utils.Respond(500, err, "Invalid Tanggal Lahir format"))
			}
		}

		data.TglLahir = tglLahir
	}
	if req.Image != "" {
		data.Image = req.Image
	}
	if req.Phone != "" {
		phone, _ := repository.GetUserByPhone(req.Phone)
		if phone.Phone != "" {
			if req.Phone == phone.Phone && data.Phone != phone.Phone {
				return c.JSON(400, utils.NewBadRequestError("Phone already exists"))
			}
		}
		data.Phone = req.Phone
	}
	if req.Address != "" {
		data.Address = req.Address
	}
	if req.RoleID != 0 {
		data.RoleID = req.RoleID
	}
	if req.Prov != 0 {
		data.Prov = req.Prov
	}
	if req.Kab != 0 {
		data.Kab = req.Kab
	}
	if req.Kec != 0 {
		data.Kec = req.Kec
	}
	if req.Kel != "" {
		data.Kel = req.Kel
	}
	if req.PostalCode != "" {
		data.PostalCode = req.PostalCode
	}
	data.IsVerify = req.IsVerify
	data.Status = 0
	if data.IsVerify {
		data.Status = 1
	}

	update, err := repository.UpdateUser(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to update user"))
	}

	dataUpdate, err := repository.GetUserByID(int(update.ID))
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    dataUpdate,
		"message": "Update User Success",
	})
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200
// @Router /v1/user/{id} [delete]
// @Security JwtToken
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.GetUserByIDPlain(id)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	dataResponse, err := repository.GetUserByID(id)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	_, err = repository.DeleteUser(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to delete user"))
	}

	DeleteFile(data.Image)

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    dataResponse,
		"message": "Delete User Success",
	})
}
