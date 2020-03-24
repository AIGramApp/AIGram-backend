package repository

import (
	"aigram-backend/forms"

	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

// GithubRepository general interface for working with github api
type GithubRepository interface {
	GetUserTokenGithubClient(form forms.GithubAuth) (*github.Client, *string, error)
	GithubClientFromToken(token *oauth2.Token) *github.Client
}
