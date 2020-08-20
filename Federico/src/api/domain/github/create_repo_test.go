package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang intro2",
		Description: "First repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWWiki:    true,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest
	//unmarchal takes input bytearray and a *pointer that we are trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
}
