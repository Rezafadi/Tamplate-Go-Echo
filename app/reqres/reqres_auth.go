package reqres

import validation "github.com/go-ozzo/ozzo-validation"

type LoginRequest struct {
	EmailOrPhone string `json:"emailorphone"`
	Password     string `json:"password"`
}

func (request *LoginRequest) Validate() error {
	return validation.ValidateStruct(
		request,
		validation.Field(&request.EmailOrPhone, validation.Required, validation.Length(1, 0)),
		validation.Field(&request.Password, validation.Required, validation.Length(1, 0)),
	)
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ChangePassword struct {
	NewPassword        string `json:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm"`
}
