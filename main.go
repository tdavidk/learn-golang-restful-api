package main

import (
	"net/http"
	"tdavidk/learn-golang-restful-api/app"
	"tdavidk/learn-golang-restful-api/controller"
	"tdavidk/learn-golang-restful-api/helper"
	"tdavidk/learn-golang-restful-api/middleware"
	"tdavidk/learn-golang-restful-api/repository"
	"tdavidk/learn-golang-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
