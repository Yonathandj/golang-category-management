package impl

import (
	"context"
	"database/sql"
	"golang-category-management/exception"
	"golang-category-management/helper"
	"golang-category-management/model/entity"
	"golang-category-management/model/request"
	"golang-category-management/model/response"
	"golang-category-management/repository"
	"golang-category-management/service"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	Database           *sql.DB
	Validator          *validator.Validate
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(database *sql.DB, validator *validator.Validate, categoryRepository repository.CategoryRepository) service.CategoryService {
	return &CategoryServiceImpl{
		Database:           database,
		Validator:          validator,
		CategoryRepository: categoryRepository,
	}
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse {
	tx, err := c.Database.Begin()
	defer helper.CommitOrRollback(tx)
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
	defer helper.CommitOrRollback(tx)
	helper.HelperPanic(err)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	tx, err := c.Database.Begin()
	defer helper.CommitOrRollback(tx)
	helper.HelperPanic(err)

	categories := c.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}

func (c *CategoryServiceImpl) Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse {
	tx, err := c.Database.Begin()
	defer helper.CommitOrRollback(tx)
	helper.HelperPanic(err)

	_, err = c.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

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
	defer helper.CommitOrRollback(tx)
	helper.HelperPanic(err)

	_, err = c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	c.CategoryRepository.Delete(ctx, tx, categoryId)
}
