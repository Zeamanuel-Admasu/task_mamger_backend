package usecases

import (
	"context"
	"time"

	"github.com/zaahidali/task_manager_api/internal/domain/models"
	"github.com/zaahidali/task_manager_api/internal/domain/ports"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	taskRepo ports.TaskRepository
}

func NewTaskUsecase(repo ports.TaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepo: repo}
}

func (u *TaskUsecase) Create(ctx context.Context, task models.Task) error {
	task.ID = primitive.NewObjectID()
	if task.DueDate.IsZero() {
		task.DueDate = time.Now()
	}
	return u.taskRepo.Create(ctx, task)
}

func (u *TaskUsecase) GetAll(ctx context.Context) ([]models.Task, error) {
	return u.taskRepo.GetAll(ctx)
}

func (u *TaskUsecase) GetByID(ctx context.Context, id string) (*models.Task, error) {
	return u.taskRepo.GetByID(ctx, id)
}

func (u *TaskUsecase) Update(ctx context.Context, id string, task models.Task) error {
	return u.taskRepo.Update(ctx, id, task)
}

func (u *TaskUsecase) Delete(ctx context.Context, id string) error {
	return u.taskRepo.Delete(ctx, id)
}
