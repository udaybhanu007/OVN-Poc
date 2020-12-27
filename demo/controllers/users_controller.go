package controllers

import (
	"demo/domain"
	"demo/helpers"
	"demo/services"
	"encoding/json"
	"io/ioutil"
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

func AddUser(response http.ResponseWriter, request *http.Request) {
	var user domain.User
	userMap := make(map[int64]*domain.User)
	body, err := ioutil.ReadAll(request.Body)
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
	json.Unmarshal(body, &user)
	userMap, apiErr := services.AddUser(&user)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(userMap)
	response.Write(jsonValue)
}
