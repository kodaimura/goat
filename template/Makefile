dcup:
	docker compose up -d

dcdown:
	docker compose down

dcstart:
	docker compose start

dcstop:
	docker compose stop

in:
	docker compose exec app bash

indb:
	docker compose exec db bash

dcbuild:
	docker compose build --no-cache

init:
	go mod tidy
	go build -o gent cmd/gent/main.go

run:
	go run cmd/goat/main.go

test:
	go test
