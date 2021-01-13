package domain

type UserDB interface {
	UserExists(string) bool
}

type UserValidate struct{}

func (r UserValidate) UserExists(email string) bool {
	for _, s := range users {
		if s.Email == email {
			return true
		}
	}
	return false
}
