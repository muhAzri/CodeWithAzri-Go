package entity

type User struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profilePicture"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
}
