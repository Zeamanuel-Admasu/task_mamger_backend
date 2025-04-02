package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	infra "github.com/zaahidali/task_manager_api/internal/infrastructure"
	dbrepo "github.com/zaahidali/task_manager_api/internal/infrastructure/db"
	"github.com/zaahidali/task_manager_api/internal/interfaces/controllers"
	"github.com/zaahidali/task_manager_api/internal/interfaces/router"
	"github.com/zaahidali/task_manager_api/internal/usecases"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found — using environment variables")
	}

	// Read env variables
	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DB")
	port := os.Getenv("PORT")
	if mongoURI == "" || mongoDB == "" || port == "" {
		log.Fatal("Missing required environment variables: MONGO_URI, MONGO_DB, or PORT")
	}

	// DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := infra.ConnectMongo(ctx, mongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	db := client.Database(mongoDB)

	// Repositories
	taskRepo := dbrepo.NewMongoTaskRepository(db.Collection("tasks"))
	userRepo := dbrepo.NewMongoUserRepo(db.Collection("users"))

	// Usecases
	taskUC := usecases.NewTaskUsecase(taskRepo)
	userUC := usecases.NewUserUsecase(userRepo)

	// Controllers
	taskCtrl := controllers.NewTaskController(taskUC)
	userCtrl := controllers.NewUserController(userUC)

	// Router
	r := router.SetupRouter(taskCtrl, userCtrl)
	fmt.Println("🚀 Server running on port", port)
	r.Run(":" + port)
}
