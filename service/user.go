package service

import (
	"aigram-backend/config"
	"aigram-backend/entities"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// UserService concrete implementation of the userRepository
type UserService struct {
	config.BaseObject
	DB *gorm.DB
}

// NewUserService initialises a new user service
func NewUserService(appConfig *config.AppConfiguration, db *gorm.DB) *UserService {
	return &UserService{
		BaseObject: config.BaseObject{
			Config: appConfig,
		},
		DB: db,
	}
}

// FindByID finds a user by id
func (userService *UserService) FindByID(githubID int64) *entities.User {
	var user entities.User
	if err := userService.DB.First(&user, githubID).Error; err != nil {
		return nil
	}
	return &user
}

// Create registers a new user and returns it
func (userService *UserService) Create(user *entities.User) *entities.User {
	userService.DB.Create(user)
	return user
}

// GetJWT returns a jwt token to authenticate the current user
func (userService *UserService) GetJWT(user *entities.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// Set 12 hours for token
		"exp": int(time.Now().Add(8 * time.Hour).Unix()),
		"id":  user.ID,
	})
	tokenString, err := token.SignedString([]byte(userService.Config.JWT.Secret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
