package controllers

import (
	"net/http"
	"project-name/app/middlewares"
	"project-name/app/repository"
	"project-name/app/reqres"
	"project-name/app/utils"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

// LoginUser godoc
// @Summary Login User
// @Description Login User
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body reqres.LoginRequest true "Login Request"
// @Success 200
// @Router /v1/auth/login/user [post]
// @Security ApiKeyAuth
func LoginUser(c echo.Context) error {
	var data reqres.LoginRequest
	if err := c.Bind(&data); err != nil {
		return utils.NewBadRequestError("Invalid request body")
	}

	if err := data.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	user, t, err := repository.Login(data.EmailOrPhone)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Invalid email"))
	}

	if err = middlewares.VerifyPassword(data.Password, user.Password); err != nil {
		return c.JSON(400, utils.Respond(400, err, "Invalid password"))
	}

	if user.RoleID != 3 {
		return c.JSON(400, utils.Respond(400, err, "You are not a user"))
	}

	userResponse, _ := repository.GetUserByID(int(user.ID))

	dataResponse := reqres.LoginResponse{
		Token: t,
		User:  userResponse,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"data":    dataResponse,
		"message": "Login Success",
	})
}

// LoginAdmin godoc
// @Summary Login Admin
// @Description Login Admin
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body reqres.LoginRequest true "Login Request"
// @Success 200
// @Router /v1/auth/login/admin [post]
// @Security ApiKeyAuth
func LoginAdmin(c echo.Context) error {
	var data reqres.LoginRequest
	if err := c.Bind(&data); err != nil {
		return utils.NewBadRequestError("Invalid request body")
	}

	if err := data.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	user, t, err := repository.Login(data.EmailOrPhone)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Invalid email"))
	}

	if err = middlewares.VerifyPassword(data.Password, user.Password); err != nil {
		return c.JSON(400, utils.Respond(400, err, "Invalid password"))
	}

	if user.RoleID == 3 {
		return c.JSON(400, utils.Respond(400, err, "You are not admin"))
	}

	userResponse, _ := repository.GetUserByID(int(user.ID))

	dataResponse := reqres.LoginResponse{
		Token: t,
		User:  userResponse,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"data":    dataResponse,
		"message": "Login Success",
	})
}

// Register godoc
// @Summary Register
// @Description Register
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body reqres.UserRequest true "Register Request"
// @Success 200
// @Router /v1/auth/register [post]
func Register(c echo.Context) error {
	var data reqres.UserRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := data.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, utils.NewInvalidInputError(errVal))
	}

	users, _ := repository.GetAllUsers()

	for _, dataUser := range users {
		if dataUser.Email == data.Email {
			return c.JSON(400, utils.NewBadRequestError("Email already exists"))
		}

		if dataUser.Phone == data.Phone {
			return c.JSON(400, utils.NewBadRequestError("Phone already exists"))
		}
	}

	user, err := repository.Register(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to create user"))
	}

	userResponse, err := repository.GetUserByID(int(user.ID))
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	// go repository.SendEmailVerificationEmail(int(user.ID))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"data":    userResponse,
		"message": "Register Success",
	})
}

// ForgotPassword godoc
// @Summary Forgot Password
// @Description Forgot Password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body reqres.ForgotPasswordRequest true "Forgot Password Request"
// @Success 200
// @Router /v1/auth/forgot-password [post]
func ForgotPassword(c echo.Context) error {
	var data reqres.ForgotPasswordRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	user, err := repository.GetUserByEmail(data.Email)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Email not found"))
	}

	// go repository.UserForgotPasswordNotification(int(user.ID))

	return c.JSON(200, map[string]interface{}{
		"status":  200,
		"data":    user,
		"message": "Permintaan Reset Password Berhasil Silahkan Cek Email Anda",
	})
}

// ResetPassword godoc
// @Summary Reset Password
// @Description Reset Password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param req body reqres.ChangePassword true "Reset Password Request"
// @Success 200
// @Router /v1/auth/reset-password/{id} [put]
func ResetPassword(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.GetUserByIDPlain(id)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	var req reqres.ChangePassword
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	if req.NewPassword == "" || req.NewPasswordConfirm == "" {
		return c.JSON(400, utils.NewUnprocessableEntityError("New Password and New Password Confirm cannot be empty"))
	}

	if req.NewPassword != req.NewPasswordConfirm {
		return c.JSON(400, utils.NewUnprocessableEntityError("New Password and New Password Confirm must be same"))
	}

	newPassword := middlewares.BcryptPassword(req.NewPassword)

	data.Password = newPassword

	update, err := repository.UpdateUser(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to update user"))
	}

	dataUpdate, err := repository.GetUserByID(int(update.ID))
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"data":    dataUpdate,
		"message": "Password Berhasil Diubah",
	})
}

// SendEmailVerifyEmail godoc
// @Summary Send Email Verify Email
// @Description Send Email Verify Email
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/auth/email-verify [post]
// @Security JwtToken
func SendEmailVerifyEmail(c echo.Context) error {
	// userID := c.Get("user_id").(int)

	// go repository.SendEmailVerificationEmail(userID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "Permintaan Verifikasi Email Berhasil Silahkan Cek Email Anda",
	})
}

// AktivateAccount godoc
// @Summary Aktivate Account
// @Description Aktivate Account
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /v1/auth/activate-account/{id} [put]
func AktivateAccount(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	data, err := repository.GetUserByIDPlain(userID)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	data.IsVerify = true

	update, err := repository.UpdateUser(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to update user"))
	}

	dataUpdate, err := repository.GetUserByID(int(update.ID))
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"data":    dataUpdate,
		"message": "Permintaan Verifikasi Email Berhasil",
	})
}

// ChangePasswordLogin godoc
// @Summary Change Password Login
// @Description Change Password Login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param req body reqres.ChangePassword true "Change Password Request"
// @Success 200
// @Router /v1/auth/change-password-login [put]
// @Security JwtToken
func ChangePasswordLogin(c echo.Context) error {
	userID := c.Get("user_id").(int)

	data, err := repository.GetUserByIDPlain(userID)
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	var req reqres.ChangePassword
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, utils.NewUnprocessableEntityError(err.Error()))
	}

	if req.NewPassword == "" || req.NewPasswordConfirm == "" {
		return c.JSON(400, utils.NewUnprocessableEntityError("New Password and New Password Confirm cannot be empty"))
	}

	if req.NewPassword != req.NewPasswordConfirm {
		return c.JSON(400, utils.NewUnprocessableEntityError("New Password and New Password Confirm must be same"))
	}

	newPassword := middlewares.BcryptPassword(req.NewPassword)

	data.Password = newPassword

	update, err := repository.UpdateUser(data)
	if err != nil {
		return c.JSON(500, utils.Respond(500, err, "Failed to update user"))
	}

	dataUpdate, err := repository.GetUserByID(int(update.ID))
	if err != nil {
		return c.JSON(400, utils.Respond(400, err, "Failed to get user"))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"data":    dataUpdate,
		"message": "Password Berhasil Diubah",
	})
}
