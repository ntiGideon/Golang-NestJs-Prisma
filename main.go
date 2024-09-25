package main

import (
	"NestJsStyle/config"
	"NestJsStyle/controllers/post"
	userCont "NestJsStyle/controllers/user"
	"NestJsStyle/helper"
	"NestJsStyle/router"
	post2 "NestJsStyle/services/post"
	userServ "NestJsStyle/services/user"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Sorry, we can't proceed!")
	}
	fmt.Printf("Server starting on PORT %s \n", os.Getenv("PORT"))

	// handle db
	db, err := config.ConnectDB()
	if err != nil {
		helper.PanicAllErrors(err)
	}

	defer db.Prisma.Disconnect()

	postService := post2.NewPrismaInjection(db)
	postController := post.NewPostControllerInjection(&postService)
	userService := userServ.NewUserServices(db)
	userController := userCont.NewUserController(userService)

	routes := router.NewRouter(postController, userController)

	//server guy
	server := http.Server{
		Addr:           os.Getenv("PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	serverError := server.ListenAndServe()
	if serverError != nil {
		helper.PanicAllErrors(serverError)
	}
}
