package tcp

import (
	"encoding/json"
	"io"
	"log"
	"net"

	"github.com/dgrijalva/jwt-go"
	"github.com/nicholaslim94/messenger_backend/pkg/account"
	"github.com/nicholaslim94/messenger_backend/pkg/contact"
	"github.com/nicholaslim94/messenger_backend/pkg/group"
	"github.com/nicholaslim94/messenger_backend/pkg/groupMessage"
	"github.com/nicholaslim94/messenger_backend/pkg/security"

	"github.com/nicholaslim94/messenger_backend/pkg/message"
)

//HandleConn handles incoming TCP connects
func HandleConn(c net.Conn, ms message.Service, cs contact.Service, as account.Service, gms groupMessage.Service, gs group.Service) {
	defer c.Close()
	var requestToken RequestToken
	conDecoder := json.NewDecoder(c)
	conEncoder := json.NewEncoder(c)
	err := conDecoder.Decode(&requestToken)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(requestToken.Token)
	claims := &security.JwtClaims{}
	token, err := jwt.ParseWithClaims(requestToken.Token, claims, func(token *jwt.Token) (interface{}, error) {
		token.SigningString()
		return security.GetSecret(), err
	})
	if err != nil || !token.Valid {
		log.Println("Bad token. Rejecting Connection: " + err.Error())
		return
	}
	log.Println("Credentials ok! Establishing connection")

	for {
		var request Request

		err := conDecoder.Decode(&request)
		if err == io.EOF {
			log.Print("Client Disconnected")
			c.Close()
			return
		}
		if err != nil {
			log.Print(err)
			continue
		}
		// log.Println(request)

		switch request.Type {
		case requestContacts:
			contacts, err := cs.GetContacts(claims.ID)
			if err != nil {
				log.Print(err)
				continue
			}
			jsonByteArray, err := json.Marshal(ToResponseContacts(contacts))
			if err != nil {
				log.Print(err)
				continue
			}
			response := Response{
				Type: responseContacts,
				Data: jsonByteArray,
			}
			conEncoder.Encode(response)

		case updateContactLastRead:
			var updateContactLastRead UpdateContactLastRead
			err = json.Unmarshal(request.Data, &updateContactLastRead)
			if err != nil {
				log.Print(err)
				continue
			}
			err = cs.UpdateLastRead(claims.ID, updateContactLastRead.AccountID, updateContactLastRead.LastReadID)
			if err != nil {
				log.Print(err)
				continue
			}

		case updateContactGroupLastRead:
			var updateContactGroupLastRead UpdateContactGroupLastRead
			err = json.Unmarshal(request.Data, &updateContactGroupLastRead)
			if err != nil {
				log.Print(err)
				continue
			}
			err = cs.UpdateGroupLastRead(claims.ID, updateContactGroupLastRead.GroupID, updateContactGroupLastRead.LastReadID)
			if err != nil {
				log.Print(err)
				continue
			}

		case addMessage:
			var addMessage AddMessage
			err = json.Unmarshal(request.Data, &addMessage)
			if err != nil {
				log.Print(err)
				continue
			}
			err = ms.AddMessage(addMessage.ToDomian(claims.ID))
			if err != nil {
				log.Print(err)
				continue
			}

		case requestMessages:
			var requestMessages RequestMessages
			err = json.Unmarshal(request.Data, &requestMessages)
			if err != nil {
				log.Print(err)
				continue
			}
			messageModels, err := ms.GetMessage(requestMessages.DateTime, claims.ID)
			if err != nil {
				log.Print(err)
				continue
			}
			jsonByteArray, err := json.Marshal(ToResponseMessages(messageModels))
			if err != nil {
				log.Print(err)
				continue
			}
			response := Response{
				Type: responseMessages,
				Data: jsonByteArray,
			}
			conEncoder.Encode(response)

		case addGroupMessage:
			var addGroupMessage AddGroupMessage
			err = json.Unmarshal(request.Data, &addGroupMessage)
			if err != nil {
				log.Print(err)
				continue
			}
			err = gms.AddGroupMessage(addGroupMessage.ToDomian(claims.ID))
			if err != nil {
				log.Print(err)
				continue
			}

		case requestGroupMessages:
			var requestGroupMessages RequestGroupMessages
			err = json.Unmarshal(request.Data, &requestGroupMessages)
			if err != nil {
				log.Print(err)
				continue
			}
			groupMessageModels, err := gms.GetGroupMessage(requestGroupMessages.DateTime, claims.ID)
			if err != nil {
				log.Print(err)
				continue
			}
			jsonByteArray, err := json.Marshal(ToResponseGroupMessages(groupMessageModels))
			if err != nil {
				log.Print(err)
				continue
			}
			response := Response{
				Type: responseGroupMessages,
				Data: jsonByteArray,
			}
			conEncoder.Encode(response)

		case requestAccount:
			var requestAccount RequestAccount
			err = json.Unmarshal(request.Data, &requestAccount)
			if err != nil {
				log.Print(err)
				continue
			}
			accountModel, err := as.GetAccountByID(requestAccount.ID)
			if err != nil {
				log.Print(err)
				continue
			}
			jsonByteArray, err := json.Marshal(ToResponseAccount(accountModel))
			if err != nil {
				log.Print(err)
				continue
			}
			response := Response{
				Type: responseAccount,
				Data: jsonByteArray,
			}
			conEncoder.Encode(response)

		case requestGroup:
			var requestGroup RequestGroup
			err = json.Unmarshal(request.Data, &requestGroup)
			if err != nil {
				log.Print(err)
				continue
			}
			groupModel, err := gs.FindGroupByID(requestGroup.ID)
			if err != nil {
				log.Print(err)
				continue
			}
			jsonByteArray, err := json.Marshal(ToResponseGroup(groupModel))
			if err != nil {
				log.Print(err)
				continue
			}
			response := Response{
				Type: responseGroup,
				Data: jsonByteArray,
			}
			conEncoder.Encode(response)
		}
	}
}
