package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUser(t *testing.T) {
	//Initialisation

	//Execution
	user, err := GetUser(0)
	//Validation
	assert.Nil(t, user, "we are not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 not found", err.Message)

}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err, "we are not expecting an error with user id 123")
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "FN", user.FirstName)
	assert.EqualValues(t, "LN", user.LastName)
	assert.EqualValues(t, "myemail@email.com", user.Email)
}
