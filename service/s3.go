package service

import (
	"aigram-backend/config"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

// S3Service implementation of s3 repo
type S3Service struct {
	config.BaseObject
	session  *session.Session
	uploader *s3manager.Uploader
}

// NewS3Service creates a new instance of s3 service
func NewS3Service(appConfig *config.AppConfiguration) *S3Service {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(appConfig.S3.Region),
		Credentials: credentials.NewStaticCredentials(
			appConfig.S3.AccessKey,
			appConfig.S3.SecretKey,
			"",
		),
	})
	if err != nil {
		panic(fmt.Sprintf("Error happened while connecting to s3 storage %s", err.Error()))
	}
	uploader := s3manager.NewUploader(session)

	return &S3Service{
		BaseObject: config.BaseObject{
			Config: appConfig,
		},
		session:  session,
		uploader: uploader,
	}
}

// Upload uploads a new file and returns the url to the file
func (s3Service *S3Service) Upload(file io.Reader, extension string) (*string, error) {
	filename := fmt.Sprintf("%s.%s", uuid.New().String(), extension)
	_, err := s3Service.uploader.Upload(&s3manager.UploadInput{
		Bucket:       aws.String(s3Service.Config.S3.BucketImages),
		Key:          aws.String(filename),
		Body:         file,
		CacheControl: aws.String("max-age=86400"),
		ACL:          aws.String("public-read"),
		ContentType:  aws.String(fmt.Sprintf("image/%s", extension)),
	})
	if err != nil {
		return nil, err
	}
	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s3Service.Config.S3.BucketImages, s3Service.Config.S3.Region, filename)
	return &fileURL, nil
}
