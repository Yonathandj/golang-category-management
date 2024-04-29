package repository

import (
	"context"
	"database/sql"
	"golang-category-management/model/entity"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}
