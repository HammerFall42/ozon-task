# Ozon internship task


## Запуск с помощью Docker

```
make run_in_mem //для запуска in-memory решения
make run_in_db //для запуска с использованием postgresql
```

## Запросы к API

```
host: "localhost:8080"
basePath: "/api/v1"
paths:
  /url
    get:
      consumes:
      - "application/json"
      produces:
      - "application/json"
      parameters:
      - name: "shortened"
        in: "query"
        description: "Get full url by shortened"
        required: true
        type: "string"
      responses:
        "200":
          type: "string"
          description: "full url"
        "500":
          description: "No url with such shortened version"
    post:
      consumes:
      - "application/json"
      produces:
      - "application/json"
      parameters:
      - name: "url"
        in: "query"
        description: "Post full url and get shortened"
        required: true
        type: "string"
      responses:
        "200":
          type: "string"
          description: "shortened url"
        "500":
          description: "Failed to add new url"
```

Пример для Postman:

```
GET localhost:8080/api/v1/url
Headers: Postman-Token, Content-Type, Host, User-Agent, Accept, Accept-Encoding, Connection
Body: raw JSON
{
  "shortened" : "0000000001"
}

POST localhost:8080/api/v1/url
Headers: Postman-Token, Content-Type, Host, User-Agent, Accept, Accept-Encoding, Connection
Body: raw JSON
{
  "url" : "https://docs.google.com/document/d/1PYRU5vLqdktzkldlO0VZS-IZfoPL6yJinJTWcDLkWU0/edit"
}
```

## База данных

База инициализируется при использовании соответствующего варианта.

Предыдущее содержание базы удаляется.
Это сделано для того, чтобы результат работы обеих реализаций был одинаковым.

## Задание

### Укорачиватель ссылок

Необходимо реализовать сервис, который должен предоставлять API по созданию сокращённых ссылок следующего формата:
- Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться только одна сокращенная ссылка.
- Ссылка должна быть длинной 10 символов
- Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание)

Сервис должен быть написан на Go и принимать следующие запросы по http:
1. Метод Post, который будет сохранять оригинальный URL в базе и возвращать сокращённый
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный URL

Решение должно быть предоставлено в «конечном виде», а именно:
- Сервис должен быть распространён в виде Docker-образа 
- В качестве хранилища ожидается использовать in-memory решение И postgresql. Какое хранилище использовать указывается параметром при запуске сервиса. 
- Покрыть реализованный функционал Unit-тестами
