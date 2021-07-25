run:
	go run ./cmd/bot/main.go

telebot:
	go run ./cmd/telebot/main.go

readDictionary:
	go run ./cmd/dictionary/main.go
	
run-tests:
	go test -v ./tests/

build-image:
	docker build -t pixel68tmb/telegram_bot:latest .

start-container:
	docker run --rm -idt --name telegram_bot --env-file .env pixel68tmb/telegram_bot:latest 

delete-unused-images:
	docker image prune -fa

delete-all-containers:
	docker rm $(shell docker ps -qa)