package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nicholaslim94/messenger_backend/pkg/account"
	"github.com/nicholaslim94/messenger_backend/pkg/contact"
	"github.com/nicholaslim94/messenger_backend/pkg/group"
	"github.com/nicholaslim94/messenger_backend/pkg/groupAccount"
	"github.com/nicholaslim94/messenger_backend/pkg/login"
)

//CreateAccountHandler creates a new account. Account should have a unique login id and email.
//Upon creation, returns status 201 else return status 400
func CreateAccountHandler(s account.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var createAccountRequest CreateAccountRequest
		err := json.NewDecoder(r.Body).Decode(&createAccountRequest)
		if err != nil {
			log.Println(err)
			return
		}
		err = s.CreateAccount(createAccountRequest.ToDomain())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok"))
	}
}

//CheckUsernameHandlerFunc checks if the given username is available
func CheckUsernameHandlerFunc(s account.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var checkUsernameRequest CheckUsernameRequest
		err := json.NewDecoder(r.Body).Decode(&checkUsernameRequest)
		if err != nil {
			log.Println(err)
			return
		}
		exist, err := s.UsernameExist(checkUsernameRequest.Username)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}
		checkUsernameResponse := CheckUsernameResponse{
			Exist: exist,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(checkUsernameResponse)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}
	}
}

//CheckEmailHandlerFunc checks if the given email is available
func CheckEmailHandlerFunc(s account.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var checkEmailRequest CheckEmailRequest
		err := json.NewDecoder(r.Body).Decode(&checkEmailRequest)
		if err != nil {
			log.Println(err)
			return
		}
		exist, err := s.EmailExist(checkEmailRequest.Email)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}
		CheckEmailResponse := CheckEmailResponse{
			Exist: exist,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CheckEmailResponse)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}
	}
}

//LoginHandleFunc returns a JWT token if credentials are valid
func LoginHandleFunc(s login.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginRequest LoginRequest
		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token, err := s.Login(loginRequest.ToDomain())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		loginDto := &LoginResponse{Token: token}
		json.NewEncoder(w).Encode(loginDto)
	}
}

//NewContactHandleFunc creates a new contact for both requestor and requested user
func NewContactHandleFunc(as account.Service, cs contact.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newContactRequest NewContactRequest
		err := json.NewDecoder(r.Body).Decode(&newContactRequest)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id := r.Context().Value(ContextIDKey).(string)
		accountDetails, err := as.GetAccountByUsername(newContactRequest.Account)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = cs.AddContact(id, accountDetails.ID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok"))
	}
}

//NewGroupHandleFunc creates a new group with its requestor as admin. Adds all other members subseqently
func NewGroupHandleFunc(as account.Service, cs contact.Service, gs group.Service, gas groupAccount.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newGroupContactRequest NewGroupContactRequest
		err := json.NewDecoder(r.Body).Decode(&newGroupContactRequest)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id := r.Context().Value(ContextIDKey).(string)
		groupID, err := gs.AddGroup(id, newGroupContactRequest.Name, newGroupContactRequest.Description)

		err = cs.AddGroupContact(id, groupID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, account := range newGroupContactRequest.Accounts {
			accountDetals, err := as.GetAccountByUsername(account)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = gas.AddAccountToGroup(accountDetals.ID, groupID, false)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = cs.AddGroupContact(accountDetals.ID, groupID)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("ok"))
	}
}

//GetMessageHandleFunc returns a list of messages from the requested date for the token's user.
//This function ASSUMES THE JWT IS VALIDATED. Username is obtain from JWT claims.
// func GetMessageHandleFunc(s message.Service) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var requestMessages RequestMessages
// 		json.NewDecoder(r.Body).Decode(&requestMessages)

// 		headerToken := r.Header.Get("Authorization")
// 		tokenString := strings.Split(headerToken, "Bearer ")[1]

// 		var claims security.JwtClaims
// 		var p jwt.Parser

// 		_, _, err := p.ParseUnverified(tokenString, &claims)
// 		if err != nil {
// 			log.Println(err.Error())
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
// 		}

// 		messages, err := s.GetMessage(requestMessages.DateTime, claims.ID)
// 		if err != nil {
// 			log.Println(err.Error())
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(messages)
// 	}
// }
