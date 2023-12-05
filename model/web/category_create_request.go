package web

type CategoryCreateRequest struct {
	Nama string `validate:"required" json:"nama"`
}
