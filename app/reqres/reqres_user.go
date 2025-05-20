package reqres

import (
	"project-name/app/models"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	TglLahir   string `json:"tgl_lahir"`
	Image      string `json:"image"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	RoleID     int    `json:"role_id"`
	IsVerify   bool   `json:"is_verify"`
	Prov       int    `json:"prov"`
	Kab        int    `json:"kab"`
	Kec        int    `json:"kec"`
	Kel        string `json:"kel"`
	PostalCode string `json:"postal_code"`
}

func (request UserRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Name, validation.Required),
	)
}

type UserResponse struct {
	models.CustomGormModel
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Image      string    `json:"image"`
	Gender     string    `json:"gender"`
	TglLahir   time.Time `json:"tgl_lahir"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	RoleID     int       `json:"role_id"`
	IsVerify   bool      `json:"is_verify"`
	Kel        string    `json:"kel"`
	PostalCode string    `json:"postal_code"`
	Status     int       `json:"status"`
}

type UserUpdateRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	TglLahir   string `json:"tgl_lahir"`
	Image      string `json:"image"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	RoleID     int    `json:"role_id"`
	IsVerify   bool   `json:"is_verify"`
	Prov       int    `json:"prov"`
	Kab        int    `json:"kab"`
	Kec        int    `json:"kec"`
	Kel        string `json:"kel"`
	PostalCode string `json:"postal_code"`
}
