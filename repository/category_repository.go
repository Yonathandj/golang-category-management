package repository

import (
	"context"
	"database/sql"
	"golang-category-management/model/entity"
	"golang-category-management/model/request"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, request request.CategoryCreateRequest) entity.Category
	FindById(ctx context.Context, tx *sql.Tx, id int) entity.Category
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
	Update(ctx context.Context, tx *sql.Tx, request request.CategoryUpdateRequest) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, id int)
}
