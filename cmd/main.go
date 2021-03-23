package main

import (
	"log"
	"net"
	"net/http"

	"github.com/nicholaslim94/messenger_backend/pkg/account"
	"github.com/nicholaslim94/messenger_backend/pkg/contact"
	"github.com/nicholaslim94/messenger_backend/pkg/group"
	"github.com/nicholaslim94/messenger_backend/pkg/groupAccount"
	"github.com/nicholaslim94/messenger_backend/pkg/groupMessage"
	"github.com/nicholaslim94/messenger_backend/pkg/login"
	"github.com/nicholaslim94/messenger_backend/pkg/message"
	"github.com/nicholaslim94/messenger_backend/pkg/net/rest"
	"github.com/nicholaslim94/messenger_backend/pkg/net/tcp"
	"github.com/nicholaslim94/messenger_backend/pkg/repository/postgres"
)

func main() {
	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Init database
	db, err := postgres.Connect("localhost", 5432, "messenger", "postgres", "password", false)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Init Repositories
	accountRepo := postgres.NewAccountRepository(db)
	contactRepo := postgres.NewContactRepository(db)
	messageRepo := postgres.NewMessageRepository(db)
	groupRepo := postgres.NewGroupRepository(db)
	groupAccountRepo := postgres.NewGroupAccountRepository(db)
	groupMessageRepo := postgres.NewGroupMessageRepository(db)

	//Init Services
	loginService := login.NewService(accountRepo)
	accountService := account.NewService(accountRepo)
	messageService := message.NewService(messageRepo)
	contactService := contact.NewService(contactRepo)
	groupService := group.NewService(groupRepo, groupAccountRepo)
	groupAccountService := groupAccount.NewService(groupAccountRepo)
	groupMessageService := groupMessage.NewService(groupMessageRepo, groupAccountRepo)

	//Init TCP and listen on port
	go func() {
		log.Println("starting tcp sever")
		listerner, err := net.Listen("tcp", ":8000")
		if err != nil {
			log.Println(err)
			return
		}
		defer listerner.Close()
		for {
			con, err := listerner.Accept()
			if err != nil {
				log.Println(err)
				return
			}
			go tcp.HandleConn(con, messageService, contactService, accountService, groupMessageService, groupService)
		}
	}()

	//Init and serve http controller
	log.Println("starting http sever")
	http.HandleFunc("/createAccount", rest.CreateAccountHandler(accountService))
	http.HandleFunc("/checkUsernameAvail", rest.CheckUsernameHandlerFunc(accountService))
	http.HandleFunc("/checkEmailAvail", rest.CheckEmailHandlerFunc(accountService))
	http.HandleFunc("/login", rest.LoginHandleFunc(loginService))

	// http handler with middleware
	newContacthandler := http.HandlerFunc(rest.NewContactHandleFunc(accountService, contactService))
	http.Handle("/newContact", rest.JwtAuthenticate(newContacthandler))

	newGrouphandler := http.HandlerFunc(rest.NewGroupHandleFunc(accountService, contactService, groupService, groupAccountService))
	http.Handle("/newGroup", rest.JwtAuthenticate(newGrouphandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
