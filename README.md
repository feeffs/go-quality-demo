# Go code quality (go-quality-demo)

Демонстрации инструментов контроля качества кода на Go.

## Структура проекта

```
go-quality-demo/
├── main.go                  # Основной код: divide, readTextFile, Add
├── math_test.go             # Тесты и бенчмарки
├── go.mod                   # Зависимости модуля
├── go.sum                   # Контрольные суммы зависимостей
├── .pre-commit-config.yaml  # Конфигурация Git Hooks
└── .gitignore
```

## Зависимости

- [go.uber.org/zap](https://github.com/uber-go/zap) — структурированное логирование

## Инструменты качества кода

| Инструмент      | Назначение                      | Команда                   |
| --------------- | ------------------------------- | ------------------------- |
| `go fmt`        | Форматирование кода             | `go fmt ./...`            |
| `go vet`        | Базовый анализ кода             | `go vet ./...`            |
| `golangci-lint` | Линтер (агрегатор анализаторов) | `golangci-lint run ./...` |
| `staticcheck`   | Статический анализ              | `staticcheck ./...`       |
| `errcheck`      | Проверка необработанных ошибок  | `errcheck ./...`          |

### Установка инструментов

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install github.com/kisielk/errcheck@latest
```

## Тесты

Запуск всех тестов:

```bash
go test ./...
```

Запуск с покрытием:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Запуск бенчмарков:

```bash
go test -bench=. -benchmem ./...
```

## Git Hooks (pre-commit)

Перед каждым коммитом автоматически запускаются: `go fmt`, `go vet`, `staticcheck`, `golangci-lint`, `errcheck`.

### Установка

```bash
pip install pre-commit
pre-commit install
```

На Windows если `pip` недоступен:

```bash
py -m pip install pre-commit
py -m pre_commit install
```

### Ручной запуск всех хуков

```bash
pre-commit run --all-files
```

## Управление зависимостями

Очистка и обновление зависимостей:

```bash
go mod tidy
```
