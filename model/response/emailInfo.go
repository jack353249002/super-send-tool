package response

type EmailInfo struct {
	Subject        string   `json:"subject"`
	ImapServerId   int64    `json:"imap_server_id"`
	SenderAddress  string   `json:"sender_address"`
	Date           string   `json:"date"`
	MessageId      string   `json:"message_id"`
	Body           string   `json:"body"`
	FilePaths      []string `json:"file_paths"`
	ImapServerText string   `json:"imap_server_text"`
}
