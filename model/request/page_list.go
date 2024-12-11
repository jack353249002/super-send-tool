package request

type PageListRequest struct {
	KeyWords string `json:"keyWords"`
	PageNo   int    `json:"pageNo" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
}
