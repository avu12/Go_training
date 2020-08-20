package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../../clients/restclient"
	"../../domain/github"
)

const (
	headerAuthorisation       = "Authorization"
	headerAuthorisationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorisationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorisation, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body!",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		err := json.Unmarshal(bytes, &errResponse)
		if err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json resp body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying unmarchal create repo succesful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "unmarchaling error create repo",
		}
	}
	return &result, nil
}
