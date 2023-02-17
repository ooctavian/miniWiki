-include .env
BIN:=${shell go env GOPATH}/bin

install-deps:
	go mod vendor
	go install -tags 'pgx' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
	go install github.com/jschaf/pggen/cmd/pggen@2023-01-27

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
	${BIN}/pggen gen go  --schema-glob "migrations/*.up.sql" --query-glob "domain/resource/query/*.sql"
	${BIN}/pggen gen go  --schema-glob "migrations/*.up.sql" --query-glob "domain/category/query/*.sql"

run:
	go run .