package repository

import "aigram-backend/entities"

// UserRepository for common operations with User
type UserRepository interface {
	FindByID(id int64) *entities.User
	Create(user *entities.User) *entities.User
	GetJWT(user *entities.User) (*string, error)
}
