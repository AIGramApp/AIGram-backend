package main

import (
	"aigram-backend/config"
	"aigram-backend/controllers"
	"aigram-backend/server"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(func() *config.AppConfiguration {
		return config.Init()
	})
	container.Provide(func(config *config.AppConfiguration) *controllers.EarlyAccess {
		return controllers.NewEarlyAccess(config)
	})
	container.Provide(func(config *config.AppConfiguration, earlyAccess *controllers.EarlyAccess) *gin.Engine {
		return server.Init(config, earlyAccess)
	})
	return container
}

func main() {
	container := buildContainer()
	container.Invoke(func(server *gin.Engine, config *config.AppConfiguration) {
		port := os.Getenv("PORT")
		if port == "" {
			port = strconv.Itoa(config.Server.Port)
		}
		server.Run(fmt.Sprintf(":%s", port))
	})
}
