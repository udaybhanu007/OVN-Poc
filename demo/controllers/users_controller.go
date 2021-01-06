package controllers

import (
	"demo/auth"
	"demo/domain"
	"demo/helpers"
	"demo/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	if isAuthorized(response, request) == false {
		return
	}
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

func AddUser(response http.ResponseWriter, request *http.Request) error {
	if isAuthorized(response, request) == false {
		return nil
	}
	var user domain.User
	userMap := make(map[int64]*domain.User)
	decoder := json.NewDecoder(request.Body)

	erro := decoder.Decode(&user)
	if erro != nil {
		return helpers.NewHTTPError(erro, "Bad request : invalid JSON.", 400)
	}
	userMap = services.AddUser(&user)
	// Custom error
	if len(userMap) == 0 {
		return helpers.NewHTTPError(nil, "json data unavailable", 400)
	}

	jsonValue, _ := json.Marshal(userMap)
	response.Write(jsonValue)
	return nil
}

func isAuthorized(response http.ResponseWriter, request *http.Request) bool {
	if request.Header.Get("Authorization") == "" ||
		auth.ValidateToken(strings.Split(request.Header.Get("Authorization"), " ")[1]) == false {
		apiErr := &helpers.ApplicationError{
			Message:    "You are not authorized to access this resource.",
			StatusCode: http.StatusUnauthorized,
			Code:       "401",
		}
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return false
	}
	return true
}
