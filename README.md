# Guitar Go API

![Go](https://img.shields.io/badge/Go-1.22+-blue.svg)
![Swagger](https://img.shields.io/badge/Swagger-2.0-green.svg)
![License](https://img.shields.io/badge/License-Apache%202.0-orange.svg)

API для управления гитарным магазином с аутентификацией JWT и админ-панелью

## Содержание

- [Требования](#требования)
- [Установка](#установка)
  - [Windows](#windows)
  - [Linux](#linux)
- [Конфигурация](#конфигурация)
- [Использование](#использование)
- [API Документация](#api-документация)
- [Разработка](#разработка)
- [Docker](#docker)

## Требования

- Go 1.22+
- SQLite3 или PostgreSQL
- Для генерации документации: Swag CLI

## Установка

### Windows

1. Установите Go с [официального сайта](https://golang.org/dl/)
2. Клонируйте репозиторий:
   ```powershell
   git clone https://github.com/yourusername/guitar-go.git
   cd guitar-go
   ```
3. Установите зависимости:
   ```powershell
   go mod download
   ```
4. Установите Swag CLI:
   ```powershell
   go install github.com/swaggo/swag/cmd/swag@latest
   ```
5. Сгенерируйте документацию:
   ```powershell
   swag init -g ./cmd/main.go --output docs
   ```

### Linux

1. Установите Go:
   ```bash
   sudo apt update
   sudo apt install golang-go
   ```
2. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/yourusername/guitar-go.git
   cd guitar-go
   ```
3. Установите зависимости:
   ```bash
   go mod download
   ```
4. Установите Swag CLI:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```
5. Сгенерируйте документацию:
   ```bash
   swag init -g ./cmd/main.go --output docs
   ```

## Конфигурация

1. Скопируйте пример конфига:
   ```bash
   cp configs/config.example.yaml configs/config.yaml
   ```
2. Отредактируйте `config.yaml` под свои нужды
3. Для разработки можно создать `.env` файл (см. `.env.example`)

## Использование

### Запуск приложения

```bash
go run cmd/main.go --config=./configs/config.yaml
```

Или с билдом:

```bash
go build -o guitar-go ./cmd/main.go
./guitar-go --config=./configs/config.yaml
```

### Создание первого пользователя

1. Отправьте POST запрос на `/register`:
   ```json
   {
     "username": "admin",
     "password": "securepassword",
     "role": "admin"
   }
   ```
2. Авторизуйтесь через `/login` чтобы получить JWT токен

## API Документация

После запуска приложения документация доступна по адресу:  
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Разработка

### Команды Makefile

```bash
make build    # Собрать приложение
make run      # Запустить приложение
make swagger  # Обновить документацию
make test     # Запустить тесты
```

### Структура проекта

```
guitar-go/
├── cmd/               # Основной пакет приложения
├── configs/           # Конфигурационные файлы
├── docs/              # Генерируемая документация Swagger
├── internal/
│   ├── app/           # Инициализация приложения
│   ├── config/        # Конфигурация
│   ├── db/            # Работа с БД
│   ├── handlers/      # HTTP обработчики
│   ├── middleware/    # Промежуточное ПО
│   ├── models/        # Модели данных
│   ├── repositories/  # Репозитории
│   └── services/      # Бизнес-логика
├── pkg/               # Вспомогательные пакеты
└── Dockerfile         # Конфигурация Docker
```

## Docker

### Сборка образа

```bash
docker build -t guitar-go .
```

### Запуск контейнера

```bash
docker run -p 8080:8080 guitar-go
```

### Docker Compose (с PostgreSQL)

1. Создайте `docker-compose.yml`:

   ```yaml
   version: "3"

   services:
     app:
       build: .
       ports:
         - "8080:8080"
       depends_on:
         - db
       environment:
         - DATABASE_DRIVER=postgres
         - DATABASE_POSTGRES_URL=postgres://user:pass@db:5432/guitar_go?sslmode=disable

     db:
       image: postgres:13
       environment:
         - POSTGRES_USER=user
         - POSTGRES_PASSWORD=pass
         - POSTGRES_DB=guitar_go
       volumes:
         - postgres_data:/var/lib/postgresql/data

   volumes:
     postgres_data:
   ```

2. Запустите:
   ```bash
   docker-compose up -d
   ```
