package handler

import (
	"CodeWithAzri/internal/app/module/user/dto"
	"CodeWithAzri/internal/app/module/user/service"
	"CodeWithAzri/pkg/requestPkg"
	"CodeWithAzri/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  *service.Service
	validate *validator.Validate
}

func NewHandler(s *service.Service, v *validator.Validate) *Handler {
	h := new(Handler)
	h.service = s
	h.validate = v
	return h
}

func (h *Handler) Create(ctx *gin.Context) {
	var d dto.CreateUpdateDto

	err := ctx.BindJSON(&d)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	d.ID = requestPkg.GetUserID(ctx)

	h.service.Create(&d, ctx)
}
