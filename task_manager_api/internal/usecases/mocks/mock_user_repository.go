package mocks

import (
	"context"

	"github.com/zaahidali/task_manager_api/internal/domain/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(ctx context.Context, user models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepo) AuthenticateUser(ctx context.Context, username, password string) (*models.User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepo) PromoteUser(ctx context.Context, username string) error {
	args := m.Called(ctx, username)
	return args.Error(0)
}

func (m *MockUserRepo) Count(ctx context.Context) (int64, error) {
	args := m.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}
