.PHONY: default
default: run;
-include .env
BIN:=${shell go env GOPATH}/bin
TEST_DB=test_db

vendor:
	go mod vendor

install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2

install-deps: vendor install-migrate install-lint

start-db:
	docker-compose -f ./deployments/docker-compose.yaml up --remove-orphans -d
	until pg_isready -qh localhost -U postgres; do sleep 0.1; done

stop-docker:
	docker-compose -f ./deployments/docker-compose.yaml down

create-s3-buckets:
	aws --endpoint-url=${AWS_ENDPOINT} s3 mb s3://${IMAGE_PROFILE_DIR}
	aws --endpoint-url=${AWS_ENDPOINT} s3 mb s3://${IMAGE_RESOURCE_DIR}

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

prepare-test-db: start-db
	$(eval DATABASE_URL := postgresql://postgres:postgres@localhost:5432/${TEST_DB}?sslmode=disable)
	PGPASSWORD=${PGPASSWORD} psql --quiet -h localhost -U postgres -w -c "drop database if exists ${TEST_DB};"
	PGPASSWORD=${PGPASSWORD} psql --quiet -h localhost -U postgres -w -c "create database ${TEST_DB};"
	${BIN}/migrate -path migrations/ -database ${DATABASE_URL} up

test-integration: vendor prepare-test-db create-s3-buckets
	go test ./internal/integrationtests

run:
	go run cmd/miniwiki/miniwiki.go
