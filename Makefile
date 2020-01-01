all: build run

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o watchlist-server main.go
	docker build --tag=watchlist .

run:
	docker-compose up watchlist db

clean:
	docker container stop watchlist-server_watchlist_1
	docker container rm watchlist-server_watchlist_1
	docker container stop watchlist-server_db_1
	docker container rm watchlist-server_db_1

.PHONY: build run clean