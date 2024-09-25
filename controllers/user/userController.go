package user

import (
	"NestJsStyle/helper"
	"NestJsStyle/model"
	"NestJsStyle/services/user"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController struct {
	User *user.UserServices
}

func NewUserController(user *user.UserServices) *UserController {
	return &UserController{User: user}
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userDto := model.RegisterUser{}
	helper.ReadRequestBody(r, &userDto)

	webResponse := controller.User.Register(r.Context(), userDto)
	helper.WriteResponseBody(w, webResponse, webResponse.Code)
}

func (controller *UserController) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userDto := model.LoginUser{}
	helper.ReadRequestBody(r, &userDto)
	webResponse := controller.User.Login(r.Context(), userDto)
	helper.WriteResponseBody(w, webResponse, webResponse.Code)
}

func (controller *UserController) UserProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := r.Context().Value("userId").(string)
	webResponse := controller.User.UserProfile(r.Context(), userId)
	helper.WriteResponseBody(w, webResponse, webResponse.Code)
}
