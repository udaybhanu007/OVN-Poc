package controllers

import (
	"demo/helpers"
	"demo/services"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	userID, err := strconv.ParseInt(request.URL.Query().Get("userId"), 10, 64)
	if err != nil {
		apiErr := &helpers.ApplicationError{
			Message:    "userId must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}
