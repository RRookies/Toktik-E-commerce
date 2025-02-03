package model

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	UserId int `json:"user"`
	jwt.RegisteredClaims
}
