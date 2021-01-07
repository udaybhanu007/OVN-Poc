package controllers

import (
	"bytes"
	"demo/domain"
	"demo/helpers"
	"demo/mock"
	"errors"
	"strconv"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

var MockAuthInterface mock.MockAuthInterface

func TestAddUser(t *testing.T) {
	user := &domain.User{
		ID:        7,
		FirstName: "acd",
		LastName:  "Soy",
		Email:     "app@email",
	}
	response, err := makeRequestToAddUser(*user, t)
	if err != nil {
		t.Fatal(err)
	}
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
			input:  domain.User{ID: 7, FirstName: "acd", LastName: "Soy", Email: "Apple@email"},
			expect: 200,
		},
		{
			input:  domain.User{ID: 7, FirstName: "", LastName: "Soy", Email: "carrot@email"},
			expect: 400,
		},
		{
			input:  domain.User{ID: 7, FirstName: "acd", LastName: "Soy", Email: ""},
			expect: 400,
		},
	}
	for _, s := range testCases {
		user := domain.User(s.input)
		response, err := makeRequestToAddUser(user, t)
		if err != nil {
			t.Fatal(err)
		}
		if response.Result().StatusCode != s.expect {
			t.Errorf("expected status %v but got %v", 200, response.Result().StatusCode)
		}
	}
}

func BenchmarkAddUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := &domain.User{
			ID:        i + 7,
			FirstName: "acd",
			LastName:  "Soy",
			Email:     "Rabb" + strconv.Itoa(i) + "@email",
		}
		response, err := makeRequestToAddUser(*user, b)
		if err != nil {
			b.Fatal(err)
		}
		if response.Result().StatusCode != 200 {
			b.Errorf("expected status %v but got %v", 200, response.Result().StatusCode)
		}
	}
}

func makeRequestToAddUser(user domain.User, t testing.TB) (response *httptest.ResponseRecorder, APIerror error) {
	// mocking for Authentication
	controller := gomock.NewController(t)
	defer controller.Finish()
	MockAuthInterface := mock.NewMockAuthInterface(controller)
	AuthInstance = MockAuthInterface

	jsonUser, _ := json.Marshal(user)
	request, err := http.NewRequest("POST", "/adduser", bytes.NewBuffer(jsonUser))
	if err != nil {
		return nil, errors.New("fatal server error")
	}
	response = httptest.NewRecorder()
	MockAuthInterface.EXPECT().IsAuthorized(response, request).Return(true)
	helpers.RootHandler(AddUser).ServeHTTP(response, request)
	return response, nil
}
