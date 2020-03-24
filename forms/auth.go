package forms

// GithubAuth is coming from the request to auth for github api
type GithubAuth struct {
	Code  string `json:"code" binding:"required"`
	State string `json:"state" binding:"required"`
}
