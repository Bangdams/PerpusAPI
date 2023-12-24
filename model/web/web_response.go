package web

type WebResponse struct {
	Code            int         `json:"code"`
	Status          string      `json:"status"`
	Data            interface{} `json:"data"`
	Pagination      int32       `json:"pagination"`
	TotalPagination int32       `jason:"total_pagination"`
}
