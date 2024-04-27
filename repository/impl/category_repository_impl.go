package impl

import (
	"context"
	"database/sql"
	"golang-category-management/model/entity"
	"golang-category-management/model/request"
)

type CategoryRepositoryImpl struct {
}

func (c CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, request request.CategoryCreateRequest) entity.Category {
	sql := "INSERT INTO categories (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, sql, request.Name)

	if err != nil {
		panic(err)
	}
	var id int64
	id, err = result.LastInsertId()

	if err != nil {
		panic(err)
	}

	category := entity.Category{
		Id:   int(id),
		Name: request.Name,
	}
	return category
}

func (c CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) entity.Category {
	sql := "SELECT id, name FROM categories WHERE id= ?"

	category := entity.Category{}
	err := tx.QueryRowContext(ctx, sql, id).Scan(&category)

	if err != nil {
		panic(err)
	}
	return category
}

func (c CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Category {
	sql := "SELECT id, name FROM categories"

	rows, err := tx.QueryContext(ctx, sql)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var categories []entity.Category
	for rows.Next() {
		category := entity.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func (c CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, request request.CategoryUpdateRequest) entity.Category {
	sql := "UPDATE categories SET name = ? WHERE id = ?"

	category := c.FindById(ctx, tx, request.Id)
	if category.Id == 0 && category.Name == "" {
		panic("Category not found")
	}

	_, err := tx.ExecContext(ctx, sql, request.Name, request.Id)
	if err != nil {
		panic(err)
	}

	return category
}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	sql := "DELETE FROM categories WHERE id = ?"

	category := c.FindById(ctx, tx, id)
	if category.Id == 0 && category.Name == "" {
		panic("Category not found")
	}

	_, err := tx.ExecContext(ctx, sql, id)
	if err != nil {
		panic(err)
	}
}
