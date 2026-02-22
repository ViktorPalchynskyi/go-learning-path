# Study-Go

Проект для изучения языка Go по программе из заметки `021-3 Go MOC.md`.

## Темы

1. Основы синтаксиса — переменные, типы, функции, пакеты
2. Structs и Methods — структуры, методы, receivers
3. Interfaces — implicit interfaces, duck typing
4. Pointers — указатели, pass by value/reference
5. Slices и Maps — коллекции данных, итерация
6. Error Handling — error, sentinel errors, wrapping
7. Packages и Modules — go.mod, структура проекта
8. JSON и REST API — encoding/json, struct tags
9. Goroutines и Channels — goroutines, channels, select
10. Context — context для запросов и отмены
11. net/http и Routing — HTTP-сервер, Chi router, middleware
12. Testing — go test, table-driven tests, mocks

## Формат коммитов

Git Flow с коротким описанием в одно предложение:

```
<type>: <короткое описание>
```

Типы:
- `feat` — новая функциональность или урок
- `fix` — исправление ошибки
- `refactor` — рефакторинг без изменения поведения
- `docs` — изменения в документации
- `test` — добавление или изменение тестов
- `chore` — прочие изменения (настройка, зависимости)
