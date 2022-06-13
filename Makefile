.PHONY: build-all
build-all: wire build

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
wire:	delete_wire
	./scripts/wire_generate.sh

.PHONY: swagger
swagger:
	./scripts/generate_swagger.sh

.PHONY: delete_wire
delete_wire:
	find . -name "*_gen.go" -exec rm -f {} \;

.DEFAULT_GOAL := build
