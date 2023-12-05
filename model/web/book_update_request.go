package web

type BookUpdateRequest struct {
	Id        int32  `validate:"required" json:"id"`
	Nama      string `validate:"required" json:"nama"`
	Penerbit  string `validate:"required" json:"penerbit"`
	Kategori  int32  `validate:"required" json:"kategori"`
	Stok      int32  `json:"stok"`
	IdPemasok int32  `json:"idPemasok"`
}
