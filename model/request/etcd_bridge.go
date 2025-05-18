package request

type EtcdBridgeUsersRequest struct {
	PageListRequest
}
type EtcdBridgeListRequest struct {
	PageListRequest
	IsSSL int `json:"is_ssl"`
}
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
type AddEtcdBridgeUserRequest struct {
	UserName string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Status   int32    `json:"status"`
	Roles    []string `json:"roles"`
}
type UpdateEtcdBridgeUserRequest struct {
	Password string   `json:"password" binding:"required"`
	ID       int      `json:"id" binding:"required"`
	Roles    []string `json:"roles"`
	Status   int32    `json:"status"`
}
type DelEtcdBridgeRequest struct {
	ID int `json:"id"`
}
type GetEtcdBridgeUserRequest struct {
	ID int `json:"id" binding:"required"`
}
type AddEtcdBridgeServerRequest struct {
	CaCerPath string `json:"ca_cer_path" binding:"required"`
	CertPath  string `json:"cert_path" binding:"required"`
	IsSSL     uint8  `json:"is_ssl"`
	KeyPath   string `json:"key_path" binding:"required"`
	Host      string `json:"host" binding:"required"`
	Type      int8   `json:"type"`
	UserName  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
type UpdateEtcdBridgeServerRequest struct {
	ID int `json:"id" binding:"required"`
	AddEtcdBridgeServerRequest
}
type DelEtcdBridgeServerRequest struct {
	ID int `json:"id" binding:"required"`
}
type GetEtcdBridgeServerRequest struct {
	ID int `json:"id" binding:"required"`
}
type AddEtcdBridgeSyncConfRequest struct {
	SourceServerID int64  `json:"source_server_id" binding:"required"`
	TargetServerID int64  `json:"target_server_id" binding:"required"`
	WatcherKey     string `json:"watcher_key" binding:"required"`
	SyncAllTime    int64  `json:"sync_all_time" binding:"required"`
}
type UpdateEtcdBridgeSyncConfRequest struct {
	AddEtcdBridgeSyncConfRequest
	ID int `json:"id" binding:"required"`
}
type DelEtcdBridgeSyncConfRequest struct {
	ID int `json:"id" binding:"required"`
}
type GetEtcdBridgeSyncConfRequest struct {
	ID int `json:"id" binding:"required"`
}
type EtcdBridgeSyncConfActionRequest struct {
	ID     int `json:"id" binding:"required"`
	Status int `json:"status"`
}
