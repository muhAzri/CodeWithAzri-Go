package handler_test

import (
	"CodeWithAzri/internal/app/module/course/dto"
	"CodeWithAzri/internal/app/module/course/handler"
	"CodeWithAzri/internal/app/module/course/service/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initializeHandler(t *testing.T) (*handler.Handler, *mocks.CourseService) {
	mockService := mocks.NewCourseService(t)
	handler := handler.NewHandler(mockService, validator.New())
	return handler, mockService
}

func TestHandler_GetCourseDetail(t *testing.T) {
	courseHandler, mockService := initializeHandler(t)

	t.Run("Get Course Detail Successfully", func(t *testing.T) {

		defer monkey.UnpatchAll()

		monkey.Patch(chi.URLParam, func(r *http.Request, key string) string {
			return "18a95d2f-a941-4a64-bbe5-256be7626db2"
		})

		mockService.On("GetDetailCourse", mock.AnythingOfType("uuid.UUID")).Return(MockCourseDTO, nil)

		req, err := http.NewRequest("GET", "/courses/18a95d2f-a941-4a64-bbe5-256be7626db2", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		courseHandler.GetCourseDetail(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestHandler_GetCourseDetail_BadRequest(t *testing.T) {
	courseHandler, _ := initializeHandler(t)

	t.Run("Get Course Detail Bad Request", func(t *testing.T) {

		defer monkey.UnpatchAll()

		monkey.Patch(chi.URLParam, func(r *http.Request, key string) string {
			return "invalid_uuid"
		})

		req, err := http.NewRequest("GET", "/courses/invalid_uuid", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		courseHandler.GetCourseDetail(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestHandler_GetCourseDetail_InternalServerError(t *testing.T) {

	courseHandler, mockService := initializeHandler(t)

	t.Run("Get Course Detail Internal Server Error", func(t *testing.T) {
		defer monkey.UnpatchAll()

		monkey.Patch(chi.URLParam, func(r *http.Request, key string) string {
			return "18a95d2f-a941-4a64-bbe5-256be7626db2"
		})

		mockService.On("GetDetailCourse", mock.AnythingOfType("uuid.UUID")).Return(dto.CourseDTO{}, errors.New("Internal Server Error"))

		req, err := http.NewRequest("GET", "/courses/18a95d2f-a941-4a64-bbe5-256be7626db2", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		courseHandler.GetCourseDetail(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		var response map[string]interface{}
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Contains(t, "Internal Server Error", response["meta"].(map[string]interface{})["message"])
		assert.Contains(t, "error", response["meta"].(map[string]interface{})["status"])
	})
}

func TestHandler_GetCourseDetail_NotFound(t *testing.T) {

	courseHandler, mockService := initializeHandler(t)

	t.Run("Get Course Detail Course Not Found", func(t *testing.T) {
		defer monkey.UnpatchAll()

		monkey.Patch(chi.URLParam, func(r *http.Request, key string) string {
			return "18a95d2f-a941-4a64-bbe5-256be7626db2"
		})

		mockService.On("GetDetailCourse", mock.AnythingOfType("uuid.UUID")).Return(dto.CourseDTO{}, nil)

		req, err := http.NewRequest("GET", "/courses/18a95d2f-a941-4a64-bbe5-256be7626db2", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		courseHandler.GetCourseDetail(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)

		var response map[string]interface{}
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Contains(t, response["meta"].(map[string]interface{})["status"], "error")

	})
}

func TestHandler_GetPaginatedCourses(t *testing.T) {
	courseHandler, mockService := initializeHandler(t)

	t.Run("Get Paginated Courses Successfully", func(t *testing.T) {
		mockService.On("GetPaginatedCourses", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(MockArrayCourseDTO, nil)

		req, err := http.NewRequest("GET", "/courses?page=1&limit=10", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		courseHandler.GetPaginatedCourses(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

	})

}

func TestHandler_GetPaginatedCourses_ServiceError(t *testing.T) {
	courseHandler, mockService := initializeHandler(t)

	t.Run("Get Paginated Courses Successfully", func(t *testing.T) {
		mockService.On("GetPaginatedCourses", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(MockArrayCourseDTO, fmt.Errorf("Internal Server Error"))

		req, err := http.NewRequest("GET", "/courses?page=1&limit=10", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		courseHandler.GetPaginatedCourses(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	})

}
