package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/zaahidali/task_manager_api/internal/domain/models"
	"github.com/zaahidali/task_manager_api/internal/usecases"
	"github.com/zaahidali/task_manager_api/internal/usecases/mocks"
)

func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepo)
	usecase := usecases.NewTaskUsecase(mockRepo)

	task := models.Task{
		Title:       "Test Task",
		Description: "Testing",
		Status:      "pending",
		DueDate:     time.Now(),
	}

	mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(t models.Task) bool {
		return t.Title == "Test Task" && t.Description == "Testing" && t.Status == "pending"
	})).Return(nil)

	err := usecase.Create(context.Background(), task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
func TestGetTaskByID_Success(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepo)
	usecase := usecases.NewTaskUsecase(mockRepo)

	expectedTask := &models.Task{Title: "Task 1"}
	mockRepo.On("GetByID", mock.Anything, "task123").Return(expectedTask, nil)

	task, err := usecase.GetByID(context.Background(), "task123")

	assert.NoError(t, err)
	assert.Equal(t, expectedTask, task)
	mockRepo.AssertExpectations(t)
}
func TestUpdateTask_Success(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepo)
	usecase := usecases.NewTaskUsecase(mockRepo)

	task := models.Task{Title: "Updated"}
	mockRepo.On("Update", mock.Anything, "task123", task).Return(nil)

	err := usecase.Update(context.Background(), "task123", task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
func TestDeleteTask_Success(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepo)
	usecase := usecases.NewTaskUsecase(mockRepo)

	mockRepo.On("Delete", mock.Anything, "task123").Return(nil)

	err := usecase.Delete(context.Background(), "task123")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
