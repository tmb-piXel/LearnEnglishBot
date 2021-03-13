run:
	go run main.go

build-image:
	docker build -t telegram-bot:v0 .

start-container:
	docker run --rm -iddot --name telegram-bot telegram-bot:v0

delete-all-images:
	docker rmi $(docker images -aq)

delete-all-containers:
	docker rm $(docker ps -qamake)