# Используем минимальный образ с поддержкой Go
FROM golang:1.23

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все файлы в контейнер
COPY . .

# Скачиваем зависимости и собираем приложение
RUN go mod tidy && go build -o app cmd/telebot/main.go

# Устанавливаем переменную окружения для логов
ENV LOG_FILE=/app/app.log

EXPOSE 443

# Запускаем приложение
ENTRYPOINT ["./app"]
