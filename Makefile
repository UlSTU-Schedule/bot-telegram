.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t ulstu-schedule/bot-telegram .

start-container:
	docker run --env-file .env --rm ulstu-schedule/bot-telegram
