package impl

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang-category-management/helper"
	"golang-category-management/model/entity"
	"golang-category-management/model/request"
	"golang-category-management/model/response"
	"golang-category-management/repository"
)

type CategoryServiceImpl struct {
	Database           *sql.DB
	Validator          *validator.Validate
	CategoryRepository repository.CategoryRepository
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse {
	tx, err := c.Database.Begin()
	helper.HelperPanic(err)

	err = c.Validator.Struct(request)
	helper.HelperPanic(err)

	category := entity.Category{
		Name: request.Name,
	}
	category = c.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) response.CategoryResponse {
	tx, err := c.Database.Begin()
	helper.HelperPanic(err)

	if categoryId == 0 {
		helper.HelperPanic(err)
	}
	category := c.CategoryRepository.FindById(ctx, tx, categoryId)
	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) findAll(ctx context.Context) []response.CategoryResponse {
	tx, err := c.Database.Begin()
	helper.HelperPanic(err)

	categories := c.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}

func (c *CategoryServiceImpl) Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse {
	tx, err := c.Database.Begin()
	helper.HelperPanic(err)

	err = c.Validator.Struct(request)
	helper.HelperPanic(err)

	category := entity.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	category = c.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := c.Database.Begin()
	helper.HelperPanic(err)

	if categoryId == 0 {
		helper.HelperPanic(err)
	}
	c.CategoryRepository.Delete(ctx, tx, categoryId)
}
