-include .env
BIN:=${shell go env GOPATH}/bin

vendor:
	go mod vendor

install-deps: vendor
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

start-db:
	docker-compose -f ./deployments/docker-compose.yaml up --remove-orphans -d 2>/dev/null
	until pg_isready -qh localhost -U postgres; do sleep 0.1; done

create-migration:
	${BIN}/migrate create -ext sql -dir migrations ${name}

migrate-up:
	${BIN}/migrate -path migrations/ -database ${DATABASE_URL} up

migrate-down:
	${BIN}/migrate -path migrations/ -database ${DATABASE_URL} down

lint:
	${BIN}/golangci-lint run

generate-swagger:
	${BIN}/swagger generate spec -o ./api/swagger.json --scan-models

test-unit: vendor
	go test $$(go list ./... | grep -v integrationtest)

run:
	go run cmd/miniwiki/miniwiki.go