package repository

import (
	"context"
	"database/sql"
	"golang-api-ulang/model/domain"
)

type SupplierRepository interface {
	Save(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) domain.Supplier
	Update(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) domain.Supplier
	Delete(ctx context.Context, tx *sql.Tx, supplier domain.Supplier)
	FindById(ctx context.Context, tx *sql.Tx, supplierId int32) (domain.Supplier, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Supplier
	Pagination(ctx context.Context, tx *sql.Tx, page int32) []domain.Supplier
}
