package main

import (
	"fmt"
	"golang-category-management/app"
	"golang-category-management/config"
	impl3 "golang-category-management/controller/impl"
	"golang-category-management/helper"
	"golang-category-management/repository/impl"
	impl2 "golang-category-management/service/impl"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error happened", err)
		}
	}()

	db := config.DatabaseConnection()
	newValidator := validator.New()
	defer func() {
		err := db.Close()
		helper.HelperPanic(err)
	}()

	categoryRepository := impl.NewCategoryRepository()
	categoryService := impl2.NewCategoryService(db, newValidator, categoryRepository)
	categoryController := impl3.NewCategoryController(categoryService)

	r := app.Router(categoryController)
	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Server is running on localhost:8080/")
	err := server.ListenAndServe()
	helper.HelperPanic(err)
}
