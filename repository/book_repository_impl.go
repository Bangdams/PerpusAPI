package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-api-ulang/helper"
	"golang-api-ulang/model/domain"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	// Save data table buku
	script := "insert into buku(nama, penerbit, kategori, stok) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, book.Nama, book.Penerbit, book.Kategori, book.Stok)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	book.Id = int32(id)

	return book
}

func (repository *BookRepositoryImpl) SaveHisSupp(ctx context.Context, tx *sql.Tx, historyS domain.HistorySupplier) {
	// Save data table history_sup / pemasok
	script := "insert into history_pemasok(id_pemasok, id_buku, stok, tanggal, ket) values(?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, historyS.IdPemasok, historyS.IdBuku, historyS.Stok, historyS.Date, historyS.Ket)
	helper.PanicIfError(err)
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	script := "update buku set nama=?, penerbit=?, kategori=? where id=?"
	_, err := tx.ExecContext(ctx, script, book.Nama, book.Penerbit, book.Kategori, book.Id)
	helper.PanicIfError(err)
	return book
}

func (repository *BookRepositoryImpl) UpdateStok(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	script := "update buku set stok=? where id=?"
	_, err := tx.ExecContext(ctx, script, book.Stok, book.Id)
	helper.PanicIfError(err)

	return book
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book domain.Book) {
	script := "delete from buku where id=?"
	_, err := tx.ExecContext(ctx, script, book.Id)
	helper.PanicIfError(err)
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int32) (domain.Book, error) {
	script := "select buku.id, buku.nama, buku.penerbit, kategori.nama as kategori, stok from buku join kategori on buku.kategori=kategori.id where buku.id=?"
	rows, err := tx.QueryContext(ctx, script, bookId)
	helper.PanicIfError(err)

	defer rows.Close()

	book := domain.Book{}

	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Nama, &book.Penerbit, &book.NamaKategori, &book.Stok)
		helper.PanicIfError(err)

		return book, nil
	} else {
		return book, errors.New("BOOK NOT FOUND")
	}
}

func (repository *BookRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, name string) (domain.Book, error) {
	script := "select id, nama from buku where nama=?"
	rows, err := tx.QueryContext(ctx, script, name)
	helper.PanicIfError(err)

	defer rows.Close()

	book := domain.Book{}

	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Nama)
		helper.PanicIfError(err)

		return book, nil
	} else {
		return book, errors.New("BOOK NOT FOUND")
	}
}

func (repository *BookRepositoryImpl) Pagination(ctx context.Context, tx *sql.Tx, page int32) ([]domain.Book, int32) {
	// get count
	var count int
	tx.QueryRow("select count(*) from buku").Scan(&count)
	pageSize := 10
	totalPages := count / pageSize
	if count%pageSize != 0 {
		totalPages++
	}

	// check if current page more then total page
	var offset int32
	currentPage := page
	if currentPage > int32(totalPages) {
		offset = 0
		currentPage = 1
	} else {
		offset = (currentPage - 1) * int32(pageSize)
	}

	script := "select buku.id, buku.nama, buku.penerbit, kategori.nama as kategori, buku.stok from buku join kategori on buku.kategori=kategori.id order by buku.id limit ? offset ?"
	rows, err := tx.QueryContext(ctx, script, pageSize, offset)
	helper.PanicIfError(err)

	defer rows.Close()

	var books []domain.Book

	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Nama, &book.Penerbit, &book.NamaKategori, &book.Stok)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	return books, currentPage
}
