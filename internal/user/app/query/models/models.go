package models

type UserIdentityDto struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type UserDto struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Personal  string `json:"personal,omitempty"`
	Secret    string `json:"secret,omitempty"`
}
