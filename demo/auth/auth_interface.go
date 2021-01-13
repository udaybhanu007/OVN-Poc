package auth

import (
	"demo/helpers"
	"encoding/json"
	"net/http"
	"strings"
)

type AuthInterface interface {
	IsAuthorized(http.ResponseWriter, *http.Request) bool
}

type UserAuthentication struct{}

func (r UserAuthentication) IsAuthorized(response http.ResponseWriter, request *http.Request) bool {
	if request.Header.Get("Authorization") == "" ||
		ValidateToken(strings.Split(request.Header.Get("Authorization"), " ")[1]) == false {
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
