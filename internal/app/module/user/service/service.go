package service

import (
	"CodeWithAzri/internal/app/module/user/dto"
	"CodeWithAzri/internal/app/module/user/entity"
	"CodeWithAzri/internal/app/module/user/repository"
	"CodeWithAzri/pkg/adapter"
	timepkg "CodeWithAzri/pkg/timePkg"
	"database/sql"
)

type UserService interface {
	Create(dto *dto.CreateUpdateDto) (entity.User, error)
}

type Service struct {
	repository repository.UserRepository
}

func NewService(r repository.UserRepository) UserService {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) Create(dto *dto.CreateUpdateDto) (entity.User, error) {
	existingUser, err := s.repository.ReadOne(dto.ID)

	if err != nil && err != sql.ErrNoRows {
		return entity.User{}, err
	}

	if existingUser != nil {

		return *existingUser, nil
	}

	user, err := adapter.AnyToType[entity.User](dto)

	if err != nil {
		return entity.User{}, err

	}

	now := timepkg.NowUnixMilli()

	user.CreatedAt = now
	user.UpdatedAt = now

	err = s.repository.Create(&user)

	if err != nil {

		return entity.User{}, err
	}

	return user, nil
}
