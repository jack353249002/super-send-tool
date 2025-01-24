package request

type SmtpListRequest struct {
	PageListRequest
}
type ReloadRequest struct {
	Token string `json:"token" binding:""`
}
type SetSmtpRequest struct {
	ID             int64  `json:"id" binding:""`
	SmtpServer     string `json:"smtp_server" binding:"required"`
	Title          string `json:"title" binding:"required"`
	SmtpSendEmail  string `json:"smtp_send_email" binding:"required"`
	Port           int64  `json:"port" binding:"required"`
	MaxConcurrency int64  `json:"max_concurrency" binding:"required"`
	SmtpPassword   string `json:"smtp_password" binding:"required"`
}
type DelSmtpRequest struct {
	ID int64 `json:"id" binding:"required"`
}
