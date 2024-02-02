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

// GetCourseDetail godoc
//
//	@Summary		Create or fetch course details
//	@Tags			Course
//	@Description	Create a new course if it does not exist or fetch the existing course based on the provided ID.
//	@ID				create-or-fetch-course
//	@Accept			json
//	@Produce		json
//	@Param			id				path	string	true	"Course ID for creation or fetching"
//	@Param			Authorization	header	string	true	"Bearer token for authentication"
//	@Security		Bearer
//	@Success		200	{object}	response.Response{data=dto.CourseDTO}	"Successful response with course details"
//	@Failure		400	{object}	response.ResponseError					"Bad request, invalid input"
//	@Failure		401	{object}	response.ResponseError					"Unauthorized, missing or invalid authentication token"
//	@Failure		404	{object}	response.ResponseError					"Course not found"
//	@Failure		500	{object}	response.ResponseError					"Internal server error"
//	@Router			/api/v1/courses/{id} [get]
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

// GetPaginatedCourses godoc
//	@Summary		Get paginated list of courses
//	@Tags			Course
//	@Description	Retrieve a paginated list of courses based on the specified page and limit parameters.
//	@ID				get-paginated-courses
//	@Accept			json
//	@Produce		json
//	@Param			page			query	int		false	"Page number for pagination (default: 1)"
//	@Param			limit			query	int		false	"Number of items per page (default: 10)"
//	@Param			Authorization	header	string	true	"Bearer token for authentication"
//	@Security		Bearer
//	@Success		200	{object}	response.Response{data=[]dto.CourseDTO}	"Successful response with paginated courses"
//	@Failure		400	{object}	response.ResponseError					"Bad request, invalid input"
//	@Failure		401	{object}	response.ResponseError					"Unauthorized, missing or invalid authentication token"
//	@Failure		500	{object}	response.ResponseError					"Internal server error"
//	@Router			/api/v1/courses [get]
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
