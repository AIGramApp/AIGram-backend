package service

import (
	"aigram-backend/config"
	"aigram-backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter inits the server
func InitRouter(config *config.AppConfiguration, userController *controllers.UserController, githubController *controllers.GithubController) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     config.CORS.Domains,
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST", "DELETE"}, // Allow all the methods above.
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept", "Content-Length", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := router.Group("api")
	authenticationMiddleware := middleware.AuthenticationRequired(config)
	api.Use(middleware.CSRF())
	{
		user := api.Group("user")
		user.Use(authenticationMiddleware)
		{
			user.GET("", userController.GetUser)
			user.POST("logout", userController.Logout)
		}
		api.POST("/auth", userController.Auth)
	}
	return router
}
