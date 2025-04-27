package request

type AddRolesRequest struct {
	Role   string `json:"role" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Action string `json:"action""`
}
type GetRolesRequest struct {
	PageListRequest
}
type DelRoleRequest []Role
type AddRoleForUserRequest struct {
	Role     string `json:"role" binding:"required"`
	Username string `json:"username" binding:"required"`
}
type DelRoleForUserRequest struct {
	Role     string `json:"role" binding:"required"`
	Username string `json:"username" binding:"required"`
}
type Role struct {
	Role   string `json:"role""`
	Path   string `json:"path"`
	Action string `json:"action""`
}
type GetRolesPermissionsRequest struct {
	Role   string `json:"role""`
	Path   string `json:"path"`
	Action string `json:"action""`
}

type DelRolesPermissionsRequest []Role
