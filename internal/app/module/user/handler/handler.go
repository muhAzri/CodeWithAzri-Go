package handler

import (
	"CodeWithAzri/internal/app/module/user/dto"
	"CodeWithAzri/internal/app/module/user/service"
	jsonpkg "CodeWithAzri/pkg/jsonPkg"
	"CodeWithAzri/pkg/requestPkg"
	"CodeWithAzri/pkg/response"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  service.UserService
	validate *validator.Validate
}

func NewHandler(s service.UserService, v *validator.Validate) *Handler {
	h := new(Handler)
	h.service = s
	h.validate = v
	return h
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var d dto.CreateUpdateDto

	err := jsonpkg.Decode(r.Body, &d)

	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}

	if err := h.validate.Struct(d); err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}

	d.ID = requestPkg.GetUserID(r)

	user, err := h.service.Create(&d)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}

	response.BuildResponse(http.StatusOK, "User Created/Fetched Successfully", "Success", user, w)
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ID := requestPkg.GetUserID(r)

	user, err := h.service.GetProfile(ID)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}

	response.BuildResponse(http.StatusOK, "User Profile Fetched Successfully", "Success", user, w)
}
