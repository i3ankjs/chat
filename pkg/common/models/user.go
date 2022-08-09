package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId             uint    `gorm:"primary key;autoIncrement" json:"user_id"`
	Name               *string `json:"name"`
	Email              *string `json:"email"`
	Password           *string `json:"password"`
	Pic                *string `json:"pic"`
	Address            *string `json:"address"`
	Telephone          *string `json:"telephone"`
	Status             *string `json:"status"`
	ShopId             *string `json:"shop_id"`
	GroupId            *string `json:"group_id"`
	TokenResetPassword *string `json:"token_reset_password"`
	DateAdded          *string `json:"date_added"`
	DateModified       *string `json:"date_modified"`
}
