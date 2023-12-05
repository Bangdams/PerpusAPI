package web

type SupplierUpdateRequest struct {
	Id   int32  `validate:"required" json:"id"`
	Nama string `validate:"required" json:"nama"`
}
