package unit

import (
	"backend-app/config"
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/auth/models"
	repo "backend-app/internal/modules/auth/repository/user"
	req "backend-app/internal/modules/auth/request/user"
	"backend-app/internal/modules/auth/service/user"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByIdentity(identity string) (*models.User, error) {
	args := m.Called(identity)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

// Ensure MockUserRepository implements UserRepository
var _ repo.UserRepository = (*MockUserRepository)(nil)

func TestAuthService_Login(t *testing.T) {
	gofakeit.Seed(0)
	mockRepo := new(MockUserRepository)
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret:     gofakeit.LetterN(32),
			Expiration: 24,
		},
	}
	authService := user.NewUserService(mockRepo, cfg)

	t.Run("Happy Path - Success Login", func(t *testing.T) {
		password := gofakeit.Password(true, true, true, true, false, 12)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		
		fakeUser := &models.User{
			Username: gofakeit.Username(),
			Email:    gofakeit.Email(),
			Password: string(hashedPassword),
		}
		fakeUser.ID = gofakeit.Uint32()

		loginReq := &req.LoginRequest{
			Identity: fakeUser.Username,
			Password: password,
		}

		mockRepo.On("FindByIdentity", loginReq.Identity).Return(fakeUser, nil).Once()

		res, err := authService.Login(loginReq)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.AccessToken)
		assert.Equal(t, "Bearer", res.TokenType)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Negative - Identity Not Found", func(t *testing.T) {
		loginReq := &req.LoginRequest{
			Identity: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, 12),
		}

		mockRepo.On("FindByIdentity", loginReq.Identity).Return(nil, errors.New("not found")).Once()

		res, err := authService.Login(loginReq)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.IsType(t, exception.UnauthorizedError{}, err)
		assert.Equal(t, "Invalid credentials", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Negative - Invalid Password", func(t *testing.T) {
		password := gofakeit.Password(true, true, true, true, false, 12)
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		
		fakeUser := &models.User{
			Username: gofakeit.Username(),
			Password: string(hashedPassword),
		}

		loginReq := &req.LoginRequest{
			Identity: fakeUser.Username,
			Password: "wrong-password-" + gofakeit.Word(),
		}

		mockRepo.On("FindByIdentity", loginReq.Identity).Return(fakeUser, nil).Once()

		res, err := authService.Login(loginReq)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.IsType(t, exception.UnauthorizedError{}, err)
		assert.Equal(t, "Invalid credentials", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Negative - Database Failure", func(t *testing.T) {
		loginReq := &req.LoginRequest{
			Identity: gofakeit.Username(),
			Password: gofakeit.Password(true, true, true, true, false, 12),
		}

		mockRepo.On("FindByIdentity", loginReq.Identity).Return(nil, errors.New("connection timeout")).Once()

		res, err := authService.Login(loginReq)

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Chaos - Buffer Stress Identity", func(t *testing.T) {
		loginReq := &req.LoginRequest{
			Identity: gofakeit.LetterN(1000),
			Password: gofakeit.Password(true, true, true, true, false, 12),
		}

		mockRepo.On("FindByIdentity", loginReq.Identity).Return(nil, errors.New("not found")).Once()

		res, err := authService.Login(loginReq)

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Security - SQL Injection Attempt", func(t *testing.T) {
		loginReq := &req.LoginRequest{
			Identity: "' OR 1=1 --",
			Password: gofakeit.Password(true, true, true, true, false, 12),
		}

		mockRepo.On("FindByIdentity", loginReq.Identity).Return(nil, errors.New("not found")).Once()

		res, err := authService.Login(loginReq)

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Security - XSS Attempt", func(t *testing.T) {
		loginReq := &req.LoginRequest{
			Identity: "<script>alert('xss')</script>",
			Password: gofakeit.Password(true, true, true, true, false, 12),
		}

		mockRepo.On("FindByIdentity", loginReq.Identity).Return(nil, errors.New("not found")).Once()

		res, err := authService.Login(loginReq)

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Boundary - Zero/Empty Input", func(t *testing.T) {
		loginReq := &req.LoginRequest{
			Identity: "",
			Password: "",
		}

		mockRepo.On("FindByIdentity", "").Return(nil, errors.New("not found")).Once()

		res, err := authService.Login(loginReq)

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})
}
