# Golang-migrate Makefile

DB_URL			 = postgres://postgres:admin@localhost:5432/tasks?sslmode=disable

.PHONY: run migrate-new migrate-up generate-sqlc migrate-down migration-status force-version

run:
	go run ./cmd/server

generate-sqlc:
	sqlc generate -f db/sqlc.yaml

# Create a new migration
migrate-new:
	migrate create -ext sql -dir db/migrations -seq $(name)

# Apply all migrations
migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

# Rollback last migration
migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down 1

# Rollback to specific version
migrate-down-to:
	migrate -path db/migrations -database "$(DB_URL)" goto $(version)

force-version:
	migrate -path db/migrations -database "$(DB_URL)" force $(v)

# Check status
migration-status:
	migrate -path db/migrations -database "$(DB_URL)" version

schema:
	pg_dump --schema-only --no-owner --file=db/schema.sql "$(DB_URL)"
