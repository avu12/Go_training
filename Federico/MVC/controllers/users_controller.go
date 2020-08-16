package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../services"
	"../utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User id must be a number!",
			StatusCode: http.StatusNotFound,
			Code:       "bad request",
		}
		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	//return user to the client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
