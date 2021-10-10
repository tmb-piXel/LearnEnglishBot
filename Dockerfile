FROM golang:1.17-alpine3.14 AS builder

RUN go version

COPY . /github.com/tmb-piXel/LearEnglishBot/
WORKDIR /github.com/tmb-piXel/LearEnglishBot/

RUN go mod download
RUN go build -o ./.bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/tmb-piXel/LearEnglishBot/.bin/bot .
COPY --from=0 /github.com/tmb-piXel/LearEnglishBot/configs configs/
COPY dictionaries /root/dictionaries

CMD ["./bot"]