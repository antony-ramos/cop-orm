## ----------------------
## Available make targets
## ----------------------
##

default: help

help: ## Display this message
	@grep -E '(^[a-zA-Z0-9_\-\.]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

##
## ----------------------
## Builds
## ----------------------
##

artifact: ## Compile app from sources (linux)
	@CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o cop-orm ./cmd/cop-orm

artifact.osx: ## Compile app from sources (osx)
	@CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o cop-orm ./cmd/cop-orm


##
## ----------------------
## Q.A
## ----------------------
##

qa: lint test ## Run all Q.A

lint.install: ## Install Go linter
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: ## Lint source code
	golangci-lint run -v

lint.fix: ## Lint and fix source code
	golangci-lint run --fix -v

test: ## Run unit tests (go test)
	@echo "### Running unit tests ..."
	go test -v ./... -coverprofile coverage.out

##
## ----------------------
## Development
## ----------------------
##

install: ## Install required goland dependencies
	go mod download

start: ## Start project
	go run ./cmd/cop-orm

init-db: create-postgres create-cop-orm-db ## Initialize the database
	@sleep 5

create-postgres: ## Create a PostgreSQL instance using Docker
	docker run --name postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

create-cop-orm-db: ## Create a cop-orm database
	docker exec -it postgres createdb -U postgres cop-orm