package controllers

import (
	"aigram-backend/config"
	"aigram-backend/entities"
	"aigram-backend/forms"
	"aigram-backend/middleware"
	"aigram-backend/repository"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v29/github"
	"github.com/sirupsen/logrus"
)

// UserController contains common public operatiosn with the user
type UserController struct {
	config.BaseObject
	userRepository   repository.UserRepository
	githubRepository repository.GithubRepository
	postRepository   repository.PostRepository
}

// NewUserController init a new controller
func NewUserController(appConfig *config.AppConfiguration, logger *logrus.Logger, userRepository repository.UserRepository, githubRepository repository.GithubRepository, postRepository repository.PostRepository) *UserController {
	return &UserController{
		BaseObject: config.BaseObject{
			Config: appConfig,
			Logger: logger,
		},
		userRepository:   userRepository,
		githubRepository: githubRepository,
		postRepository:   postRepository,
	}
}

func currentUser(c *gin.Context) middleware.Claims {
	currentClaim, _ := c.Get("currentUser")
	currentUser := currentClaim.(middleware.Claims)
	return currentUser
}

// GetUser finds the user by id and returns it in json response
func (userController *UserController) GetUser(c *gin.Context) {
	currentUser := userController.userRepository.FindByID(currentUser(c).ID)
	c.JSON(http.StatusOK, currentUser)
}

// Auth method will either register a new or login an existing user
func (userController *UserController) Auth(c *gin.Context) {
	var form forms.GithubAuth
	err := c.BindJSON(&form)
	if err != nil {
		userController.Logger.Errorf("Error happened while trying to bind the form %s", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	var client *github.Client
	client, _, err = userController.githubRepository.GetUserTokenGithubClient(form)
	if err != nil {
		userController.Logger.Errorf("Error happened while creating a new github client %s", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	user, _, _ := client.Users.Get(context.Background(), "")
	// Query the user by id
	// If it doesn`t exist register
	// If it exists, ignore, we will login anyways
	existingUser := userController.userRepository.FindByID(user.GetID())
	if existingUser == nil {
		// Register a new user
		email, _, _ := client.Users.ListEmails(context.Background(), nil)
		existingUser = userController.userRepository.Create(&entities.User{
			ID:       user.GetID(),
			Name:     user.GetName(),
			Avatar:   user.GetAvatarURL(),
			Username: user.GetLogin(),
			Email:    email[0].GetEmail(),
		})
	}
	var token *string
	token, err = userController.userRepository.GetJWT(existingUser)
	if err != nil {
		userController.Logger.Errorf("Error happened while trying to create a token for user %+v %s", existingUser, err.Error())
		c.Status(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Profile will return the full profile info for a user
func (userController *UserController) Profile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		userController.Logger.Errorf("ID is not given %s", err.Error())
		c.Status(http.StatusNotFound)
		return
	}
	user := userController.userRepository.FindByID(int64(id))
	if user == nil {
		userController.Logger.Errorf("User cannot be found %d", id)
		c.Status(http.StatusNotFound)
		return
	}
	posts := userController.postRepository.PostsByUser(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"posts": posts,
	})
}
