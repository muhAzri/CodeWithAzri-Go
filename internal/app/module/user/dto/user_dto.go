package dto

type UserDTO struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	ProfilePicture string `json:"profilePicture,omitempty"`
	CreatedAt      int64  `json:"createdAt,omitempty"`
	UpdatedAt      int64  `json:"updatedAt,omitempty"`
}

type UserProfileDTO struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	ProfilePicture string `json:"profilePicture,omitempty"`
}

type CreateUpdateDto struct {
	ID             string `json:"id"`
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	ProfilePicture string `json:"profilePicture" validate:"required"`
}
