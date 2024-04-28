package impl

import (
	"golang-category-management/controller"
	"golang-category-management/helper"
	request2 "golang-category-management/model/request"
	"golang-category-management/model/response"
	"golang-category-management/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) controller.CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (c *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryRequest request2.CategoryCreateRequest
	helper.DecodeJSONBody(request, &categoryRequest)

	categoryResponse := c.CategoryService.Create(request.Context(), categoryRequest)
	standardResponse := response.StandardResponse{
		StatusCode: http.StatusCreated,
		Message:    "Successfully create new category",
		Data:       categoryResponse,
	}

	helper.EncodeJSONBody(writer, standardResponse)
}

func (c *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryRequest request2.CategoryUpdateRequest
	helper.DecodeJSONBody(request, &categoryRequest)

	categoryResponse := c.CategoryService.Update(request.Context(), categoryRequest)
	standardResponse := response.StandardResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully update category",
		Data:       categoryResponse,
	}

	helper.EncodeJSONBody(writer, standardResponse)
}

func (c *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	categoryId, err := strconv.Atoi(id)
	helper.HelperPanic(err)

	c.CategoryService.Delete(request.Context(), categoryId)
	standardResponse := response.StandardResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully delete category",
		Data:       nil,
	}

	helper.EncodeJSONBody(writer, standardResponse)
}

func (c *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	categoryId, err := strconv.Atoi(id)
	helper.HelperPanic(err)

	category := c.CategoryService.FindById(request.Context(), categoryId)
	standardResponse := response.StandardResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully find id category",
		Data:       category,
	}

	helper.EncodeJSONBody(writer, standardResponse)
}

func (c *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := c.CategoryService.FindAll(request.Context())
	standardResponse := response.StandardResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully find all category",
		Data:       categoryResponses,
	}

	helper.EncodeJSONBody(writer, standardResponse)
}
