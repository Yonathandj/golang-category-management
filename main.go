package main

import (
	"fmt"
	"golang-category-management/app"
	"golang-category-management/config"
	controller "golang-category-management/controller/impl"
	"golang-category-management/helper"
	repository "golang-category-management/repository/impl"
	service "golang-category-management/service/impl"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error happened", err)
		}
	}()

	newDatabase := config.DatabaseConnection()
	newValidator := validator.New()
	defer func() {
		err := newDatabase.Close()
		helper.HelperPanic(err)
	}()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(newDatabase, newValidator, categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	r := app.Router(categoryController)
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server is running on localhost:8080/")
	err := server.ListenAndServe()
	helper.HelperPanic(err)
}
