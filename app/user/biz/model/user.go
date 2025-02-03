package model

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Email    string `gorm:"column:email;unique;not null"`
	Username string `gorm:"column:username;unique;not null"`
	Nickname string `gorm:"column:nickname;"`
	Password string `gorm:"column:password;"`
	gorm.Model
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}

func GetByUserName(db *gorm.DB, ctx context.Context, userName string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Username: userName}).First(&user).Error
	return
}
