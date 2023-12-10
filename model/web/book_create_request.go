package web

type BookCreateRequest struct {
	Nama       string `validate:"required" json:"nama"`
	IdPenerbit int32  `validate:"required,min=1" json:"idPenerbit"`
	Kategori   int32  `validate:"required,min=1" json:"kategori"`
	Stok       int32  `validate:"required,min=1" json:"stok"`
	IdPemasok  int32  `validate:"required,min=1" json:"idPemasok"`
}
