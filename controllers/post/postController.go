package post

import (
	"NestJsStyle/helper"
	"NestJsStyle/model"
	"NestJsStyle/services/post"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type PostControllerInjection struct {
	Post *post.PrismaInjection
}

func NewPostControllerInjection(post *post.PrismaInjection) *PostControllerInjection {
	return &PostControllerInjection{Post: post}
}

func (controller *PostControllerInjection) CreatePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postCreate := model.Post{}
	helper.ReadRequestBody(r, &postCreate)
	userId := r.Context().Value("userId").(string)
	postCreate.UserId = userId

	webResponse := controller.Post.CreatePostService(r.Context(), postCreate)
	helper.WriteResponseBody(w, webResponse, webResponse.Code)
}

func (controller *PostControllerInjection) UpdatePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postUpdate := model.PostUpdate{}
	helper.ReadRequestBody(r, &postUpdate)
	userId := r.Context().Value("userId").(string)
	postId := params.ByName("postId")
	postUpdate.Id = postId
	postUpdate.UserId = userId

	webResponse := controller.Post.UpdatePost(r.Context(), postUpdate)
	helper.WriteResponseBody(w, webResponse, webResponse.Code)
}

func (controller *PostControllerInjection) DeletePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	webResponse := controller.Post.DeletePost(r.Context(), postId)
	helper.WriteResponseBody(w, webResponse, webResponse.Code)
}

func (controller *PostControllerInjection) GetAllPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	pageNumber, _ := strconv.Atoi(page)
	limitNumber, _ := strconv.Atoi(limit)
	webResponse := controller.Post.GetAllPost(r.Context(), pageNumber, limitNumber)
	helper.WriteResponseBody(w, webResponse, webResponse.Code)
}

func (controller *PostControllerInjection) GetPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	webResponse := controller.Post.FindPostById(r.Context(), postId)
	helper.WriteResponseBody(w, webResponse, webResponse.Code)
}
