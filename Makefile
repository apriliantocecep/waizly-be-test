MIGRATION_DIR=database/migrations
MIGRATION_EXT=sql
MIGRATION_COMMAND=migrate
DATABASE_URL=mysql://root:admin@tcp(localhost:3306)/waizly

create_table:
	@echo "Creating table $(table)"
	@$(MIGRATION_COMMAND) create -ext $(MIGRATION_EXT) -dir $(MIGRATION_DIR) create_$(table)_table

migrate_up:
	@echo "Running database migrations"
	@$(MIGRATION_COMMAND) -database "$(DATABASE_URL)" -path $(MIGRATION_DIR) up

wire:
	wire ./internal/server/

run:
	go run cmd/web/main.go


t:
	go test -v ./test/ -coverprofile=coverage.out

c:
	go tool cover -html=coverage.out -o coverage.html