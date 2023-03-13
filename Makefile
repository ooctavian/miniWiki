-include .env
BIN:=${shell go env GOPATH}/bin

vendor:
	go mod vendor

install-deps: vendor
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	go install github.com/jschaf/pggen/cmd/pggen@2023-01-27
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2
	#go install github.com/vektra/mockery/v2@v2.20.0

start-db:
	docker-compose up --remove-orphans -d 2>/dev/null
	until pg_isready -qh localhost -U postgres; do sleep 0.1; done

create-migration:
	${BIN}/migrate create -ext sql -dir migrations ${name}

migrate-up:
	${BIN}/migrate -path migrations/ -database ${DATABASE_URL} up

migrate-down:
	${BIN}/migrate -path migrations/ -database ${DATABASE_URL} down

generate-queries:
	${BIN}/pggen gen go --schema-glob "migrations/*.up.sql" --query-glob "domain/resource/query/*.sql"
	${BIN}/pggen gen go --schema-glob "migrations/*.up.sql" --query-glob "domain/category/query/*.sql"
	${BIN}/pggen gen go --schema-glob "migrations/*.up.sql" --query-glob "domain/profile/query/*.sql"
	${BIN}/pggen gen go --schema-glob "migrations/*.up.sql" --query-glob "domain/account/query/*.sql" --go-type 'domain_email=string' --go-type 'varchar=string'
	${BIN}/pggen gen go --schema-glob "migrations/*.up.sql" --query-glob "domain/auth/query/*.sql" --go-type 'domain_email=string' --go-type 'timestamp=time.Time' --go-type 'varchar=string'

seed-db:
	go run cmd/polluter/polluter.go

lint:
	${BIN}/golangci-lint run

run:
	go run cmd/miniwiki/miniwiki.go