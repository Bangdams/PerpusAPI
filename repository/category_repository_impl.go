package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-api-ulang/helper"
	"golang-api-ulang/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	script := "insert into kategori(id,nama) values (?,?)"
	result, err := tx.ExecContext(ctx, script, category.Id, category.Nama)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int32(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	script := "update kategori set nama=? where id=?"
	_, err := tx.ExecContext(ctx, script, category.Nama, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	script := "delete from kategori where id=?"
	_, err := tx.ExecContext(ctx, script, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int32) (domain.Category, error) {
	script := "select id, nama from kategori where id=?"
	rows, err := tx.QueryContext(ctx, script, categoryId)
	helper.PanicIfError(err)

	defer rows.Close()

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Nama)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("CATEGORY NOT FOUND")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	script := "select id, nama from kategori"
	rows, err := tx.QueryContext(ctx, script)

	helper.PanicIfError(err)

	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Nama)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) Pagination(ctx context.Context, tx *sql.Tx, page int32) []domain.Category {
	// get count
	var count int
	tx.QueryRow("select count(*) from kategori").Scan(&count)
	pageSize := 5
	totalPages := count / pageSize
	if count%pageSize != 0 {
		totalPages++
	}

	// check if current page more then total page
	var offset int32
	currentPage := page
	if currentPage > int32(totalPages) {
		offset = 0
		currentPage = 1
	} else {
		offset = (currentPage - 1) * int32(pageSize)
	}

	script := "select id, nama from kategori limit ? offset ?"
	rows, err := tx.QueryContext(ctx, script, pageSize, offset)
	helper.PanicIfError(err)

	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Nama)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
