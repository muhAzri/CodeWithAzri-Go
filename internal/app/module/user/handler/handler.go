package handler

import (
	"CodeWithAzri/internal/app/module/user/dto"
	"CodeWithAzri/internal/app/module/user/service"
	"CodeWithAzri/pkg/requestPkg"
	"CodeWithAzri/pkg/response"
	"encoding/json"
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

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	defer r.Body.Close()

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
