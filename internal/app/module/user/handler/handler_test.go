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
		// Create a sample user input data
		userInput := []byte(`{"name": "John Doe", "email": "john.doe@example.com", "profilePicture": "https://example.com/image.png"}`)

		// Mock the service's Create method
		mockService.On("Create", mock.AnythingOfType("*dto.CreateUpdateDto")).Return(dto.UserDTO{ID: "123", Name: "John Doe", Email: "john.doe@example.com"}, nil)

		// Create a request with the sample user input
		req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(userInput))
		assert.NoError(t, err)

		// Set the user ID in the request context (you may need to update this based on your actual implementation)
		// req = requestPkg.SetUserID(req, "user123")
		patch := monkey.Patch(requestPkg.GetUserID, func(r *http.Request) string {
			return "user123"
		})
		defer patch.Unpatch()

		// Create a response recorder to record the response
		recorder := httptest.NewRecorder()

		// Call the handler's Create method
		userHandler.Create(recorder, req)

		// Assertions
		assert.Equal(t, http.StatusOK, recorder.Code) // Check if the status code is OK

		// Parse the response body to check the response content
		var response map[string]interface{}
		err = json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "User Created/Fetched Successfully", response["meta"].(map[string]interface{})["message"])
		assert.Equal(t, "Success", response["meta"].(map[string]interface{})["status"])

		// Ensure that the user service's Create method was called with the correct input
		mockService.AssertCalled(t, "Create", mock.AnythingOfType("*dto.CreateUpdateDto"))

		// Reset the mock to prepare for the next test
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

		fmt.Println(recorder)

		userHandler.Create(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

	})
}

func TestHandler_Create_Validator_Error(t *testing.T) {

	userHandler, _ := initializeHandler(t)

	t.Run("Create User Error Decoding", func(t *testing.T) {
		userInput := []byte(`{"name": "John Doe", "email": "john.doe@example.com"`)

		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userInput))
		fmt.Println(err)

		recorder := httptest.NewRecorder()

		fmt.Println(recorder)

		userHandler.Create(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

	})

}
