# Load environment variables from .env
include .env
export $(shell sed 's/=.*//' .env)

# Start the development server with air
dev:
	air

# Database migration settings
MIGRATION_DIR=./db/migrations
DB_URI=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Database migration commands
migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

# Get current migration version
migrate-version:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URI)" version

# Apply all up migrations
migrate-up:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URI)" up

# Rollback flag to last version
migrate-fix:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URI)" force $(version)

# Rollback to a specific version
migrate-rollback:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URI)" goto $(version)