package app

import (
	"log"
	"os"
	"time"

	_ "github.com/likesense/task-service/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	postgres "github.com/likesense/task-service/internal/database"
	repository "github.com/likesense/task-service/internal/repositories"
	service "github.com/likesense/task-service/internal/services"
	http "github.com/likesense/task-service/internal/transport"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	gin *gin.Engine
}

func New() (*App, error) {
	a := new(App)
	gin.SetMode(os.Getenv("GIN_MODE"))
	a.gin = gin.Default()
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}
	a.gin.Use(
		cors.New(corsConfig),
		gin.Recovery(),
	)

	a.gin.GET("/task/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	postgresDbConnection := postgres.NewPostgresDbConnection()
	repos := repository.NewRepositories(postgresDbConnection)
	services := service.NewServices(repos)
	handlers := http.NewHandler(services)
	handlers.RegisterAPI(a.gin.Group("/"))

	return a, nil
}

func (a *App) Run() error {
	log.Println("task service launch successfully")
	err := a.gin.Run()
	if err != nil {
		return err
	}
	return nil
}
