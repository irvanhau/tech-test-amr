package handler

type InputRequest struct {
	Name        string `json:"name" form:"name" validate:"required,max=250"`
	Description string `json:"description" form:"description" validate:"required"`
	CategoryID  string `json:"category_id" form:"category_id" validate:"required"`
}
