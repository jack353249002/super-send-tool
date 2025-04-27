package request

type SetConfigRequest struct {
	ID  int64  `json:"id"`
	Key string `json:"key" binding:"required"`
	Val string `json:"val" binding:"required"`
}
type GetConfigListRequest struct {
	KeyWords string `json:"keyWords"`
}
type DelConfigRequest struct {
	ID int64 `json:"id" binding:"required"`
}
