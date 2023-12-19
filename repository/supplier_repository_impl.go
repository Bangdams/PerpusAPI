package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-api-ulang/helper"
	"golang-api-ulang/model/domain"
)

type SupplierRepositoryImpl struct {
}

func NewSupplierRepository() SupplierRepository {
	return &SupplierRepositoryImpl{}
}

func (repository *SupplierRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) domain.Supplier {
	script := "insert into pemasok(nama) values(?)"
	result, err := tx.ExecContext(ctx, script, supplier.Nama)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	supplier.Id = int32(id)

	return supplier
}

func (repository *SupplierRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) domain.Supplier {
	script := "update pemasok set nama=? where id=?"
	_, err := tx.ExecContext(ctx, script, supplier.Nama, supplier.Id)
	helper.PanicIfError(err)

	return supplier
}

func (repository *SupplierRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, supplier domain.Supplier) {
	script := "delete from pemasok where id=?"
	_, err := tx.ExecContext(ctx, script, supplier.Id)
	helper.PanicIfError(err)
}

func (repository *SupplierRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, supplierId int32) (domain.Supplier, error) {
	script := "select id, nama from pemasok where id=?"
	rows, err := tx.QueryContext(ctx, script, supplierId)
	helper.PanicIfError(err)

	defer rows.Close()

	supplier := domain.Supplier{}

	if rows.Next() {
		err := rows.Scan(&supplier.Id, &supplier.Nama)
		helper.PanicIfError(err)

		return supplier, nil
	} else {
		return supplier, errors.New("SUPPLIER NOT FOUND")
	}
}

func (repository *SupplierRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Supplier, error) {
	script := "select nama from pemasok where nama=?"
	rows, err := tx.QueryContext(ctx, script, name)
	helper.PanicIfError(err)

	defer rows.Close()

	supplier := domain.Supplier{}

	if rows.Next() {
		err := rows.Scan(&supplier.Nama)
		helper.PanicIfError(err)

		return supplier, nil
	} else {
		return supplier, errors.New("SUPPLIER NOT FOUND")
	}
}

func (repository *SupplierRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Supplier {
	script := "select id, nama from pemasok"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)

	defer rows.Close()

	var suppliers []domain.Supplier

	for rows.Next() {
		supplier := domain.Supplier{}
		err := rows.Scan(&supplier.Id, &supplier.Nama)
		helper.PanicIfError(err)

		suppliers = append(suppliers, supplier)
	}

	return suppliers
}

func (repository *SupplierRepositoryImpl) Pagination(ctx context.Context, tx *sql.Tx, page int32, nameQuery string) ([]domain.Supplier, int32) {
	// * Global Variable
	var count int
	var name string
	var script string
	var rows *sql.Rows
	var err error

	// * Get Name Book By Query Parameter
	getName := nameQuery
	name = "%" + getName + "%"

	// * Get Count
	if name != "%%" {
		script = "select count(*) from pemasok where nama like ?"
		tx.QueryRow(script, name).Scan(&count)
	} else {
		tx.QueryRow("select count(*) from pemasok").Scan(&count)
	}

	pageSize := 5
	totalPages := count / pageSize
	if count%pageSize != 0 {
		totalPages++
	}

	// * check if current page more then total page
	var offset int32
	currentPage := page
	if currentPage > int32(totalPages) {
		offset = 0
		currentPage = 1
	} else {
		offset = (currentPage - 1) * int32(pageSize)
	}

	// * Check if var name on nameQuery is %% or can call null
	if name != "%%" {
		script = "select id, nama from pemasok where nama like ? limit ? offset ?"
		rows, err = tx.QueryContext(ctx, script, name, pageSize, offset)
	} else {
		script := "select id, nama from pemasok limit ? offset ?"
		rows, err = tx.QueryContext(ctx, script, pageSize, offset)
	}

	helper.PanicIfError(err)

	defer rows.Close()

	var suppliers []domain.Supplier

	for rows.Next() {
		supplier := domain.Supplier{}
		err := rows.Scan(&supplier.Id, &supplier.Nama)
		helper.PanicIfError(err)

		suppliers = append(suppliers, supplier)
	}

	return suppliers, currentPage
}
