bot:
	go run ./cmd/bot/main.go

dictionary:
	go run ./cmd/dictionary/main.go

db:
	go run ./cmd/db/main.go
	
tests:
	go test -v ./tests/

build-image:
	docker build -t pixel68tmb/telegram_bot:latest .

start-container:
	docker run --rm -idt --name telegram_bot --env-file .env pixel68tmb/telegram_bot:latest 

delete-unused-images:
	docker image prune -fa

delete-all-containers:
	docker rm $(shell docker ps -qa)