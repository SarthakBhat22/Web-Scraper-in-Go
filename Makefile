BINARY_NAME=Web-Scraper-in-Go
DB_URL=postgres://<username>:<password>@localhost:5432/<yourdb>

.PHONY: setup build run

setup:
	@echo "Running SQLC to generate Go code from SQL..."
	sqlc generate
	@echo "Running database migrations with Goose..."
	goose -dir sql/schema postgres "$(DB_URL)" up

build:
	go build -o $(BINARY_NAME) .

run: build
	@echo "Starting the server..."
	./$(BINARY_NAME)
