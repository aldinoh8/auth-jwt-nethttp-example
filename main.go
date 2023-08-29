package main

import (
	"example/config"
	"example/controller"
	"example/middleware"
	"example/repository"
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.InitDb(os.Getenv("DB"))
	router := httprouter.New()

	userController := controller.User{
		Repository: repository.User{DB: db},
	}
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	productController := controller.Product{}
	router.GET("/product", productController.Index)
	router.GET("/product/auth",
		middleware.Authentication(productController.Auth, db),
	)
	router.GET("/product/superadmin",
		middleware.Authentication(
			middleware.AuthorizeSuperAdmin(productController.SuperAdmin),
			db),
	)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.Log(router),
	}

	fmt.Println("server running on port 8080")
	server.ListenAndServe()
}
