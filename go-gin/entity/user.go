package entity

type User struct {
	ID          int     `json:"id"`
	FirstName   string  `json:"first-name"`
	LastName    string  `json:"last-name"`
	Age         int     `json:"age"`
	UserAddress Address `json:"user-address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Zipcode int    `json:"zipcode"`
}
