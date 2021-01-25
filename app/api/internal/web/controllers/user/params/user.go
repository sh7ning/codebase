package params

type UserCreateRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type UserListRequest struct {
	Page     int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"page_size" form:"page_size" binding:"required"`
}
