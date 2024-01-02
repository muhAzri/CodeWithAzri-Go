package repository

import (
	"CodeWithAzri/internal/app/module/user/entity"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	r := &Repository{db: db}
	return r
}

func (r *Repository) Create(e *entity.User) error {
	return r.db.Create(e).Error
}

func (r *Repository) ReadMany(limit, offset int) ([]entity.User, error) {
	var users []entity.User
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) ReadOne(id string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Update(id string, e *entity.User) error {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}

	user.Name = e.Name

	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	if err := r.db.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
