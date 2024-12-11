package request

type SuperSendListRequest struct {
	PageListRequest
	IsSSL int `json:"is_ssl"`
}
type AddSuperSendRequest struct {
	Address  string `json:"address" binding:"required"`
	IsSSL    uint8  `json:"is_ssl" binding:""`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UpdateSuperSendRequest struct {
	ID       int    `json:"id" binding:"required"`
	Address  string `json:"address" binding:"required"`
	IsSSL    uint8  `json:"is_ssl" binding:""`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type SetSuperSendOnlineRequest struct {
	ID     int   `json:"id" binding:"required"`
	Online uint8 `json:"online"`
}
type DeleteSuperSendRequest struct {
	ID int `json:"id" binding:"required"`
}
type LoginSuperSendRequest struct {
	ID int `json:"id" binding:"required"`
}
