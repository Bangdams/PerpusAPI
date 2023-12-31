package repository

import (
	"context"
	"database/sql"
	"golang-api-ulang/model/domain"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	SaveHisSupp(ctx context.Context, tx *sql.Tx, book domain.HistorySupplier)
	Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	UpdateStok(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	Delete(ctx context.Context, tx *sql.Tx, book domain.Book)
	FindById(ctx context.Context, tx *sql.Tx, bookId int32) (domain.Book, error)
	FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Book, error)
	Pagination(ctx context.Context, tx *sql.Tx, page int32, nameQuery string) ([]domain.Book, int32, int32)
	ReportPagination(ctx context.Context, tx *sql.Tx, page int32, nameQuery string, bookStatus string, startDate string, endDate string) ([]domain.HistorySupplier, int32, int32)
}
