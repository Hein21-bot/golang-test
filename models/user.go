package models

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	OrgName   string `json:"orgnamme"`
	Email     string `json:"email"`
	PhoneNo   string `json:"phoneNo"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipcode"`
	Country   string `json:"country"`
	Password  string `json:"password"`
}
