DB_DSN ?= postgres://postgres:1987@localhost:8088/tasksdb?sslmode=disable
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создание новой миграции: make migrate-new NAME=tasks
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

format:
	gofmt -s -w .
	
lint:
	golangci-lint run --output.text.colors=true
	
lint-fix:
	golangci-lint run --fix