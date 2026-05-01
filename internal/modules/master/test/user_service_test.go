package test

import (
	"backend-app/internal/modules/auth/models"
	repo "backend-app/internal/modules/master/repository/user"
	"backend-app/internal/modules/master/service/user"
	"backend-app/pkg/pagination"
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

func (m *MockUserRepository) FindAll(ctx context.Context, req pagination.BaseRequest) ([]models.User, int64, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]models.User), args.Get(1).(int64), args.Error(2)
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
	req := pagination.BaseRequest{Page: 1, Paginate: 10}

	t.Run("Success", func(t *testing.T) {
		users := []models.User{
			{Username: "testuser", Email: "test@example.com"},
		}
		mockRepo.On("FindAll", ctx, req).Return(users, int64(1), nil).Once()

		result, meta, err := userService.GetAllUsers(ctx, req)

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "testuser", result[0].Username)
		assert.NotNil(t, meta)
		assert.Equal(t, int64(1), meta.Total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty Data Error", func(t *testing.T) {
		mockRepo.On("FindAll", ctx, req).Return([]models.User{}, int64(0), nil).Once()

		result, meta, err := userService.GetAllUsers(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, meta)
		assert.Equal(t, "Data Not Found", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("FindAll", ctx, req).Return(nil, int64(0), errors.New("db error")).Once()

		result, meta, err := userService.GetAllUsers(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, meta)
		assert.Equal(t, "db error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
