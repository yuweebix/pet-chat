# Pet-Chat
## Go Real-Time Chat Application

This is my pet project implementing Chatting functionality without the use of third-party services with the new advanced features of Go 1.22. This application provides multiple chat rooms. Users can join any chat room and start chatting.

## Technologies and Methods

- **Backend**: The backend of the application is built with Go. It handles user requests, interacts with the database, and sends responses back to the client.

- **Database**: The application uses a PostgreSQL database to store data. GORM is used as an ORM tool for easier and more efficient database operations.

- **Frontend**: The frontend of the application is built with HTML and CSS. HTMX is used to add interactivity to the web pages without needing to write JavaScript.

## Project Structure

```bash
├── Dockerfile
├── README.md
├── cmd
│   └── server
│       └── main.go
├── go.mod
├── go.sum
├── pkg
│   ├── handlers
│   │   ├── admin.go
│   │   ├── handler.go
│   │   ├── message.go
│   │   ├── room.go
│   │   └── user.go
│   ├── middleware
│   │   ├── chaining.go
│   │   └── logging.go
│   ├── models
│   │   ├── message.go
│   │   ├── room.go
│   │   └── user.go
│   ├── repository
│   │   ├── admin.go
│   │   ├── db.go
│   │   ├── message.go
│   │   ├── room.go
│   │   └── user.go
│   └── utils
│       └── utils.go
└── web
    ├── static
    │   ├── css
    │   │   └── style.css
    │   └── js
    │       └── script.js
    └── templates
        ├── admin.html
        ├── chatroom.html
        ├── index.html
        ├── layout.html
        └── user
            ├── login.html
            └── register.html
```
