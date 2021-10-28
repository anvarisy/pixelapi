package models

import (
	"time"

	"github.com/anvarisy/pixelapi/auths"
	"gorm.io/gorm"
)

type User struct {
	Username     string         `gorm:"not null; primaryKey; size:16" json:"username" form:"username"`
	UserPassword string         `gorm:"not null" json:"user_password" form:"user_password"`
	UserFullname string         `json:"user_fullname" form:"user_fullname"`
	UserMobile   string         `json:"user_mobile" form:"user_mobile"`
	IsAdmin      bool           `json:"is_admin" form:"is_admin" gorm:"default:false"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	Deleted      gorm.DeletedAt `gorm:"index" json:"-"`
}

type UserLogin struct {
	Username     string `json:"username" form:"username"`
	UserPassword string `json:"user_password" form:"user_password"`
}

type UserLoginSuccess struct {
	Username     string `json:"username" form:"username"`
	UserFullname string `json:"user_fullname" form:"user_fullname"`
	UserMobile   string `json:"user_mobile" form:"user_mobile"`
	UserToken    string `json:"user_token" form:"user_token"`
}

type UserRegister struct {
	Username     string `gorm:"not null; primaryKey" json:"username" form:"username"`
	UserPassword string `gorm:"not null" json:"user_password" form:"user_password"`
	UserFullname string `json:"user_fullname" form:"user_fullname"`
	UserMobile   string `json:"user_mobile" form:"user_mobile"`
}

type UserRegisterSuccess struct {
	Username     string `json:"username" form:"username"`
	UserFullname string `json:"user_fullname" form:"user_fullname"`
	UserMobile   string `json:"user_mobile" form:"user_mobile"`
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {

	err := db.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	hashedPassword, err := auths.Hash(u.UserPassword)
	if err != nil {
		return err
	}
	u.UserPassword = string(hashedPassword)
	return nil
}
