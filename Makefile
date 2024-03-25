postgres:
	@echo "Starting postgres"
	docker run --name postgres_alpine -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=simple_bank -p 5432:5432 -d postgres:alpine
	@echo "Postgres started"

startpostgres:
	brew services start postgresql@15

stoppostgres:
	brew services stop postgresql@15

startdockerimage:
	docker start postgres_alpine

stopdockerimage:
	docker stop postgres_alpine

createdb:
	@echo "Creating database"
	docker exec -it postgres_alpine createdb --username=user --owner=user simple_bank
	@echo "Database created"

dropdb:
	@echo "Dropping database"
	docker exec -it postgres_alpine dropdb --username=user simple_bank
	@echo "Database dropped"

migrateup:
	@echo "Migrating up"
	goose postgres postgresql://user:password@localhost:5432/simple_bank up
	@echo "Migrate up completed"

migratedown:
	@echo "Migrating down"
	goose postgres postgresql://user:password@localhost:5432/simple_bank down
	@echo "Migrate down completed"

sqlc:
	@echo "Generating sqlc"
	sqlc generate
	@echo "Sqlc generated"

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test startpostgres startdockerimage stopdockerimage