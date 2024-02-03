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
	Create(dto *dto.CreateUpdateDto) (dto.UserDTO, error)
	GetProfile(ID string) (dto.UserProfileDTO, error)
}

type Service struct {
	repository repository.UserRepository
}

func NewService(r repository.UserRepository) UserService {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) Create(inputDTO *dto.CreateUpdateDto) (dto.UserDTO, error) {
	existingUser, err := s.repository.ReadOne(inputDTO.ID)

	if err != nil && err != sql.ErrNoRows {
		return dto.UserDTO{}, err
	}

	if existingUser.ID != "" {
		existingUserDTO, _ := adapter.AnyToType[dto.UserDTO](existingUser)

		return existingUserDTO, nil
	}

	user, err := adapter.AnyToType[entity.User](inputDTO)

	if err != nil {
		return dto.UserDTO{}, err
	}

	now := timepkg.NowUnixMilli()

	user.CreatedAt = now
	user.UpdatedAt = now

	err = s.repository.Create(user)

	if err != nil {
		return dto.UserDTO{}, err
	}

	userDTO, _ := adapter.AnyToType[dto.UserDTO](user)

	return userDTO, nil
}

func (s *Service) GetProfile(ID string) (dto.UserProfileDTO, error) {
	user, err := s.repository.ReadOne(ID)

	if err != nil && err != sql.ErrNoRows {
		return dto.UserProfileDTO{}, err
	}

	userDTO, err := adapter.AnyToType[dto.UserProfileDTO](user)
	if err != nil {
		return dto.UserProfileDTO{}, err
	}

	return userDTO, nil
}
