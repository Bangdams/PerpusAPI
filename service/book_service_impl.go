package service

import (
	"context"
	"database/sql"
	"golang-api-ulang/exception"
	"golang-api-ulang/helper"
	"golang-api-ulang/model/domain"
	"golang-api-ulang/model/web"
	"golang-api-ulang/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewBookService(bookRepository repository.BookRepository, DB *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book := domain.Book{}

	historyS := domain.HistorySupplier{}

	//* Check if name is EXISTS in database
	_, err = service.BookRepository.FindByName(ctx, tx, request.Nama)
	if err != nil {
		book.Nama = request.Nama
		book.IdPenerbit = request.IdPenerbit
		book.Kategori = request.Kategori
		book.Stok = request.Stok

		book = service.BookRepository.Save(ctx, tx, book)

		historyS.IdPemasok = request.IdPemasok
		historyS.Stok = book.Stok
		historyS.IdBuku = book.Id
		historyS.Date = time.Now().String()
		historyS.Ket = "Buku Baru"

		service.BookRepository.SaveHisSupp(ctx, tx, historyS)
	} else {
		panic(exception.NewDuplicateName("DATA IS EXISTS")) // todo Omeken deui dam cara jalan nu lain tong pake ieu
	}

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, request.Id)

	historyS := domain.HistorySupplier{}

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Check request stock

	// * Coba dam Cek deui dam ieu asa salah cara jalan anu lebih bagus deui jadi di refaktoring ken
	if request.Stok != 0 {
		if request.Stok > 0 {
			book.Stok += request.Stok
			book = service.BookRepository.UpdateStok(ctx, tx, book)

			historyS.IdPemasok = request.IdPemasok
			historyS.IdBuku = book.Id
			historyS.Stok = request.Stok
			historyS.Date = time.Now().String()
			historyS.Ket = "Tambah Stok"

			service.BookRepository.SaveHisSupp(ctx, tx, historyS)
		} else {
			panic(exception.NewNotFoundError(err.Error())) // ! ganti buat deui tong pake newnotfound
		}
	} else if request.Stok == 0 && request.Nama != "" {
		err := service.Validate.Struct(request)
		helper.PanicIfError(err)

		//* Check if name is EXISTS in database
		body, err := service.BookRepository.FindByName(ctx, tx, request.Nama)

		if err != nil || book.Nama == body.Nama {
			book.Nama = request.Nama
			book.IdPenerbit = request.IdPenerbit
			book.Kategori = request.Kategori

			book = service.BookRepository.Update(ctx, tx, book)
		} else {
			panic(exception.NewDuplicateName("DATA IS EXISTS"))
		}
	} else {
		err := service.Validate.Struct(request)
		helper.PanicIfError(err)
	}

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BookRepository.Delete(ctx, tx, book)
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int32) web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) FindByName(ctx context.Context, name string) web.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	book, err := service.BookRepository.FindByName(ctx, tx, name)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBookResponse(book)
}
func (service *BookServiceImpl) Pagination(ctx context.Context, page int32) ([]web.BookResponse, int32) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	books, currentPage := service.BookRepository.Pagination(ctx, tx, page)

	return helper.ToBookResponses(books), currentPage
}
