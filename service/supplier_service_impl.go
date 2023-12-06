package service

import (
	"context"
	"database/sql"
	"golang-api-ulang/exception"
	"golang-api-ulang/helper"
	"golang-api-ulang/model/domain"
	"golang-api-ulang/model/web"
	"golang-api-ulang/repository"

	"github.com/go-playground/validator/v10"
)

type SupplierServiceImpl struct {
	SupplierRepository repository.SupplierRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewSupplierService(supplierRepository repository.SupplierRepository, DB *sql.DB, validate *validator.Validate) SupplierService {
	return &SupplierServiceImpl{
		SupplierRepository: supplierRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *SupplierServiceImpl) Create(ctx context.Context, request web.SupplierCreateRequest) web.SupplierResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	supplier := domain.Supplier{
		Nama: request.Nama,
	}

	supplier = service.SupplierRepository.Save(ctx, tx, supplier)

	return helper.ToSupplierResponse(supplier)
}

func (service *SupplierServiceImpl) Update(ctx context.Context, request web.SupplierUpdateRequest) web.SupplierResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	supplier, err := service.SupplierRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	supplier.Nama = request.Nama

	supplier = service.SupplierRepository.Update(ctx, tx, supplier)

	return helper.ToSupplierResponse(supplier)
}

func (service *SupplierServiceImpl) Delete(ctx context.Context, supplierId int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	supplier, err := service.SupplierRepository.FindById(ctx, tx, supplierId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.SupplierRepository.Delete(ctx, tx, supplier)
}

func (service *SupplierServiceImpl) FindById(ctx context.Context, supplierId int32) web.SupplierResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	supplier, err := service.SupplierRepository.FindById(ctx, tx, supplierId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSupplierResponse(supplier)
}

func (service *SupplierServiceImpl) FindAll(ctx context.Context) []web.SupplierResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	suppliers := service.SupplierRepository.FindAll(ctx, tx)

	return helper.ToSupplierResponses(suppliers)
}

func (service *SupplierServiceImpl) Pagination(ctx context.Context, page int32) ([]web.SupplierResponse, int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	suppliers, currentPage := service.SupplierRepository.Pagination(ctx, tx, page)

	return helper.ToSupplierResponses(suppliers), currentPage
}
