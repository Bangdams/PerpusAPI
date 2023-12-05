package web

type BookCreateRequest struct {
	Nama      string `validate:"required" json:"nama"`
	Penerbit  string `validate:"required" json:"penerbit"`
	Kategori  int32  `validate:"required" json:"kategori"`
	Stok      int32  `validate:"required" json:"stok"`
	IdPemasok int32  `validate:"required" json:"idPemasok"`
}
