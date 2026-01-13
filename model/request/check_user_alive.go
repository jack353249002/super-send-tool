package request

type CheckUserAliveRequest struct {
	PageListRequest
}
type CreateUserAliveRequest struct {
	Username               string `json:"username" binding:"required"`
	Password               string `json:"password" binding:"required"`
	DayLoginFirstTime      int    `json:"day_login_first_time" binding:""`
	SendID                 int    `json:"send_id"`
	SendEmailActionTimeout int    `json:"send_email_action_timeout"`
	SuperSendConnInfoID    int64  `json:"super_send_conn_info_id"`
}
type UserAlivePing struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type DelUserAliveRequest struct {
	ID int64 `json:"id" binding:"required"`
}
type GetUserAliveInfoRequest struct {
	ID int64 `json:"id" binding:"required"`
}
type SetUserAliveRequest struct {
	ID                     int64  `json:"id" binding:"required"`
	Username               string `json:"username" binding:"required"`
	Password               string `json:"password" binding:""`
	DayLoginFirstTime      int    `json:"day_login_first_time" binding:""`
	SendID                 int    `json:"send_id"`
	SendEmailActionTimeout int    `json:"send_email_action_timeout"`
	SuperSendConnInfoID    int64  `json:"super_send_conn_info_id"`
}
