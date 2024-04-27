package impl

import (
	"context"
	"database/sql"
	"errors"
	"golang-category-management/helper"
	"golang-category-management/model/entity"
)

type CategoryRepositoryImpl struct {
}

func (c CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	sql := "INSERT INTO categories (name) VALUES (?)"

	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.HelperPanic(err)

	var id int64
	id, err = result.LastInsertId()
	helper.HelperPanic(err)

	category.Id = int(id)
	return category
}

func (c CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) entity.Category {
	sql := "SELECT id, name FROM categories WHERE id= ?"
	category := entity.Category{}

	err := tx.QueryRowContext(ctx, sql, categoryId).Scan(&category)
	helper.HelperPanic(err)

	return category
}

func (c CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
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

func (c CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	sql := "UPDATE categories SET name = ? WHERE id = ?"

	foundedCategory := c.FindById(ctx, tx, category.Id)
	if foundedCategory.Id == 0 && foundedCategory.Name == "" {
		helper.HelperPanic(errors.New("category not found"))
	}

	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.HelperPanic(err)

	return category
}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	sql := "DELETE FROM categories WHERE id = ?"

	foundedCategory := c.FindById(ctx, tx, categoryId)
	if foundedCategory.Id == 0 && foundedCategory.Name == "" {
		helper.HelperPanic(errors.New("category not found"))
	}

	_, err := tx.ExecContext(ctx, sql, categoryId)
	helper.HelperPanic(err)
}
