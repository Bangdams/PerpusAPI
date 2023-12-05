package service

import (
	"context"
	"golang-api-ulang/model/web"
)

type SupplierService interface {
	Create(ctx context.Context, request web.SupplierCreateRequest) web.SupplierResponse
	Update(ctx context.Context, request web.SupplierUpdateRequest) web.SupplierResponse
	Delete(ctx context.Context, supplierId int32)
	FindById(ctx context.Context, supplierId int32) web.SupplierResponse
	FindAll(ctx context.Context) []web.SupplierResponse
	Pagination(ctx context.Context, page int32) []web.SupplierResponse
}
