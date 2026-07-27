# Example Go Microservice with Database

Учебный проект микросервиса на Go, построенный на принципах **Clean Architecture**.
Проект демонстрирует правильную организацию кода, разделение ответственности, использование интерфейсов для абстракции инфраструктуры и внедрение зависимостей.


## 🎯 Назначение проекта

Этот проект создан в **учебных целях** для изучения:

- Архитектурных паттернов в Go
- Работы с HTTP-сервером и middleware
- Генерации уникальных идентификаторов (UUID v7)
- Абстракции времени для тестируемости
- Интеграции с базой данных (SQLite)
- Dependency Injection и интерфейсов

Проект представляет собой минимальный API-сервер, который возвращает текущее время с использованием разных источников (системное время или время из БД).


### Структура каталогов
```
├── cmd/
│ └── server/
│ │ ├── main.go # Точка входа
│ │ └── http_handler.go # Сборка роутера и middleware
├── internal/ # Внутренний код
│ ├── domains/ # БИЗНЕС-СЛОЙ
│ │ ├── clock.go # Clock интерфейс
│ │ ├── datetime_service.go # Сервис времени
│ │ └── id_gen.go # IDGenerator интерфейс
│ ├── infra/ # ИНФРАСТРУКТУРА
│ │ ├── sqlite_clock/ # SQLite реализация Clock
│ │ │ ├── clock.go # SQLiteClock
│ │ │ └── open.go # Подключение к SQLite
│ │ ├── system_clock.go # Системная реализация Clock
│ │ └── uuid_generator.go # UUIDGenerator (v7)
│ └── protocol/ # ПРОТОКОЛ
│ └── httpapi/ # HTTP API слой
│ ├── datetime_handler.go # Обработчик времени
│ ├── middleware.go # RequestID middleware
│ ├── request_context.go # Контекст для RequestID
│ └── request_id.go # Константы и helpers
├── go.mod
└── go.sum


├── cmd
│ └── server
│     ├── http_handler.go
│     └── main.go
├── go.mod
├── go.sum
├── internal
│ ├── domains
│ │   ├── clock.go
│ │   ├── datetime_service.go
│ │   └── id_gen.go
│ ├── infra
│ │        ├── sqlite_clock
│ │        │     ├── clock.go
│ │        │     └── open.go
│ │        ├── system_clock.go
│ │        └── uuid_generator.go
│ └── protocol
│     └── httpapi
│         ├── datetime_handler.go
│         ├── middleware.go
│         ├── request_context.go
│         └── request_id.go
└── README.md


```

## 🚀 Быстрый старт

### Установка и запуск

```bash
git clone https://github.com/AssemblerBossss/example-go-microservice-with-db.git
cd example-go-microservice-with-db
go mod tidy

go run cmd/server/main.go