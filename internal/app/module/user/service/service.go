package service

import (
	"CodeWithAzri/internal/app/module/user/dto"
	"CodeWithAzri/internal/app/module/user/entity"
	"CodeWithAzri/internal/app/module/user/repository"
	"CodeWithAzri/pkg/adapter"
	"CodeWithAzri/pkg/response"
	timepkg "CodeWithAzri/pkg/timePkg"
	"database/sql"
	"net/http"
)

type UserService interface {
	Create(dto *dto.CreateUpdateDto, w http.ResponseWriter, r *http.Request)
}

type Service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) Create(dto *dto.CreateUpdateDto, w http.ResponseWriter, r *http.Request) {
	// Check if the user already exists
	existingUser, err := s.repository.ReadOne(dto.ID)

	if err != nil && err != sql.ErrNoRows {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}

	if existingUser != nil {
		response.Respond(http.StatusOK, response.Meta{
			Message: "User Getted Successfully",
			Code:    http.StatusOK,
			Status:  "success",
		}, existingUser, w)
		return
	}

	user, err := adapter.AnyToType[entity.User](dto)

	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}

	now := timepkg.NowUnixMilli()

	user.CreatedAt = now
	user.UpdatedAt = now

	err = s.repository.Create(&user)

	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}

	response.Respond(http.StatusCreated, response.Meta{
		Message: "User Created Successfully",
		Code:    http.StatusCreated,
		Status:  "success",
	}, user, w)
}
