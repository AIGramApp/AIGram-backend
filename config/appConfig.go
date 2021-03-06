package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// AppConfiguration main config for the app
type AppConfiguration struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Github struct {
		ClientID     string `yaml:"clientId"`
		ClientSecret string `yaml:"clientSecret"`
	} `yaml:"github"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	CORS struct {
		Domains []string `yaml:"domains"`
	} `yaml:"cors"`
	S3 struct {
		AccessKey    string `yaml:"accessKey"`
		SecretKey    string `yaml:"secretKey"`
		Region       string `yaml:"region"`
		BucketImages string `yaml:"bucketImages"`
	} `yaml:"s3"`
}

// Print current configuration
func (config *AppConfiguration) Print(logger *logrus.Logger) {
	logger.Debug(fmt.Sprintf("The configuration loaded: %+v", config))
}

// BaseObject represents base object for controllers and services
type BaseObject struct {
	Config *AppConfiguration
	Logger *logrus.Logger
}
