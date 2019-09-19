-include .env
export

ifeq ($(ENV),)
	ENV := development
endif
include .env.$(ENV)

# Server Configuration.
start:
	@GO111MODULE=on go run main.go

install: # use goose with github version 
	@go get github.com/pressly/goose/cmd/goose
	@go get ./...
	@GO111MODULE=on go mod tidy
	@GO111MODULE=on go mod vendor

migrate: # Run the MySQL migration.
	@echo operating on database: ${DB_NAME}
	@goose -dir migrations mysql "${DB_USER}:${DB_PASS}@tcp(${DB_HOST})/${DB_NAME}" up

migrate-dbtest: # Run the MySQL migration.
	@echo operating on database: ${TEST_DB_NAME}
	@goose -dir migrations mysql "${TEST_DB_USER}:${TEST_DB_PASS}@tcp(${TEST_DB_HOST})/${TEST_DB_NAME}" up


rollback: # Rollback to previous migration.
	@echo operating on database: ${DB_NAME}
	@goose -dir migrations mysql "${DB_USER}:${DB_PASS}@tcp(${DB_HOST})/${DB_NAME}" down

test: # Run the test.
	@go test -bench=. -v ./...

clean: # Clear the MySQL tmp/ directory.
	@rm -rf tmp

review:
	@go get -u github.com/kisielk/errcheck; ls -la; errcheck ./...
	@go get honnef.co/go/tools/cmd/staticcheck; staticcheck -checks all ./...
	@go vet ./...
	@go get github.com/securego/gosec/cmd/gosec; gosec ./...

