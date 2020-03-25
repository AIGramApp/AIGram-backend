package repository

import "io"

// S3Repository common operations with s3 storage
type S3Repository interface {
	Upload(file io.Reader, extension string) (*string, error)
}
