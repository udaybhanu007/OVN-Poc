package domain

type UserDB interface {
	userExists(string) bool
}

type UserValidate struct{}

func (r UserValidate) userExists(email string) bool {
	for _, s := range users {
		if s.Email == email {
			return true
		}
	}
	return false
}
