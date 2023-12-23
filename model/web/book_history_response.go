package web

type BookHistoryResponse struct {
	IdBuku   int32  `json:"id"`
	Nama     string `json:"nama"`
	Pemasok  string `json:"pemasok"`
	Kategori string `json:"kategori"`
	Stok     int32  `json:"stok"`
	Tanggal  string `json:"tanggal"`
	Ket      string `json:"ket"`
}
