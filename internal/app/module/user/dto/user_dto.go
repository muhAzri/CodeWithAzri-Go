package dto

type CreateUpdateDto struct {
	ID    string
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
