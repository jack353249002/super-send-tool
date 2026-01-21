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
	SendEmailAccounts      string `json:"send_email_accounts"`
	MessageID              int    `json:"message_id"`
	SmtpIds                string `json:"smtp_ids"`
}
type UserAlivePing struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Position string `json:"position"`
	TimeZone string `json:"time_zone" binding:"required"` //时区
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
	SendEmailAccounts      string `json:"send_email_accounts"`
	MessageID              int    `json:"message_id"`
	SmtpIds                string `json:"smtp_ids"`
}
