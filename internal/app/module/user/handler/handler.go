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

// Create godoc
//
//	@Summary		Create or fetch a user
//	@Tags			User
//	@Description	Create a new user if not exists or fetch the existing user based on the provided data.
//	@ID				create-or-fetch-user
//	@Accept			json
//	@Produce		json
//	@Param			input			body	dto.CreateUpdateDto	true	"User data for creation or fetching"
//	@Param			Authorization	header	string				true	"With the bearer started"
//	@Security		Bearer
//	@Success		200	{object}	response.Response{data=dto.UserDTO}
//	@Failure		400	{object}	response.ResponseError
//	@Failure		401	{object}	response.ResponseError
//	@Failure		500	{object}	response.ResponseError
//	@Router			/api/v1/users [post]
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

// GetProfile godoc
//
//	@Summary		Fetch user profile
//	@Tags			User
//	@Description	Fetch the profile of the authenticated user based on the provided Authorization token.
//	@ID				get-user-profile
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"Bearer token for authentication"
//	@Security		Bearer
//	@Success		200	{object}	response.Response{data=dto.UserProfileDTO}
//	@Failure		401	{object}	response.ResponseError
//	@Failure		500	{object}	response.ResponseError
//	@Router			/api/v1/users/profile [get]
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ID := requestPkg.GetUserID(r)

	user, err := h.service.GetProfile(ID)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}

	response.BuildResponse(http.StatusOK, "User Profile Fetched Successfully", "Success", user, w)
}
