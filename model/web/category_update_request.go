package web

type CategoryUpdateRequest struct {
	Id   int32  `validate:"required" json:"id"`
	Nama string `validate:"required" json:"nama"`
}
