package handler

import (
	"CodeWithAzri/internal/app/module/user/dto"
	"CodeWithAzri/internal/app/module/user/service"
	jsonpkg "CodeWithAzri/pkg/jsonPkg"
	"CodeWithAzri/pkg/requestPkg"
	"CodeWithAzri/pkg/response"
	"fmt"
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

func (h *Handler) parseValidateRequestBody(ctx *gin.Context) (dto.CreateUpdateDto, error) {
	var d dto.CreateUpdateDto

	err := jsonpkg.Decode(ctx.Request.Body, &d)
	if err != nil {
		return d, err
	}
	// validate request body
	err = h.validate.Struct(d)
	if err != nil {
		// Struct is invalid
		// for checking only
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	}
	return d, err
}

func (h *Handler) Create(ctx *gin.Context) {
	d, err := h.parseValidateRequestBody(ctx)
	d.ID = requestPkg.GetUserID(ctx)

	if err != nil {
		response.RespondError(http.StatusBadRequest, err, ctx)
		return
	}
	h.service.Create(&d, ctx)
}
