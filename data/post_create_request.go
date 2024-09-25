package data

type PostCreateRequest struct {
	Title       string `json:"title" validate:"required,min=10,max=100"`
	Published   bool   `json:"published" validate:"required"`
	Description string `json:"description" validate:"required,min=10,max=1000"`
}
