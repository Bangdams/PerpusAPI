package service

import (
	"context"
	"golang-api-ulang/model/web"
)

type BookService interface {
	Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse
	Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse
	Delete(ctx context.Context, bookId int32)
	FindById(ctx context.Context, bookId int32) web.BookResponse
	FindByName(ctx context.Context, name string) web.BookResponse
	Pagination(ctx context.Context, page int32, nameQuery string) ([]web.BookResponse, int32, int32)
	ReportPagination(ctx context.Context, page int32, nameQuery string, bookStatus string, startDate string, endDate string) ([]web.BookHistoryResponse, int32, int32)
}
