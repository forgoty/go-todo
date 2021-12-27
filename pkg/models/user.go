package models

type SignedInUser struct {
	UserId      string
	Username    string
	IsAnonymous bool
}

type UserSignInSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
