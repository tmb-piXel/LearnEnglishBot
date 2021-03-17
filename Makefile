run:
	go run ./cmd/bot/main.go

build-image:
	docker build -t pixel68tmb/telegram_bot:latest

start-container:
	docker run --rm -idt --name telegram_bot pixel68tmb/telegram_bot:latest

delete-unused-images:
	docker image prune -fa

delete-all-containers:
	docker rm $(shell docker ps -qa)

readDictionary:
	go run ./cmd/dictionary/main.go