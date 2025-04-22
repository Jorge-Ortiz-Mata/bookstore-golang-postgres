include .env
export

DEVELOPMENT_DATABASE_URL=postgres://$(PGUSER):$(PGPASSWORD)@$(PGHOST):$(PGPORT)/$(DEVELOPMENT_DATABASE)?sslmode=disable
TEST_DATABASE_URL=postgres://$(PGUSER):$(PGPASSWORD)@$(PGHOST):$(PGPORT)/$(TEST_DATABASE)?sslmode=disable

# DB Operations

test-db-conn:
	PGUSER=$(PGUSER) PGPASSWORD=$(PGPASSWORD) PGHOST=$(PGHOST) PGPORT=$(PGPORT) psql -d postgres -c '\conninfo'

create-db-development:
	PGUSER=$(PGUSER) PGPASSWORD=$(PGPASSWORD) PGHOST=$(PGHOST) PGPORT=$(PGPORT) createdb $(DEVELOPMENT_DATABASE)

create-db-test:
	PGUSER=$(PGUSER) PGPASSWORD=$(PGPASSWORD) PGHOST=$(PGHOST) PGPORT=$(PGPORT) createdb $(TEST_DATABASE)

drop-db-development:
	PGUSER=$(PGUSER) PGPASSWORD=$(PGPASSWORD) PGHOST=$(PGHOST) PGPORT=$(PGPORT) dropdb $(DEVELOPMENT_DATABASE)

drop-db-test:
	PGUSER=$(PGUSER) PGPASSWORD=$(PGPASSWORD) PGHOST=$(PGHOST) PGPORT=$(PGPORT) dropdb $(TEST_DATABASE)

migrate-db-development:
	migrate -path $(MIGRATIONS_DIR) -database $(DEVELOPMENT_DATABASE_URL) up

migrate-db-test:
	migrate -path $(MIGRATIONS_DIR) -database $(TEST_DATABASE_URL) up

rollback-db-development:
	migrate -path $(MIGRATIONS_DIR) -database $(DEVELOPMENT_DATABASE_URL) down

rollback-db-test:
	migrate -path $(MIGRATIONS_DIR) -database $(TEST_DATABASE_URL) down

generate-migration:
	@read -p "Migration name (snake_case): " name; \
	timestamp=$$(date +%Y%m%d%H%M%S); \
	filename="$${timestamp}_$${name}"; \
	mkdir -p $(MIGRATIONS_DIR); \
	touch "$(MIGRATIONS_DIR)/$${filename}.up.sql" "$(MIGRATIONS_DIR)/$${filename}.down.sql"; \
	echo "Created: $${filename}.up.sql and .down.sql in $(MIGRATIONS_DIR)"

#  App operations

run:
	go run .

build: 
	go build

test:
	go test ./... --cover

test-logs:
	go test ./... -v --cover