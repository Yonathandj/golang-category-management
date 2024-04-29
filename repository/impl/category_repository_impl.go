package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-category-management/helper"
	"golang-category-management/model/entity"
	"golang-category-management/repository"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() repository.CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	sql := "INSERT INTO categories (name) VALUES ($1) RETURNING id"

	var id int64
	err := tx.QueryRowContext(ctx, sql, category.Name).Scan(&id)
	helper.HelperPanic(err)

	category.Id = int(id)
	return category
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entity.Category, error) {
	sql := "SELECT id, name FROM categories WHERE id= $1"

	rows, err := tx.QueryContext(ctx, sql, categoryId)
	defer func() {
		err := rows.Close()
		helper.HelperPanic(err)
	}()
	helper.HelperPanic(err)

	var category entity.Category
	if rows.Next() {
		err = rows.Scan(&category.Id, &category.Name)
		helper.HelperPanic(err)

		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	sql := "SELECT id, name FROM categories"

	rows, err := tx.QueryContext(ctx, sql)
	defer func() {
		err := rows.Close()
		helper.HelperPanic(err)
	}()
	helper.HelperPanic(err)

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.HelperPanic(err)
		categories = append(categories, category)
	}
	return categories
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	sql := "UPDATE categories SET name = $1 WHERE id = $2"
	
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.HelperPanic(err)

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	sql := "DELETE FROM categories WHERE id = $1"

	_, err := tx.ExecContext(ctx, sql, categoryId)
	helper.HelperPanic(err)
}
