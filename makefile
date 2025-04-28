APP_NAME=aio
DB_URL=postgres://aio:aio_secure_password@localhost:5432/aio_development?sslmode=disable
MAIN=cmd/server/main.go

run-local:
	ENVIRONMENT=local go run ${MAIN}
run-dev:
	ENVIRONMENT=development go run ${MAIN}
run-stg:
	ENVIRONMENT=staging go run ${MAIN}
run:
	ENVIRONMENT=production go run ${MAIN}

build-local:
	ENVIRONMENT=local go build -o bin/$(APP_NAME) $(MAIN)
build-dev:
	ENVIRONMENT=development go build -o bin/$(APP_NAME) $(MAIN)
build-stg:
	ENVIRONMENT=staging go build -o bin/$(APP_NAME) $(MAIN)
build:
	go build -o bin/$(APP_NAME) $(MAIN)

test:
	go test ./...

test-coverage:
	go test ./... -coverprofile=./tmp/coverage.out
	go tool cover -html=./tmp/coverage.out