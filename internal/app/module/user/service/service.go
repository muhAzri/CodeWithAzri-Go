package service

import (
	"CodeWithAzri/internal/app/module/user/dto"
	"CodeWithAzri/internal/app/module/user/entity"
	"CodeWithAzri/internal/app/module/user/repository"
	"CodeWithAzri/pkg/adapter"
	"CodeWithAzri/pkg/response"
	timepkg "CodeWithAzri/pkg/timePkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) Create(dto *dto.CreateUpdateDto, ctx *gin.Context) {
	user, err := adapter.AnyToType[entity.User](dto)

	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}

	now := timepkg.NowUnixMilli()

	user.CreatedAt = now
	user.UpdatedAt = now

	err = s.repository.Create(&user)

	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, ctx)
		return
	}

	response.Respond(http.StatusCreated, response.BuildData(dto), ctx)
}
