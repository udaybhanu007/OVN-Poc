package domain

import (
	"demo/mock"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestAddUserDao(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	MockInterface := mock.NewMockUserDB(controller)
	DBInstance = MockInterface
	user := &User{
		ID:        7,
		FirstName: "acd",
		LastName:  "Soy",
		Email:     "Rabb@email",
	}
	MockInterface.EXPECT().UserExists(user.Email).Return(false)
	if len(AddUser(user)) == 0 {
		t.Errorf("function didnt add the user")
	}
}
