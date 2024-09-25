package user

import (
	"NestJsStyle/data"
	"NestJsStyle/helper"
	"NestJsStyle/model"
	"NestJsStyle/prisma/db"
	"context"
	"net/http"
)

type UserServices struct {
	Db *db.PrismaClient
}

func NewUserServices(db *db.PrismaClient) *UserServices {
	return &UserServices{Db: db}
}

func (p *UserServices) Register(ctx context.Context, userDto model.RegisterUser) *data.WebResponse {
	err := helper.RequestValidators(userDto)
	if err != nil {
		return &data.WebResponse{
			Data:   err.Error(),
			Code:   http.StatusBadRequest,
			Status: "Validation errors",
		}
	}

	existingUserByEmailOrUsername, _ := p.Db.User.FindFirst(
		db.User.Or(
			db.User.Username.Equals(userDto.Username),
			db.User.Email.Equals(userDto.Email),
		),
	).Exec(ctx)

	if existingUserByEmailOrUsername != nil {
		return &data.WebResponse{
			Code:   http.StatusBadRequest,
			Data:   nil,
			Status: "User Already Exists",
		}
	}
	hashedPassword, err := helper.HashPassword(userDto.Password)
	if err != nil {
		return &data.WebResponse{
			Data:   err.Error(),
			Code:   http.StatusBadRequest,
			Status: "Could not hash password",
		}
	}

	_, err = p.Db.User.CreateOne(
		db.User.Username.Set(userDto.Username),
		db.User.Email.Set(userDto.Email),
		db.User.Firstname.Set(userDto.Firstname),
		db.User.Lastname.Set(userDto.Lastname),
		db.User.Password.Set(hashedPassword),
	).Exec(ctx)

	if err != nil {
		return &data.WebResponse{
			Data:   nil,
			Code:   http.StatusBadRequest,
			Status: "Could not register user",
		}
	}
	return &data.WebResponse{
		Data:   nil,
		Code:   http.StatusOK,
		Status: "Success",
	}

}

func (p *UserServices) Login(ctx context.Context, userDto model.LoginUser) *data.LoginResponse {
	err := helper.RequestValidators(userDto)
	if err != nil {
		return &data.LoginResponse{
			Code:        http.StatusBadRequest,
			Status:      "Validation errors",
			AccessToken: nil,
		}
	}
	user, _ := p.Db.User.FindFirst(
		db.User.Or(
			db.User.Email.Equals(userDto.EmailOrUsername),
			db.User.Username.Equals(userDto.EmailOrUsername),
		)).Exec(ctx)
	if user == nil {
		return &data.LoginResponse{
			Code:        http.StatusBadRequest,
			Status:      "User Not Found",
			AccessToken: nil,
		}
	}
	isValidPassword := helper.CheckPasswordHash(userDto.Password, user.Password)
	if !isValidPassword {
		return &data.LoginResponse{
			Code:        http.StatusBadRequest,
			Status:      "Invalid Password",
			AccessToken: nil,
		}
	}
	jwtPayload := &model.JWTPayload{
		Username: user.Username,
		Email:    user.Email,
		Id:       user.ID,
	}
	tokenString, err := helper.GenerateJwt(jwtPayload)
	if err != nil {
		return &data.LoginResponse{
			Code:        http.StatusBadRequest,
			Status:      "Bad request error",
			AccessToken: nil,
		}
	}

	return &data.LoginResponse{
		Code:        http.StatusOK,
		Status:      "Success",
		AccessToken: tokenString,
	}

}

func (p *UserServices) UserProfile(ctx context.Context, userId string) *data.WebResponse {
	user, err := p.Db.User.FindUnique(db.User.ID.Equals(userId)).Omit(
		db.User.Password.Field(),
		db.User.CreatedAt.Field(),
		db.User.UpdatedAt.Field(),
	).Exec(ctx)
	if err != nil {
		return &data.WebResponse{
			Data:   nil,
			Code:   http.StatusBadRequest,
			Status: "User Not Found",
		}
	}
	return &data.WebResponse{
		Data: &data.UserProfile{
			Id:        user.ID,
			FirstName: user.Firstname,
			LastName:  user.Lastname,
			Username:  user.Username,
			Email:     user.Email,
		},
		Code:   http.StatusOK,
		Status: "Success",
	}
}
