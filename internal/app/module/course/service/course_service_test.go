package service_test

import (
	"CodeWithAzri/internal/app/module/course/dto"
	"CodeWithAzri/internal/app/module/course/entity"
	"CodeWithAzri/internal/app/module/course/repository/mocks"
	"CodeWithAzri/internal/app/module/course/service"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initializeService(t *testing.T) (service.CourseService, *mocks.CourseRepository) {
	mockRepo := mocks.NewCourseRepository(t)
	service := service.NewCourseService(mockRepo)
	return service, mockRepo
}

func TestService_GetDetailCourse(t *testing.T) {
	courseService, mockRepo := initializeService(t)

	t.Run("Get Detail Course Success", func(t *testing.T) {
		expectedCourse := MockEntity

		mockRepo.On("ReadOne", mock.AnythingOfType("uuid.UUID")).Return(expectedCourse, nil)

		actualCourse, err := courseService.GetDetailCourse(expectedCourse.ID)

		assert.NoError(t, err)
		assert.NotNil(t, actualCourse)
		assert.Equal(t, expectedCourse.ID, actualCourse.ID)
		assert.Equal(t, expectedCourse.Name, actualCourse.Name)
		assert.Equal(t, expectedCourse.Description, actualCourse.Description)
	})

}

func TestService_GetDetailCourseRepositoryError(t *testing.T) {
	courseService, mockRepo := initializeService(t)

	t.Run("Get Detail Course Failed Repository", func(t *testing.T) {
		expectedCourse := MockEntity

		mockRepo.On("ReadOne", mock.AnythingOfType("uuid.UUID")).Return(entity.Course{}, fmt.Errorf("Repository Failure"))

		courseDTO, err := courseService.GetDetailCourse(expectedCourse.ID)

		assert.Error(t, err)
		assert.Equal(t, dto.CourseDTO{}, courseDTO)

	})
}

func TestService_GetDetailCourseAdapterError(t *testing.T) {
	courseService, mockRepo := initializeService(t)

	t.Run("Get Detail Course Failed Adapter", func(t *testing.T) {

		patch := monkey.Patch(json.Marshal, func(v any) ([]byte, error) {
			return nil, errors.New("mocked error during json.Marshal")
		})
		defer patch.Unpatch()

		expectedCourse := MockEntity

		mockRepo.On("ReadOne", mock.AnythingOfType("uuid.UUID")).Return(expectedCourse, nil)

		_, err := courseService.GetDetailCourse(expectedCourse.ID)

		assert.Error(t, err)

	})
}

func TestService_GetPaginatedCourse(t *testing.T) {

	courseService, mockRepo := initializeService(t)

	t.Run("Get Paginated Course Success", func(t *testing.T) {
		expectedCourse := MockArrayEntity

		mockRepo.On("ReadMany", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(expectedCourse, nil)

		actualCourse, err := courseService.GetPaginatedCourses(10, 1)

		assert.NoError(t, err)
		assert.NotNil(t, actualCourse)
		assert.Equal(t, expectedCourse[0].ID, actualCourse[0].ID)

	})
}

func TestService_GetPaginatedCourseRepositoryError(t *testing.T) {
	courseService, mockRepo := initializeService(t)

	t.Run("Get Paginated Course Repository Error", func(t *testing.T) {
		expectedCourse := MockArrayEntity

		mockRepo.On("ReadMany", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(expectedCourse, fmt.Errorf("Repository Failure"))

		_, err := courseService.GetPaginatedCourses(10, 1)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Repository Failure")
	})
}

func TestService_GetPaginatedCourseAdapterError(t *testing.T) {
	courseService, mockRepo := initializeService(t)

	t.Run("Get Paginated Course Repository Error", func(t *testing.T) {
		patch := monkey.Patch(json.Marshal, func(v any) ([]byte, error) {
			return nil, errors.New("mocked error during json.Marshal")
		})
		defer patch.Unpatch()

		expectedCourse := MockArrayEntity

		mockRepo.On("ReadMany", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(expectedCourse, nil)

		_, err := courseService.GetPaginatedCourses(10, 1)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mocked error during json.Marshal")
	})
}
