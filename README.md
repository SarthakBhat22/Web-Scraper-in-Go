# Web Scraper in Go
This is a backend server that is used to aggregate data from RSS feeds, which will be used to collect posts (blogs, articles, podcasts, etc.)

This project was developed using Go (v 1.23.2) and PostgreSQL (v 17).

<img width="644" height="256" alt="Screenshot 2025-08-26 at 3 01 47 PM" src="https://github.com/user-attachments/assets/af93e76a-c1dd-45a9-9697-1925e9d554a7" />

It also uses chi, sqlc and goose. They can be installed using the following commands:

```bash
# Install the chi library
go get github.com/go-chi/chi

# Install the CLI tools
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
```
