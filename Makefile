build:
	go build cmd/goat/main.go

run:
	ENV=local go run cmd/goat/main.go

start:
	nohup ./main &

test:
	go test