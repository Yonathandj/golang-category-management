package app

import (
	"golang-category-management/controller"
	"golang-category-management/exception"

	"github.com/julienschmidt/httprouter"
)

func Router(categoryController controller.CategoryController) *httprouter.Router {
	r := httprouter.New()

	r.GET("/categories", categoryController.FindAll)
	r.GET("/categories/:id", categoryController.FindById)
	r.POST("/categories", categoryController.Create)
	r.PUT("/categories", categoryController.Update)
	r.DELETE("/categories/:id", categoryController.Delete)

	r.PanicHandler = exception.ErrorHandler

	return r
}
