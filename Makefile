# ========================
# VARIABLES
# ========================
APP_NAME=loan-api
MAIN_PATH=./cmd/api/main.go
BINARY=bin/$(APP_NAME)

DB_URL=mysql://user:password@tcp(localhost:3306)/loan_db

# ========================
# GO COMMANDS
# ========================

.PHONY: run
run:
	go run $(MAIN_PATH)

.PHONY: build
build:
	go build -o $(BINARY) $(MAIN_PATH)

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

# ========================
# SQLC
# ========================

.PHONY: sqlc
sqlc:
	sqlc generate

# ========================
# DATABASE (MIGRATIONS)
# ========================

MIGRATE=migrate -path db/migrations -database "$(DB_URL)"

.PHONY: migrate-up
migrate-up:
	$(MIGRATE) up

.PHONY: migrate-down
migrate-down:
	$(MIGRATE) down

.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

# usage:
# make migrate-create name=create_loans_table

# ========================
# DOCKER
# ========================

.PHONY: docker-build
docker-build:
	docker build -t $(APP_NAME) .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(APP_NAME)

# ========================
# TESTING
# ========================

.PHONY: test
test:
	go test ./... -v

# ========================
# DEVELOPMENT SHORTCUTS
# ========================

.PHONY: dev
dev: tidy fmt vet sqlc run

.PHONY: rebuild
rebuild: clean build