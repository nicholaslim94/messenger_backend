package tcp

import (
	"encoding/json"

	"github.com/nicholaslim94/messenger_backend/pkg/account"
	"github.com/nicholaslim94/messenger_backend/pkg/contact"
	"github.com/nicholaslim94/messenger_backend/pkg/group"
	"github.com/nicholaslim94/messenger_backend/pkg/groupMessage"
	"github.com/nicholaslim94/messenger_backend/pkg/message"
)

//Response is a wraps the type and data
type Response struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

//Request is a wrapper containing the type and data
type Request struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

//RequestToken is a dto containing the token
type RequestToken struct {
	Token string `json:"token"`
}

//RequestContacts is a dto containing the object for removing Account
// type RequestContacts struct {
// }

//UpdateContactLastRead is a request dto to update last read id in contact
type UpdateContactLastRead struct {
	AccountID  string `json:"accountId"`
	LastReadID string `json:"lastReadId"`
}

//UpdateContactGroupLastRead is a request dto to update group last read id in contact
type UpdateContactGroupLastRead struct {
	GroupID    string `json:"groupId"`
	LastReadID string `json:"lastReadId"`
}

//AddMessage is a request dto to add a new message
type AddMessage struct {
	To  string `json:"to"`
	Msg string `json:"msg"`
}

//ToDomian maps dto to domain model
func (s *AddMessage) ToDomian(from string) *message.Model {
	return &message.Model{
		From: from,
		To:   s.To,
		Msg:  s.Msg,
	}
}

//AddGroupMessage is a request dto to add a new group message
type AddGroupMessage struct {
	GroupID string `json:"groupId"`
	Msg     string `json:"msg"`
}

//ToDomian maps dto to domain model
func (s *AddGroupMessage) ToDomian(from string) *groupMessage.Model {
	return &groupMessage.Model{
		AccountID: from,
		GroupID:   s.GroupID,
		Msg:       s.Msg,
	}
}

//RequestMessages is a request dto to get messages
type RequestMessages struct {
	DateTime string `json:"dateTime"`
}

//RequestGroupMessages is a request dto to get group messages
type RequestGroupMessages struct {
	DateTime string `json:"dateTime"`
}

//RequestAccount is a request dto to get Account
type RequestAccount struct {
	ID string `json:"id"`
}

//RequestGroup is a request dto to get group
type RequestGroup struct {
	ID string `json:"id"`
}

//ResponseContact is a response dto to send Contacts
type ResponseContact struct {
	ID          string `json:"id"`
	AccountID   string `json:"accountId"`
	GroupID     string `json:"groupId"`
	LastReadID  string `json:"lastReadId"`
	CreatedDate string `json:"createdDt"`
}

//ToResponseContacts maps slice of contact domain to slice of ResponseContacts dto
func ToResponseContacts(contacts []*contact.Model) []*ResponseContact {
	var responseContacts []*ResponseContact
	for _, contact := range contacts {
		responseContact := &ResponseContact{
			ID:          contact.ID,
			AccountID:   contact.ChildID,
			GroupID:     contact.GroupID,
			LastReadID:  contact.LastReadID,
			CreatedDate: contact.CreatedDt,
		}
		responseContacts = append(responseContacts, responseContact)
	}
	return responseContacts
}

//ResponseMessage is a response dto to send Messages
type ResponseMessage struct {
	ID       string `json:"id"`
	From     string `json:"from"`
	To       string `json:"to"`
	Msg      string `json:"msg"`
	DateTime string `json:"dateTime"`
}

//ToResponseMessages maps slice of contact domain to slice of ResponseMessage dto
func ToResponseMessages(messages []*message.Model) []*ResponseMessage {
	var responseMessages []*ResponseMessage
	for _, message := range messages {
		responseMessage := &ResponseMessage{
			ID:       message.ID,
			From:     message.From,
			To:       message.To,
			Msg:      message.Msg,
			DateTime: message.DateTime,
		}
		responseMessages = append(responseMessages, responseMessage)
	}
	return responseMessages
}

//ResponseMessage is a response dto to send Messages
type ResponseGroupMessage struct {
	ID        string `json:"id"`
	GroupID   string `json:"groupId"`
	AccountID string `json:"accountId"`
	Msg       string `json:"msg"`
	DateTime  string `json:"dateTime"`
}

//ToResponseMessages maps slice of contact domain to slice of ResponseMessage dto
func ToResponseGroupMessages(groupMessages []*groupMessage.Model) []*ResponseGroupMessage {
	var responseGroupMessages []*ResponseGroupMessage
	for _, message := range groupMessages {
		responseGroupMessage := &ResponseGroupMessage{
			ID:        message.ID,
			GroupID:   message.GroupID,
			AccountID: message.AccountID,
			Msg:       message.Msg,
			DateTime:  message.DateTime,
		}
		responseGroupMessages = append(responseGroupMessages, responseGroupMessage)
	}
	return responseGroupMessages
}

//ResponseAccount is a response dto to send Account
type ResponseAccount struct {
	Username  string `json:"username"`
	AccountID string `json:"accountId"`
	CreatedDt string `json:"createdDt"`
}

//ToResponseAccount maps account domain to ResponseAccount dto
func ToResponseAccount(account *account.Model) *ResponseAccount {
	return &ResponseAccount{
		Username:  account.Username,
		AccountID: account.ID,
		CreatedDt: account.CreatedDt,
	}
}

//ResponseGroup is a response dto to send Group
type ResponseGroup struct {
	Name      string `json:"name"`
	GroupID   string `json:"groupId"`
	CreatedDt string `json:"createdDt"`
}

//ToResponseGroup maps account domain to ResponseGroup dto
func ToResponseGroup(group *group.Model) *ResponseGroup {
	return &ResponseGroup{
		Name:      group.Name,
		GroupID:   group.ID,
		CreatedDt: group.CreatedDt,
	}
}
