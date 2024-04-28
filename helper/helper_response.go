package helper

import (
	"golang-category-management/model/entity"
	"golang-category-management/model/response"
)

func ToCategoryResponse(category entity.Category) response.CategoryResponse {
	return response.CategoryResponse{
		CategoryId:   category.Id,
		CategoryName: category.Name,
	}
}

func ToCategoryResponses(categories []entity.Category) []response.CategoryResponse {
	var categoriesResponse []response.CategoryResponse
	for _, v := range categories {
		categoriesResponse = append(categoriesResponse, response.CategoryResponse{CategoryId: v.Id, CategoryName: v.Name})
	}
	return categoriesResponse
}
