package handler

import (
	"CodeWithAzri/internal/app/module/course/service"
	"CodeWithAzri/pkg/requestPkg"
	"CodeWithAzri/pkg/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Handler struct {
	service  service.CourseService
	validate *validator.Validate
}

func NewHandler(s service.CourseService, v *validator.Validate) *Handler {
	h := new(Handler)
	h.service = s
	h.validate = v
	return h
}

func (h *Handler) GetCourseDetail(w http.ResponseWriter, r *http.Request) {
	id := requestPkg.GetURLParam(r, "id")
	courseID, err := uuid.Parse(id)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}

	courseDetail, err := h.service.GetDetailCourse(courseID)
	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}

	if courseDetail.ID == uuid.Nil {
		err = errors.New("course not found")
		response.RespondError(http.StatusNotFound, err, w)
		return
	}

	response.BuildResponse(http.StatusOK, "Course Detail Fetched Successfully", "Success", courseDetail, w)
}

func (h *Handler) GetPaginatedCourses(w http.ResponseWriter, r *http.Request) {
	pageStr := requestPkg.GetQueryParam(r, "page")
	limitStr := requestPkg.GetQueryParam(r, "limit")

	page := 1
	limit := 10

	if pageStr != "" {
		pageInt, err := strconv.Atoi(pageStr)
		if err == nil {
			page = pageInt
		}
	}

	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = limitInt
		}
	}

	courses, err := h.service.GetPaginatedCourses(limit, page)

	if err != nil {
		response.RespondError(http.StatusInternalServerError, err, w)
		return
	}

	response.BuildResponse(http.StatusOK, "Courses Fetched Successfully", "Success", courses, w)
}
