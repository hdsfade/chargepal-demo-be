setup:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init --dir cmd/server/
	env GOOS=linux GOARCH=amd64 go build -o bin/server cmd/server/main.go

build:
	docker compose build --no-cache

up:
	docker compose up

down:
	docker compose down

restart:
	docker compose restart

clean:
	docker stop chargepal-demo-be
	docker stop dockerPostgres
	docker stop dockerRedis
	docker rm chargepal-demo-be
	docker rm dockerPostgres
	docker rm dockerRedis
	docker image rm chargepal-demo-be
	rm -rf .dbdata
