# Overview
Messenger_backend is a GOLANG messaging backend service using raw TCP sockets, RESTFul API and SQL database(Postgres).

## Project Purpose
Coming from the Java world, this project is my first attempt at the GO language.
This project purpose is to
1. Learn the GO Language
2. Better understand how TCP works
3. My own reasoning to how a messaging backend may work (with SQL database)
4. Experiment with flutter (Not in this project)
5. My first post to GitHub
6. Get a job :/

## Description
Messenger_backend exposes both RESTFUL API and TCPSOCKET to interact with other users either individually or as a group.
Occational tasks such as creating accounts or creating groups are done RESTFUL-ly. 
More extensive task such as checking for new messages and sending messages are done over TCP Socket.
### RESTFUL API
See /pkg/net/rest/README.md
### TCP
See /pkg/net/tcp/README.md

## Limitation
As this is Not a messaging backend based on Firesbase Cloud Messaging, there is no Push Notification.
From my understanding implementing a custom Push notification is no small feat and it is a project on its own. :(

## How to run?
Prerequisite - Latest GO version, Postgres Database
1. Create database name it: [messenger]
2. Create all tables provided in sql.txt in the ROOT folder
3. In CMD/CLI Navigate to /CMD folder
4. Run go project command: [go run main.go]
