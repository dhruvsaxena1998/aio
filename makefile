APP_NAME=expense-manager
DB_URL=postgres://aio:aio_secure_password@localhost:5432/aio_development?sslmode=disable
MAIN=cmd/server/main.go

run:
	go run ${MAIN}

build:
	go build -o bin/$(APP_NAME) $(MAIN)