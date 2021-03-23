package rest

import (
	"github.com/nicholaslim94/messenger_backend/pkg/account"
	"github.com/nicholaslim94/messenger_backend/pkg/login"
)

//CreateAccountRequest is a dto to create a new Account
type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//ToDomain maps dto to domain model
func (createAccountRequest *CreateAccountRequest) ToDomain() *account.Model {
	return &account.Model{
		Username: createAccountRequest.Username,
		Password: createAccountRequest.Password,
		Email:    createAccountRequest.Email,
	}
}

//CheckUsernameRequest is a dto to check if username exisit
type CheckUsernameRequest struct {
	Username string `json:"username"`
}

//CheckUsernameResponse is a dto stating if username exisit
type CheckUsernameResponse struct {
	Exist bool `json:"exist"`
}

//CheckEmailRequest is a dto to check if email exisit
type CheckEmailRequest struct {
	Email string `json:"email"`
}

//CheckEmailResponse is a dto stating if email exisit
type CheckEmailResponse struct {
	Exist bool `json:"exist"`
}

//LoginRequest is a dto for requesting login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//ToDomain maps dto to domain model
func (loginRequest *LoginRequest) ToDomain() *login.Model {
	return &login.Model{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}
}

//LoginResponse is a dto for responding a token after login
type LoginResponse struct {
	Token string `json:"token"`
}

//NewContactRequest is a dto for requesting a new contact
type NewContactRequest struct {
	Account string `json:"account"`
}

//NewGroupContactRequest is a dto for requesting a new group contact
type NewGroupContactRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Accounts    []string `json:"accounts"`
}
