FROM golang:1.15-alpine3.12 AS builder

RUN go version

COPY . /github.com/tmb-piXel/telegramBotForLearningEnglish/
WORKDIR /github.com/tmb-piXel/telegramBotForLearningEnglish/

RUN go mod download
RUN go build -o ./.bin/bot main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/tmb-piXel/telegramBotForLearningEnglish/.bin/bot .
COPY dictionary /root/ 

EXPOSE 80

CMD ["./bot"]