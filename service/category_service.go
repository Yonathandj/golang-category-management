package service

import (
	"context"
	"golang-category-management/model/request"
	"golang-category-management/model/response"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	FindById(ctx context.Context, categoryId int) response.CategoryResponse
	findAll(ctx context.Context) []response.CategoryResponse
	Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
}
