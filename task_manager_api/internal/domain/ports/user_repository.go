package ports

import (
	"context"

	"github.com/zaahidali/task_manager_api/internal/domain/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	AuthenticateUser(ctx context.Context, username, password string) (*models.User, error)
	PromoteUser(ctx context.Context, username string) error
	Count(ctx context.Context) (int64, error) // ✅ Add this!
}
