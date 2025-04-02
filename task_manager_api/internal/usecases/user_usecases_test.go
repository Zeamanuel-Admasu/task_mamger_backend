package usecases_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/zaahidali/task_manager_api/internal/domain/models"
	"github.com/zaahidali/task_manager_api/internal/usecases"
	"github.com/zaahidali/task_manager_api/internal/usecases/mocks"
)

func TestRegister_FirstUserBecomesAdmin(t *testing.T) {
	mockRepo := new(mocks.MockUserRepo)
	usecase := usecases.NewUserUsecase(mockRepo)

	user := models.User{
		Username: "firstuser",
		Password: "password123",
	}

	// Simulate 0 users → first user becomes admin
	mockRepo.On("Count", mock.Anything).Return(int64(0), nil)

	// Expect user with role=admin to be created
	mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(u models.User) bool {
		return u.Role == "admin"
	})).Return(nil)

	err := usecase.Register(context.Background(), user)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
func TestAuthenticateUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepo)
	usecase := usecases.NewUserUsecase(mockRepo)

	expectedUser := &models.User{Username: "testuser", Role: "user"}
	mockRepo.On("AuthenticateUser", mock.Anything, "testuser", "password").Return(expectedUser, nil)

	user, err := usecase.Login(context.Background(), "testuser", "password")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}
func TestPromoteUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepo)
	usecase := usecases.NewUserUsecase(mockRepo)

	mockRepo.On("PromoteUser", mock.Anything, "someuser").Return(nil)

	err := usecase.Promote(context.Background(), "someuser")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
