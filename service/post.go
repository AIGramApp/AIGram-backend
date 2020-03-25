package service

import (
	"aigram-backend/config"
	"aigram-backend/entities"

	"github.com/jinzhu/gorm"
)

// PostService can publish, view posts
type PostService struct {
	config.BaseObject
	DB *gorm.DB
}

// NewPostService inits a new post service
func NewPostService(appConfig *config.AppConfiguration, db *gorm.DB) *PostService {
	return &PostService{
		BaseObject: config.BaseObject{
			Config: appConfig,
		},
		DB: db,
	}
}

// Publish adds a new post
func (postService *PostService) Publish(post *entities.Post) *entities.Post {
	postService.DB.Create(post)
	return post
}

// Feed returns all the latest posts
func (postService *PostService) Feed() []entities.Post {
	posts := []entities.Post{}
	postService.DB.Preload("User").Order("created_at desc").Find(&posts)
	return posts
}

// PostsByUser returns all the posts by a specific user
func (postService *PostService) PostsByUser(id int64) []entities.Post {
	posts := []entities.Post{}
	postService.DB.Preload("User").Where("user_refer = ?", id).Order("created_at desc").Find(&posts)
	return posts
}
