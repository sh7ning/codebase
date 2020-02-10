package params

type UserCreateRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}
