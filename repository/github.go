package repository

import (
	"aigram-backend/forms"

	"github.com/google/go-github/v29/github"
)

// GithubRepository general interface for working with github api
type GithubRepository interface {
	GetUserTokenGithubClient(form forms.GithubAuth) (*github.Client, *string, error)
}
