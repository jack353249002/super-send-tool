package request

type ReceiveRequest struct {
	ID                 int    `json:"id"`                // 主键
	Title              string `json:"title"`             // 标题
	ReceiveRule        string `json:"receive_rule"`      // 接收规则
	ReceiveServerID    string `json:"receive_server_id"` // 接收服务器
	ReceiveModel       int8   `json:"receive_model"`     // 接收模型:0=邮件
	Status             int8   `json:"status"`            // 状态:0=未开始,1=进行中,2=已结束
	EndType            int8   `json:"end_type"`          // 结束类型:0=永久运行,1=自动结束
	EndRule            string `json:"end_rule"`          // 结束规则
	LastNewMessageMark string `json:"last_new_message_mark"`
}
type ReceiveListRequest struct {
	PageListRequest
	Status int8 `json:"status"`
}
type ReceiveMessagesListRequest struct {
	ReceiveID int `json:"receive_id"`
}
type ReceiveMessagesInfoRequest struct {
	ID           int    `json:"id"`
	UID          uint64 `json:"uid"`
	ImapServerId int    `json:"imap_server_id"`
}
type ReceiveMessageActionRequest struct {
	ReceiveID int `json:"receive_id"`
	Status    int `json:"status"`
}
