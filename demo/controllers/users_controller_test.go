package controllers_test

import (
	"bytes"
	"demo/controllers"
	"demo/domain"
	"demo/helpers"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddUser(t *testing.T) {
	user := &domain.User{
		ID:        7,
		FirstName: "acd",
		LastName:  "Soy",
		Email:     "Rabb@email",
	}
	jsonUser, _ := json.Marshal(user)
	request, err := http.NewRequest("POST", "/adduser", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	helpers.RootHandler(controllers.AddUser).ServeHTTP(response, request)
	if response.Result().StatusCode != 200 {
		t.Errorf("expected status %v but got %v", 200, response.Result().StatusCode)
	}
}

func TestAddUserTableDrivenTest(t *testing.T) {
	testCases := []struct {
		input  domain.User
		expect int
	}{
		{
			input:  domain.User{ID: 7, FirstName: "acd", LastName: "Soy", Email: "Rabb@email"},
			expect: 200,
		},
		{
			input:  domain.User{ID: 7, FirstName: "", LastName: "Soy", Email: "Rabb@email"},
			expect: 400,
		},
		{
			input:  domain.User{ID: 7, FirstName: "acd", LastName: "Soy", Email: ""},
			expect: 400,
		},
	}
	for _, s := range testCases {
		jsonUser, _ := json.Marshal(s.input)
		request, err := http.NewRequest("POST", "/adduser", bytes.NewBuffer(jsonUser))
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()
		helpers.RootHandler(controllers.AddUser).ServeHTTP(response, request)
		if response.Result().StatusCode != s.expect {
			t.Errorf("expected status %v but got %v", 200, response.Result().StatusCode)
		}
	}
}
