## Запуск приложения

### Конфигурация
Хранится в директории configs/
* database_url - параметры для запуска бд
* bind_addr - параметры для запуска сервера

Флаги при запуске
* -config-path - путь к файлу с конфигурацией (-config-path=path)
  по дефолту стоит ../configs/shorturlservice.toml
* -db-option - выбор бд ( -db-option=postgres, -db-option=inmemory) по дефолту стоит postgres 

### Запуск с докера
Для запуска приложения с использованием docker-compose: build/docker-compose.yml
```shell
# нужно находится в папке build
cd build/
sudo docker-compose up -d
#по дефолту запуститься с postgres базой данных, изменить это можно в build/build.sh
```

### Запуск и тестирование локально
### Make
```shell
# нужно находится в папке build
cd build/
make build
#соберет бинарник shorturlservice
make test
#запуск unit-тестов
```

### Go build
```shell
# нужно находится в папке build
cd build/
go build -v ./../cmd/shorturlservice
./shorturlservice -db-option=postgres # запуск с postgres
./shorturlservice -db-option=inmemory # запуск с локальной бд
#запуск unit-тестов
go test -v  -race -timeout 30s ./../...
```

## Структура проекта

* /build/ - содержит инструменты для запуска тестов и окружения
* /cmd/ - приложение http-сервер обеспечивающий прием и обработку REST API запросов
* /docs/ - документация к проекту
* /internal/ - реализация микросервиса.
* /configs/ - файлы конфигурации для запуска сервиса.(host, url базы данных)
* /migrations/ - миграции базы данных

### /
* build.sh - скрипт для запуска сервиса в докере
* database.env - файл конфигурации бд в докере
* docker-compose.yml - запуск микросервиса (вместе с субд)
* Makefile - make файл для локального запуска

# Тестирование
```shell
➜  build git:(main) ✗ docker-compose up -d   
➜  build git:(main) ✗ http POST http://localhost:8080/urls url=yandex.ru
HTTP/1.1 201 Created
Content-Length: 26
Content-Type: text/plain; charset=utf-8
Date: Sun, 24 Oct 2021 12:28:22 GMT

{
    "shortUrl": "bi.ly/5hnO"
}

➜  build git:(main) ✗ http GET http://localhost:8080/urls url=bi.ly/5hnO          
HTTP/1.1 200 OK
Content-Length: 24
Content-Type: text/plain; charset=utf-8
Date: Sun, 24 Oct 2021 12:28:37 GMT

{
    "longUrl": "yandex.ru"
}

```