package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"../services"
	"../utils"
)

func GetUser(c *gin.Context) {

	userId, err := (strconv.ParseInt(c.Param("user_id"), 10, 64))
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User id must be a number!",
			StatusCode: http.StatusNotFound,
			Code:       "bad request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	//return user to the client
	utils.Respond(c, http.StatusOK, user)

}
