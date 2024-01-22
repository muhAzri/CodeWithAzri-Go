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

	"bou.ke/monkey"
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
		patch := monkey.Patch(timepkg.NowUnixMilli, func() int64 { return 12121212 })
		defer patch.Unpatch()

		createUpdateDto := &dto.CreateUpdateDto{
			ID:    "123",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}

		expectedUser := entity.User{
			ID:        createUpdateDto.ID,
			Name:      createUpdateDto.Name,
			Email:     createUpdateDto.Email,
			CreatedAt: 12121212,
			UpdatedAt: 12121212,
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

	// // Test Case 3: Error when creating user

}

func TestService_CreateUserExisted(t *testing.T) {

	userService, mockRepo := initializeService(t)

	// Test Case 2: Existing user, should return existing user without creating a new one
	t.Run("Existing User", func(t *testing.T) {
		existingUser := entity.User{
			ID:        "123",
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: 12121212,
			UpdatedAt: 12121212,
		}

		patch := monkey.Patch(timepkg.NowUnixMilli, func() int64 { return 12121212 })
		defer patch.Unpatch()

		// Set up the mock with the expected argument and return value
		mockRepo.On("ReadOne", existingUser.ID).Return(existingUser, nil)

		createUpdateDto := &dto.CreateUpdateDto{
			ID:    existingUser.ID,
			Name:  "John Doe",             // The name in the DTO should not affect the existing user
			Email: "john.doe@example.com", // The email in the DTO should not affect the existing user
		}

		createdUser, err := userService.Create(createUpdateDto)

		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, &existingUser, &createdUser)
	})
}

func TestService_CreateCheckExistError(t *testing.T) {

	userService, mockRepo := initializeService(t)

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

func TestService_CreateError(t *testing.T) {

	userService, mockRepo := initializeService(t)

	t.Run("Error Creating User", func(t *testing.T) {
		existingUser := entity.User{
			ID:        "456",
			Name:      "Error Case",
			Email:     "error.case@example.com",
			CreatedAt: 12121212,
			UpdatedAt: 12121212,
		}

		patch := monkey.Patch(timepkg.NowUnixMilli, func() int64 { return 12121212 })
		defer patch.Unpatch()

		mockRepo.On("ReadOne", "456").Return(entity.User{}, nil)
		mockRepo.On("Create", existingUser).Return(sql.ErrTxDone)

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
