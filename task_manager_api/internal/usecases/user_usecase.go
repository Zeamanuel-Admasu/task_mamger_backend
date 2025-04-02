package usecases

import (
	"context"
	"errors"

	"github.com/zaahidali/task_manager_api/internal/domain/models"
	"github.com/zaahidali/task_manager_api/internal/domain/ports"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo ports.UserRepository
}

func NewUserUsecase(r ports.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

// Register a new user and assign role
func (u *UserUsecase) Register(ctx context.Context, user models.User) error {
	// Assign role based on user count
	count, err := u.repo.Count(ctx)
	if err != nil {
		return err
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	// Hash the password before storing
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPwd)

	return u.repo.CreateUser(ctx, user)
}

// Login verifies credentials
func (u *UserUsecase) Login(ctx context.Context, username, password string) (*models.User, error) {
	return u.repo.AuthenticateUser(ctx, username, password)
}

// Promote upgrades a user to admin
func (u *UserUsecase) Promote(ctx context.Context, username string) error {
	return u.repo.PromoteUser(ctx, username)
}
