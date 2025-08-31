# Web Scraper in Go
This is a backend server that is used to aggregate data from RSS feeds, which will be used to collect posts (blogs, articles, podcasts, etc.)

This project was developed using Go (v 1.23.2) and PostgreSQL (v 17).

<img width="425" height="202" alt="Screenshot 2025-08-26 at 6 43 26 PM" src="https://github.com/user-attachments/assets/594c1f11-876c-45f5-9850-f846329302a5" />


It also uses chi, sqlc and goose. They can be installed using the following commands:

```bash
# Install the chi library
go get github.com/go-chi/chi

# Install the CLI tools
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
```

To setup the server, run the following command:

```bash
make setup
```

Change DB_URL variable in the Makefile to the correct one based on your system.

Create a .env file specifying the PORT and DB_URL variables.

Run the following command to start the server:

```bash
make run
```

Postman or curl can be used to test the endpoints. Some require an authentication header with the user's API Key. 
That would look something like this:

<img width="884" height="162" alt="Screenshot 2025-08-30 at 9 58 13 PM" src="https://github.com/user-attachments/assets/5f49fce9-dd49-4fe1-b245-12a849f4357a" />

The endpoints are listed below that can be used with http://localhost:8080/v1/

1. POST /users
2. POST /feeds
