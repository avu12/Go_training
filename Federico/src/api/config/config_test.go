package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", apiGithubAccessToken)
}

func TestGetGithubAccessToken(t *testing.T) {
	assert.EqualValues(t, "f36172560521502ab348f31b06d1ea4b98435072", GetGithubAccessToken())
}
