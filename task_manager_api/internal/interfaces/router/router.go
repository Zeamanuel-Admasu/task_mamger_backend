package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/internal/interfaces/controllers"
	"github.com/zaahidali/task_manager_api/internal/interfaces/middleware"
)

func SetupRouter(taskCtrl *controllers.TaskController, userCtrl *controllers.UserController) *gin.Engine {
	r := gin.Default()

	r.POST("/register", userCtrl.Register)
	r.POST("/login", userCtrl.Login)
	r.POST("/promote", middleware.Authenticate(), middleware.Authorize("admin"), userCtrl.Promote)

	tasks := r.Group("/tasks")
	tasks.Use(middleware.Authenticate())
	{
		tasks.GET("/", taskCtrl.GetTasks)
		tasks.GET("/:id", taskCtrl.GetTaskByID)
		tasks.POST("/", middleware.Authorize("admin"), taskCtrl.CreateTask)
		tasks.PUT("/:id", middleware.Authorize("admin"), taskCtrl.UpdateTask)
		tasks.DELETE("/:id", middleware.Authorize("admin"), taskCtrl.DeleteTask)
	}

	return r
}
