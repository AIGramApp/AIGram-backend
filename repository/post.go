package repository

import "aigram-backend/entities"

// PostRepository interface for common operations with posts
type PostRepository interface {
	Publish(post *entities.Post) *entities.Post
	Feed() []entities.Post
	PostsByUser(id int64) []entities.Post
}
