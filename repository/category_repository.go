package repository

import (
	"context"
	"database/sql"
	"golang-api-ulang/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int32) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Category, error)
	Pagination(ctx context.Context, tx *sql.Tx, page int32) ([]domain.Category, int32)
}
