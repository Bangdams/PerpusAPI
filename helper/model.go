package helper

import (
	"golang-api-ulang/model/domain"
	"golang-api-ulang/model/web"
)

// Response For Book

func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		Id:       book.Id,
		Nama:     book.Nama,
		Penerbit: book.Penerbit,
		Kategori: book.NamaKategori,
		Stok:     book.Stok,
	}
}

func ToBookResponses(books []domain.Book) []web.BookResponse {
	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}

	return bookResponses
}

// Respinse for history Book

func ToBookHistoryResponse(historySupplier domain.HistorySupplier) web.BookHistoryResponse {
	return web.BookHistoryResponse{
		IdBuku:   historySupplier.IdBuku,
		Nama:     historySupplier.NamaBuku,
		Pemasok:  historySupplier.NamaPemasok,
		Kategori: historySupplier.Kategori,
		Stok:     historySupplier.Stok,
		Tanggal:  historySupplier.Tanggal,
		Ket:      historySupplier.Ket,
	}
}

func ToBookHistoryResponses(hisSuppliers []domain.HistorySupplier) []web.BookHistoryResponse {
	var bookHistoryResponses []web.BookHistoryResponse
	for _, bookHistory := range hisSuppliers {
		bookHistoryResponses = append(bookHistoryResponses, ToBookHistoryResponse(bookHistory))
	}

	return bookHistoryResponses
}

// Response For Category

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Nama: category.Nama,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoriesResponses []web.CategoryResponse
	for _, category := range categories {
		categoriesResponses = append(categoriesResponses, ToCategoryResponse(category))
	}

	return categoriesResponses
}

// Response For Supplier

func ToSupplierResponse(supplier domain.Supplier) web.SupplierResponse {
	return web.SupplierResponse{
		Id:   supplier.Id,
		Nama: supplier.Nama,
	}
}

func ToSupplierResponses(suppliers []domain.Supplier) []web.SupplierResponse {
	var supplierResponses []web.SupplierResponse
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, ToSupplierResponse(supplier))
	}

	return supplierResponses
}
