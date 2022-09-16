
install:
	@cd src && \
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

migrate:
	migrate -path ./src/database/migrations -database ${POSTGRES_URL} -verbose up

db-down:
	migrate -path ./src/database/migrations -database ${POSTGRES_URL} -verbose down

db-reset:
	migrate -path ./src/database/migrations -database ${POSTGRES_URL} drop -f

reset-db: db-reset migrate

build:
	@cd src && \
	go build -o ../bank .

run:
	@cd src && \
	go run .

vrun:
	@cd src && \
	go run . -verbose

test:
	@cd src && \
	go run . -test

vtest:
	@cd src && \
	go run . -test -verbose

sqlc:
	@cd src/database && \
	sqlc generate

dbshell:
	psql ${POSTGRES_URL}