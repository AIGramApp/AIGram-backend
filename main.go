package main

import (
	"aigram-backend/config"
	"aigram-backend/controllers"
	"aigram-backend/service"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(func() *logrus.Logger {
		return config.InitLogging()
	})
	container.Provide(func(logger *logrus.Logger) *config.AppConfiguration {
		return config.LoadConfig(logger)
	})
	container.Provide(func(config *config.AppConfiguration) *gorm.DB {
		return service.InitializeDatabase(config)
	})
	container.Provide(func(config *config.AppConfiguration, db *gorm.DB) *service.UserService {
		return service.NewUserService(config, db)
	})
	container.Provide(func(config *config.AppConfiguration) *service.S3Service {
		return service.NewS3Service(config)
	})
	container.Provide(func(config *config.AppConfiguration, logger *logrus.Logger, userService *service.UserService, githubService *service.GithubService) *controllers.UserController {
		return controllers.NewUserController(config, logger, userService, githubService)
	})
	container.Provide(func(config *config.AppConfiguration, logger *logrus.Logger, s3Service *service.S3Service) *controllers.PostController {
		return controllers.NewPostController(config, logger, s3Service)
	})
	container.Provide(func(config *config.AppConfiguration) *service.GithubService {
		return service.NewGithubService(config)
	})
	container.Provide(func(config *config.AppConfiguration, userController *controllers.UserController, postController *controllers.PostController) *gin.Engine {
		return service.InitRouter(config, userController, postController)
	})
	return container
}

func main() {
	container := buildContainer()
	container.Invoke(func(router *gin.Engine, config *config.AppConfiguration) {
		router.Run(fmt.Sprintf(":%d", config.Server.Port))
	})
}
