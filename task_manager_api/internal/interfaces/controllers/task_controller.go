package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/internal/domain/models"
	"github.com/zaahidali/task_manager_api/internal/usecases"
)

type TaskController struct {
	usecase *usecases.TaskUsecase
}

func NewTaskController(u *usecases.TaskUsecase) *TaskController {
	return &TaskController{usecase: u}
}

func (ctl *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctl.usecase.Create(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

func (ctl *TaskController) GetTasks(c *gin.Context) {
	tasks, err := ctl.usecase.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (ctl *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := ctl.usecase.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (ctl *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ctl.usecase.Update(context.Background(), id, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func (ctl *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := ctl.usecase.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
