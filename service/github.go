package service

import (
	"aigram-backend/config"
	"aigram-backend/forms"
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GithubService concrete implementation
type GithubService struct {
	config.BaseObject
}

// NewGithubService initialises a new github service
func NewGithubService(appConfig *config.AppConfiguration) *GithubService {
	return &GithubService{
		BaseObject: config.BaseObject{
			Config: appConfig,
		},
	}
}

// GetUserTokenGithubClient is used to create a github client with currently provided user code using oauth2
func (service *GithubService) GetUserTokenGithubClient(form forms.GithubAuth) (*github.Client, *string, error) {
	config := &oauth2.Config{
		ClientID:     service.Config.Github.ClientID,
		ClientSecret: service.Config.Github.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	setState := oauth2.SetAuthURLParam("state", form.State)
	token, err := config.Exchange(context.Background(), form.Code, setState)
	if err != nil {
		return nil, nil, err
	}
	client := service.GithubClientFromToken(token)
	return client, &token.AccessToken, nil
}

// GithubClientFromToken makes a new github client from the token
func (service *GithubService) GithubClientFromToken(token *oauth2.Token) *github.Client {
	ts := oauth2.StaticTokenSource(
		token,
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)
	return client
}
