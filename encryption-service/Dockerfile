FROM golang:1.22.4

# Устанавливаем рабочую директорию для приложения
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь код проекта
COPY . /app

# Устанавливаем рабочую директорию для команды "go build"
WORKDIR /app

# Собираем исполняемый файл
RUN go build -o /encryption-service 

# Копируем исполняемый файл в корневую директорию образа: COPY /encryption-service /


# Запускаем исполняемый файл
CMD ["/encryption-service"]
