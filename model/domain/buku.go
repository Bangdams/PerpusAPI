package domain

type Book struct {
	Id           int32
	Nama         string
	Penerbit     string
	IdPenerbit   int32
	Kategori     int32
	Stok         int32
	NamaKategori string //* Join from category name
}
