package controllers

import (
	"aigram-backend/config"
	"aigram-backend/repository"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// PostController to add new posts/images
type PostController struct {
	config.BaseObject
	s3Repository repository.S3Repository
}

// NewPostController inits a new post controller
func NewPostController(appConfig *config.AppConfiguration, logger *logrus.Logger, s3Repository repository.S3Repository) *PostController {
	return &PostController{
		BaseObject: config.BaseObject{
			Config: appConfig,
			Logger: logger,
		},
		s3Repository: s3Repository,
	}
}

// UploadImage uploads a new image
func (postController *PostController) UploadImage(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		postController.Logger.Errorf("Image file cannot be opened %s", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	f, err := fileHeader.Open()
	if err != nil {
		postController.Logger.Errorf("Cannot open the image file %s", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	// Get the file extension without the dot
	filename, err := postController.s3Repository.Upload(f, filepath.Ext(fileHeader.Filename)[1:])
	if err != nil {
		postController.Logger.Errorf("Image cannot be uploaded %s", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(200, gin.H{
		"filename": filename,
	})
}
