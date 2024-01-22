package service_test

import (
	dto "CodeWithAzri/internal/app/module/user/dto"
	"CodeWithAzri/internal/app/module/user/entity"
	"CodeWithAzri/internal/app/module/user/repository/mocks"
	"CodeWithAzri/internal/app/module/user/service"
	timepkg "CodeWithAzri/pkg/timePkg"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initializeService(t *testing.T) (service.UserService, *mocks.UserRepository) {
	mockRepo := mocks.NewUserRepository(t)
	service := service.NewService(mockRepo)
	return service, mockRepo
}

func TestService_Create(t *testing.T) {
	userService, mockRepo := initializeService(t)

	// Test Case 1: Successful creation
	t.Run("Create User Successfully", func(t *testing.T) {
		createUpdateDto := &dto.CreateUpdateDto{
			ID:    "123",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}

		now := timepkg.NowUnixMilli()

		expectedUser := entity.User{
			ID:        createUpdateDto.ID,
			Name:      createUpdateDto.Name,
			Email:     createUpdateDto.Email,
			CreatedAt: now,
			UpdatedAt: now,
		}

		mockRepo.On("ReadOne", createUpdateDto.ID).Return(entity.User{}, sql.ErrNoRows)
		mockRepo.On("Create", expectedUser).Return(nil)

		createdUser, err := userService.Create(createUpdateDto)

		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, createUpdateDto.ID, createdUser.ID)
		assert.Equal(t, createUpdateDto.Name, createdUser.Name)
		assert.Equal(t, createUpdateDto.Email, createdUser.Email)
	})

	// Test Case 2: Existing user, should return existing user without creating a new one
	// t.Run("Existing User", func(t *testing.T) {
	// 	existingUser := entity.User{
	// 		ID:        "123",
	// 		Name:      "Jane Doe",
	// 		Email:     "Jane Doe",
	// 		CreatedAt: 12121212,
	// 		UpdatedAt: 12121212,
	// 	}

	// 	// Set up the mock with the expected argument and return value
	// 	mockRepo.On("ReadOne", existingUser.ID).Return(existingUser, nil)

	// 	createUpdateDto := &dto.CreateUpdateDto{
	// 		ID:    existingUser.ID,
	// 		Name:  "Jane Doe", // The name in the DTO should not affect the existing user
	// 		Email: "Jane Doe", // The email in the DTO should not affect the existing user
	// 	}

	// 	createdUser, err := userService.Create(createUpdateDto)

	// 	assert.NoError(t, err)
	// 	assert.NotNil(t, createdUser)
	// 	assert.Equal(t, &existingUser, &createdUser)
	// })

	// // Test Case 3: Error when creating user
	t.Run("Error Creating User", func(t *testing.T) {
		mockRepo.On("ReadOne", "456").Return(entity.User{}, errors.New("some error"))

		createUpdateDto := &dto.CreateUpdateDto{
			ID:    "456",
			Name:  "Error Case",
			Email: "error.case@example.com",
		}

		createdUser, err := userService.Create(createUpdateDto)

		assert.Error(t, err)
		assert.Equal(t, entity.User{}, createdUser)
	})
}
