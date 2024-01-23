package handler_test

import (
	"CodeWithAzri/internal/app/module/user/dto"
	"CodeWithAzri/internal/app/module/user/handler"
	"CodeWithAzri/internal/app/module/user/service/mocks"
	jsonpkg "CodeWithAzri/pkg/jsonPkg"
	"CodeWithAzri/pkg/requestPkg"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initializeHandler(t *testing.T) (*handler.Handler, *mocks.UserService) {
	mockService := mocks.NewUserService(t)
	handler := handler.NewHandler(mockService, validator.New())
	return handler, mockService
}

func TestHandler_Create(t *testing.T) {
	userHandler, mockService := initializeHandler(t)

	t.Run("Create User Successfully", func(t *testing.T) {
		userInput := []byte(`{"name": "John Doe", "email": "john.doe@example.com", "profilePicture": "https://example.com/image.png"}`)

		mockService.On("Create", mock.AnythingOfType("*dto.CreateUpdateDto")).Return(dto.UserDTO{ID: "123", Name: "John Doe", Email: "john.doe@example.com"}, nil)

		req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(userInput))
		assert.NoError(t, err)

		patch := monkey.Patch(requestPkg.GetUserID, func(r *http.Request) string {
			return "user123"
		})
		defer patch.Unpatch()

		recorder := httptest.NewRecorder()

		userHandler.Create(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response map[string]interface{}
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "User Created/Fetched Successfully", response["meta"].(map[string]interface{})["message"])
		assert.Equal(t, "Success", response["meta"].(map[string]interface{})["status"])

		mockService.AssertCalled(t, "Create", mock.AnythingOfType("*dto.CreateUpdateDto"))

		mockService.AssertExpectations(t)
	})
}

func TestHandler_Create_DecodeError(t *testing.T) {
	userHandler, _ := initializeHandler(t)

	t.Run("Create User Error Decoding", func(t *testing.T) {
		userInput := []byte(`<invalid json>`)

		patch := monkey.Patch(jsonpkg.Decode, func(r io.Reader, v any) error {
			return errors.New("invalid json")
		})
		patch.Unpatch()

		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userInput))
		fmt.Println(err)

		recorder := httptest.NewRecorder()

		userHandler.Create(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

	})
}

func TestHandler_Create_Validator_Error(t *testing.T) {

	userHandler, _ := initializeHandler(t)

	t.Run("Create User Error Validation", func(t *testing.T) {
		userInput := []byte(`{"name": "John Doe", "email": "john.doe@example.com"}`)

		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userInput))
		fmt.Println(err)

		recorder := httptest.NewRecorder()

		fmt.Println(recorder)

		userHandler.Create(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

	})

}

func TestHandler_Create_ServiceError(t *testing.T) {
	userHandler, mockService := initializeHandler(t)

	t.Run("Create User Service Error", func(t *testing.T) {
		userInput := []byte(`{"name": "John Doe", "email": "john.doe@example.com", "profilePicture": "https://example.com/image.png"}`)

		mockService.On("Create", mock.AnythingOfType("*dto.CreateUpdateDto")).Return(dto.UserDTO{ID: "123", Name: "John Doe", Email: "john.doe@example.com"}, errors.New("Internal Server Error"))

		req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(userInput))
		assert.NoError(t, err)

		patch := monkey.Patch(requestPkg.GetUserID, func(r *http.Request) string {
			return "user123"
		})
		defer patch.Unpatch()

		recorder := httptest.NewRecorder()

		userHandler.Create(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		var response map[string]interface{}
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Error", response["meta"].(map[string]interface{})["message"])
		assert.Equal(t, "error", response["meta"].(map[string]interface{})["status"])

		mockService.AssertCalled(t, "Create", mock.AnythingOfType("*dto.CreateUpdateDto"))

		mockService.AssertExpectations(t)
	})
}

func TestHandler_GetProfile(t *testing.T) {
	userHandler, mockService := initializeHandler(t)

	t.Run("Get User Profile Successfully", func(t *testing.T) {
		mockService.On("GetProfile", mock.AnythingOfType("string")).Return(dto.UserProfileDTO{ID: "123", Name: "John Doe", ProfilePicture: "https://example.com/img.png"}, nil)

		patch := monkey.Patch(requestPkg.GetUserID, func(r *http.Request) string {
			return "user123"
		})
		defer patch.Unpatch()

		req, err := http.NewRequest("GET", "/user/profile", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		userHandler.GetProfile(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestHandler_GetProfile_Service_error(t *testing.T) {
	userHandler, mockService := initializeHandler(t)

	t.Run("Get User Profile Successfully", func(t *testing.T) {
		mockService.On("GetProfile", mock.AnythingOfType("string")).Return(dto.UserProfileDTO{ID: "123", Name: "John Doe", ProfilePicture: "https://example.com/img.png"}, errors.New("Internal Server Error"))

		patch := monkey.Patch(requestPkg.GetUserID, func(r *http.Request) string {
			return "user123"
		})
		defer patch.Unpatch()

		req, err := http.NewRequest("GET", "/user/profile", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		userHandler.GetProfile(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		var response map[string]interface{}
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "Error", response["meta"].(map[string]interface{})["message"])
		assert.Equal(t, "error", response["meta"].(map[string]interface{})["status"])

	})
}
