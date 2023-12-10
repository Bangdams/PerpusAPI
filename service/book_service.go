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
	FindByName(ctx context.Context, name string, method string) web.BookResponse
	Pagination(ctx context.Context, page int32) ([]web.BookResponse, int32)
}
