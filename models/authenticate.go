package models

type LoginModel struct {
	UserName string
	Password string
}

type PasswordResetModel struct {
	Password string
}
