package request

type ImapListRequest struct {
	PageListRequest
}
type ImapReloadRequest struct {
	Token string `json:"token" binding:""`
}
type SetImapRequest struct {
	ID           int64  `json:"id" binding:""`
	ImapServer   string `json:"imap_server" binding:"required"`
	ImapPassword string `json:"imap_password" binding:"required"`
	Title        string `json:"title" binding:"required"`
	ImapEmail    string `json:"imap_email" binding:"required"`
	Port         int64  `json:"port" binding:"required"`
	MaxClient    int64  `json:"max_client" binding:"required"`
}
type DelImapRequest struct {
	ID int64 `json:"id" binding:"required"`
}
