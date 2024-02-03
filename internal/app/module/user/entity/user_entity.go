package entity

type User struct {
	ID             string `json:"id" gorm:"primaryKey"`
	Name           string `json:"name" gorm:"type:varchar(255)"`
	Email          string `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	ProfilePicture string `json:"profilePicture" gorm:"type:text"`
	CreatedAt      int64  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      int64  `json:"updatedAt" gorm:"autoUpdateTime"`
}
