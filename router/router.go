package router

import (
	"NestJsStyle/controllers/post"
	"NestJsStyle/controllers/user"
	"NestJsStyle/middleware"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(postController *post.PostControllerInjection, userController *user.UserController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/post", middleware.AuthMiddleware(postController.CreatePost))
	router.PUT("/api/post/:postId", middleware.AuthMiddleware(postController.UpdatePost))
	router.DELETE("/api/post/:postId", middleware.AuthMiddleware(postController.DeletePost))
	router.GET("/api/post", middleware.AuthMiddleware(postController.GetAllPost))
	router.GET("/api/post/:postId", middleware.AuthMiddleware(postController.GetPost))
	router.POST("/api/user/register", userController.CreateUser)
	router.POST("/api/user/login", userController.Login)
	router.GET("/api/user/profile", middleware.AuthMiddleware(userController.UserProfile))
	return router
}
