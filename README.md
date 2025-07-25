# Task API

Минимальный HTTP API для управления долгими задачами (создание, удаление, получение статуса). Все данные хранятся в памяти сервиса.

## Запуск

```sh
git clone https://github.com/AlexAntinomy/task_api.git
cd task_api
go run .
```

Сервер стартует на порту 8080.

## Эндпоинты

### Создать задачу
- **POST** `/tasks`
- Ответ: JSON с id и статусом задачи

### Получить статус/результат
- **GET** `/tasks/{id}`
- Ответ: JSON с id, статусом, датой создания, длительностью, результатом

### Удалить задачу
- **DELETE** `/tasks/{id}`
- Ответ: 204 No Content

## Пример запроса

```sh
curl -X POST http://localhost:8080/tasks
```

## Примечания
- Все задачи выполняются в фоне, результат появляется через несколько минут.
- После перезапуска сервиса все задачи теряются. 