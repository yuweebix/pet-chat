# Pet-Chat
## GoLang Real-Time Chat Application

This is my pet project implementing Chatting functionality. This application is built using GoLang and provides multiple chat rooms. Users can join any chat room and start chatting instantly.

## Key Features

- **Real-Time Messaging**: The application uses WebSockets for real-time bidirectional communication between the server and the client. This allows messages to be sent and received instantly.

- **User Authentication**: User authentication is handled using JWT (JSON Web Tokens). After a successful login, the server generates a token that the client stores and includes in all subsequent requests.

- **Admin Users**: The application supports admin users who have the ability to manage chat rooms and users. Admins can delete chat rooms and ban users.

- **Caching**: To improve performance, frequently accessed data is cached using Redis.

## Technologies and Methods

- **Backend**: The backend of the application is built with GoLang. It handles user requests, interacts with the database, and sends responses back to the client.

- **Database**: The application uses a PostgreSQL database to store data. GORM (Go Object-Relational Mapper) is used as an ORM (Object-Relational Mapping) tool for easier and more efficient database operations.

- **Frontend**: The frontend of the application is built with HTML and CSS. HTMX is used to add interactivity to the web pages without needing to write JavaScript.

## Project Structure

```
/pet-chat
├── cmd
│   └── server
│       └── main.go
├── pkg
│   ├── handlers
│   │   ├── handler.go
│   │   ├── user.go
│   │   ├── room.go
│   │   ├── message.go
│   │   └── admin.go
│   ├── models
│   │   ├── user.go
│   │   ├── room.go
│   │   └── message.go
│   └── repository
│       ├── db.go
│       ├── user.go
│       ├── room.go
│       ├── message.go
│       └── admin.go
├── middleware
│   └── auth.go
├── scripts
├── web
│   ├── static
│   │   ├── css
│   │   │   └── style.css
│   └── templates
│       ├── layout.html
│       ├── index.html
│       ├── login.html
│       ├── register.html
│       ├── chatroom.html
│       └── admin.html
├── .env
├── .gitignore
├── Dockerfile
├── go.mod
└── README.md
```
