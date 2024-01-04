package entity

type User struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Email     string `gorm:"uniqueIndex;column:email" json:"email"`
	CreatedAt int64  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at" json:"updatedAt"`
}
