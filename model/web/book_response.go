package web

type BookResponse struct {
	Id       int32  `json:"id"`
	Nama     string `json:"nama"`
	Penerbit string `json:"penerbit"`
	Kategori string `json:"kategori"`
	Stok     int32  `json:"stok"`
}
