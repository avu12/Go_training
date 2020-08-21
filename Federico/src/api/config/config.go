package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	//apiGithubAccessToken = "f36172560521502ab348f31b06d1ea4b98435072"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	//return githubAccessToken
	return "f36172560521502ab348f31b06d1ea4b98435072"
}
