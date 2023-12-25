package web

type WebResponse struct {
	Code            int         `json:"code"`
	Status          string      `json:"status"`
	Data            interface{} `json:"data"`
	Pagination      int32       `json:"page"`
	TotalPagination int32       `json:"total_page"`
}
