# Brain Games
Лабораторная работа №1 по Методологии разработки ПО Котов Дмитрий Викторович БИВТ-22-СП-2 
## Описание игр

### 1. Наименьшее общее кратное (НОК)
Игра, в которой пользователю предлагается вычислить наименьшее общее кратное трёх случайных чисел.

### 2. Геометрическая прогрессия
Игра, в которой пользователю предлагается определить пропущенное число в геометрической прогрессии.

## Установка и запуск

1. **Установка Go**: Убедитесь, что у вас установлен Go на компьютере.
2. **Клонирование репозитория**: Клонируйте репозиторий с помощью команды: `git clone https://github.com/your-username/brain-games.git`
3. **Переход в папку проекта**: `cd brain-games`
4. **Запуск игры**: `go run main.go`

## Настройка линтера

Для проверки кода используется `golangci-lint`. Чтобы настроить линтер:

1. **Установка golangci-lint**: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
2. **Запуск линтера**: `golangci-lint run ` или через makefile `make lint`

[![Maintainability](https://api.codeclimate.com/v1/badges/79c532d89c61e77fc4f4/maintainability)](https://codeclimate.com/github/kotovski000/brain-games/maintainability)

