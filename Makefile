.PHONY: build
build:
	go build cmd/todo-server/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...


.PHONY: up
up:
	docker compose up -d

.PHONY: stop
stop:
	docker compose stop

.PHONY: dsh
dsh:
	dsh todo-app

.PHONY: wire
wire:
	./scripts/wire_generate.sh

.PHONY: swagger
swagger:
	./scripts/generate_swagger.sh

.DEFAULT_GOAL := build
