.PHONY: build
build:
	go build .

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

.DEFAULT_GOAL := build
