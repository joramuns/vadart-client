FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app cmd/telebot/main.go

EXPOSE 443

ENTRYPOINT ["./app"]
