package model

import "github.com/dgrijalva/jwt-go"

type RegisterUser struct {
	Username  string `json:"username" validate:"required,min=5,max=100"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
	Firstname string `json:"firstname" validate:"required,min=5,max=100"`
	Lastname  string `json:"lastname" validate:"required,min=5,max=100"`
	Email     string `json:"email" validate:"required,email"`
}

type LoginUser struct {
	EmailOrUsername string `json:"email_or_username" validate:"required"`
	Password        string `json:"password" validate:"required,min=8,max=100"`
}

type JWTPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Id       string `json:"id"`
}

type JWTClaim struct {
	Id         string      `json:"id"`
	Username   string      `json:"username"`
	Email      string      `json:"email"`
	CustomData interface{} `json:"customData"`
	jwt.StandardClaims
}
