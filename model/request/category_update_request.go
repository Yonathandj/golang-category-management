package request

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"validate"`
}
