package request

type UsersRequest struct {
	PageListRequest
}
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type DelUserRequest struct {
	ID int64 `json:"id" binding:"required"`
}
type GetUserInfoRequest struct {
	ID int64 `json:"id" binding:"required"`
}
type SetPasswordRequest struct {
	Password string `json:"password" binding:"required"`
}
type SetUserPasswordRequest struct {
	ID       int64  `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
