package test

import (
	"backend-app/internal/modules/auth/models"
	repo "backend-app/internal/modules/master/repository/user"
	"backend-app/internal/modules/master/service/user"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByUsername(ctx context.Context, username string) (bool, error) {
	args := m.Called(ctx, username)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (bool, error) {
	args := m.Called(ctx, email)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) FindByNIP(ctx context.Context, nip string) (bool, error) {
	args := m.Called(ctx, nip)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) Create(ctx context.Context, u *models.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

// Ensure MockUserRepository implements UserRepository
var _ repo.UserRepository = (*MockUserRepository)(nil)

func TestUserService_GetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := user.NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		users := []models.User{
			{Username: "testuser", Email: "test@example.com"},
		}
		mockRepo.On("FindAll", ctx).Return(users, nil).Once()

		result, err := userService.GetAllUsers(ctx)

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "testuser", result[0].Username)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Data Error", func(t *testing.T) {
		mockRepo.On("FindAll", ctx).Return([]models.User{}, nil).Once()

		result, err := userService.GetAllUsers(ctx)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "Data Not Found", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("FindAll", ctx).Return(nil, errors.New("db error")).Once()

		result, err := userService.GetAllUsers(ctx)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
