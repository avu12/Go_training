package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"../../clients/restclient"
	"../../domain/github"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorisation)
	assert.EqualValues(t, "token %s", headerAuthorisationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)

}
func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}
func TestCreateRepoErrorRestClient(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid restclient response", err.Message)
	restclient.StopMockups()

}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid response body!", err.Message)
	restclient.StopMockups()
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1 }")`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json resp body", err.Message)
	restclient.StopMockups()
}

func TestCreateRepoUnathorized(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
	restclient.StopMockups()
}

func TestCreateInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "unmarchaling error create repo", err.Message)
	restclient.StopMockups()
}

func TestCreateNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123,"name": "golang-tutorial","full_name": "federicoleon/golang-tutorial"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.Id)

	restclient.StopMockups()
}
