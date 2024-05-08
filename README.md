# Pet-Chat

This is my pet project implementing Chatting functionality with the new features of Go 1.22.

## Technologies and Methods

- **Backend**: The backend uses Go 1.22 with the hopes that the new advanced routing capabilities will be enough to route to create a full-fledged app that looks decent.

- **Middleware**: For this project I have implemented middleware that does logging, checks for whether a user is authenticated or not and chaining.

- **Database**: Seeing as it's a fairly small project, I opted to use GORM with PostgreSQL as its dialect.

- **Frontend**: The frontend will be built with HTML, CSS, and HTMX will be used to add interactivity to the web pages without the need to write JavaScript or use React.

## How to make it work

- Check go.mod
- Start up your PostgreSQL server
- Set up your .env file like so:
```
# Database variables
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=257.257.257.257
DB_PORT=5432
DB_NAME=pet_db

# IP Adress
DOMAIN=localhost
PORT=42069

# Admin initialisation
ADMIN_EMAIL=admin@admin.com
ADMIN_USERNAME=admin
ADMIN_PASSWORD=12345678

# Session 
EXPIRY=31  # in minutes
```
- Start up the server as per example:
```
go run ./cmd/server/
```
