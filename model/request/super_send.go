package request

type SuperSendListRequest struct {
	PageListRequest
	IsSSL int `json:"is_ssl"`
}
type GetSuperSendInfoRequest struct {
	ID int `json:"id" binding:"required"`
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
type LogOutUserAllDeviceRequest struct {
	ID int `json:"id" binding:"required"`
}
type DeleteSuperSendRequest struct {
	ID int `json:"id" binding:"required"`
}
type LoginSuperSendRequest struct {
	ID int `json:"id" binding:"required"`
}
type RegisterSuperSendRequest struct {
	ID    int    `json:"id" binding:"required"`
	Token string `json:"token" binding:"required"`
}
type SetMessageRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Body        string `json:"body" binding:"required"`
	Attach      string `json:"attach"`
	ContentType string `json:"content_type"`
}
type DelMessageRequest struct {
	ID int `json:"id"`
}
type MessageListRequest struct {
	PageListRequest
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

type SetSendRequest struct {
	Title         string `json:"title" binding:"required"`
	MessageID     int64  `json:"message_id" binding:"required"`
	CreateTime    int64  `json:"create_time"`
	SendModel     uint32 `json:"send_model"`
	SendServerID  string `json:"send_server_id" binding:"required"`
	DispatchModel uint32 `json:"dispatch_model"`
	Receive       string `json:"receive" binding:"required"`
}
type GetSendInfoListRequest struct {
	PageListRequest
	Status int `json:"status"`
}
type GetSendListRequest struct {
	PageListRequest
	Status int `json:"status"`
	SendID int `json:"send_id"`
}
type SendInfoActionRequest struct {
	Status int `json:"status"`
	SendID int `json:"send_id"`
}
type SetSendInfoRequest struct {
	ID int64 `json:"id" binding:"required"`
	/*Title         string `json:"title" binding:"required"`
	MessageID     int64  `json:"message_id" binding:"required"`
	CreateTime    int64  `json:"create_time"`
	SendModel     uint32 `json:"send_model" binding:"required"`*/
	SendServerID string `json:"send_server_id" binding:"required"`
	/*DispatchModel uint32 `json:"dispatch_model" binding:"required"`
	Receive       string `json:"receive" binding:"required"`*/
}
