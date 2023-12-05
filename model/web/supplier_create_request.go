package web

type SupplierCreateRequest struct {
	Nama string `validate:"required" json:"nama"`
}
