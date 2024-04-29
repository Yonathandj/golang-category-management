package app

import (
	"github.com/julienschmidt/httprouter"
	"golang-category-management/controller"
	"golang-category-management/exception"
)

func Router(categoryController controller.CategoryController) *httprouter.Router {
	r := httprouter.New()

	r.GET("/categories", categoryController.FindAll)
	r.GET("/categories/:id", categoryController.FindById)
	r.POST("/categories", categoryController.Create)
	r.PUT("/categories", categoryController.Update)
	r.DELETE("/categories/:id", categoryController.Delete)

	r.PanicHandler = exception.ExceptionHandler

	return r
}
