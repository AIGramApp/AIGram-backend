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
	container.Provide(func(config *config.AppConfiguration, userService *service.UserService, logger *logrus.Logger, githubService *service.GithubService) *controllers.UserController {
		return controllers.NewUserController(config, logger, userService, githubService)
	})
	container.Provide(func(config *config.AppConfiguration) *service.GithubService {
		return service.NewGithubService(config)
	})
	container.Provide(func(config *config.AppConfiguration, userController *controllers.UserController) *gin.Engine {
		return service.InitRouter(config, userController)
	})
	return container
}

func main() {
	container := buildContainer()
	container.Invoke(func(router *gin.Engine, config *config.AppConfiguration) {
		router.Run(fmt.Sprintf(":%d", config.Server.Port))
	})
}
