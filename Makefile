# Production
deploy:
	go build cmd/goat/main.go
	nohup ./main &

# Local
lrun:
	ENV=local go run cmd/goat/main.go

ltest:
	ENV=local go test

# Local (docker compose)
up:
	docker compose up -d

down:
	docker compose down

start:
	docker compose start

stop:
	docker compose stop

in:
	docker exec -i -t goat_app bash

db:
	docker exec -i -t goat_db bash

clear:
	docker compose build --no-cache

# Docker Container
run:
	go run cmd/goat/main.go



