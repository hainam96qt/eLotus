package model

type CreateRegistrationRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
