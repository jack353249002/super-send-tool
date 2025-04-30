package request

type AddEtcdBridgeRequest struct {
	Address  string `json:"address" binding:"required"`
	IsSSL    uint8  `json:"is_ssl" binding:""`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UpdateEtcdBridgeRequest struct {
	Address  string `json:"address" binding:"required"`
	IsSSL    uint8  `json:"is_ssl" binding:""`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	ID       int    `json:"id" binding:"required"`
}
type DelEtcdBridgeRequest struct {
	ID int `json:"id"`
}
