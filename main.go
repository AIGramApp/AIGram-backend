package main

import (
	"aigram-backend/config"
	"aigram-backend/controllers"
	"aigram-backend/service"
	"fmt"
	"os"
	"strconv"

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
	container.Provide(func(config *config.AppConfiguration, db *gorm.DB) *service.PostService {
		return service.NewPostService(config, db)
	})
	container.Provide(func(config *config.AppConfiguration, logger *logrus.Logger, userService *service.UserService, githubService *service.GithubService, postService *service.PostService) *controllers.UserController {
		return controllers.NewUserController(config, logger, userService, githubService, postService)
	})
	container.Provide(func(config *config.AppConfiguration, logger *logrus.Logger, s3Service *service.S3Service, postService *service.PostService, userService *service.UserService) *controllers.PostController {
		return controllers.NewPostController(config, logger, s3Service, postService, userService)
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
		port := os.Getenv("PORT")
		if port == "" {
			port = strconv.Itoa(config.Server.Port)
		}
		router.Run(fmt.Sprintf(":%s", port))
	})
}
