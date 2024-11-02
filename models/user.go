package models

import "gorm.io/gorm"

type Users struct {
	*gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	IDCode      string `json:"id_code"`
	IsAdmin     bool   `json:"is_admin"`
	Gender      bool   `json:"gender"`
	IsHidden    bool   `json:"is_hidden"`
}
