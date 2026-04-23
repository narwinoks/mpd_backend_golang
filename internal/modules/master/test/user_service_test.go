package test

import (
	"backend-app/internal/modules/auth/models"
	repo "backend-app/internal/modules/master/repository/user"
	"backend-app/internal/modules/master/service/user"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll() ([]models.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(id uint) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByUsername(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) FindByNIP(nip string) (bool, error) {
	args := m.Called(nip)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) Create(u *models.User) error {
	args := m.Called(u)
	return args.Error(0)
}

// Ensure MockUserRepository implements UserRepository
var _ repo.UserRepository = (*MockUserRepository)(nil)

func TestUserService_GetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := user.NewUserService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		users := []models.User{
			{Username: "testuser", Email: "test@example.com"},
		}
		mockRepo.On("FindAll").Return(users, nil).Once()

		result, err := userService.GetAllUsers()

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "testuser", result[0].Username)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Data Error", func(t *testing.T) {
		mockRepo.On("FindAll").Return([]models.User{}, nil).Once()

		result, err := userService.GetAllUsers()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "Data Not Found", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("FindAll").Return(nil, errors.New("db error")).Once()

		result, err := userService.GetAllUsers()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
