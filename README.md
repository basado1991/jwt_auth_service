## jwt_auth_service

### Конфигурация
`.env.app`:
```sh
AS_POSTGRES=<URL для подключения к PostgreSQL>
AS_ADDR=<адрес, куда биндиться>
AS_KEY_PATH=<путь к файлу HS512 ключа>

AS_MAIL_USERNAME=<имя пользователя на почтовом сервисе>
AS_MAIL_PASSWORD=<пароль на почтовом сервисе>

AS_MAIL_FROM=<имя, от которого слать письма>
AS_MAIL_HOST=<хост почтового сервиса (yandex: smtp.yandex.ru)>
AS_MAIL_ADDR=<адрес почтового сервиса (yandex: smtp.yandex.ru:25)>
```

`.env.postgres`:
```sh
POSTGRES_PASSWORD=<пароль>
POSTGRES_USER=<имя пользователя>
```

### Запуск
Для начала нужно настроить базу данных:
```sh
$ dbmate -u "<URL для подключения к PostgreSQL>" up
```

И... запуск!
```sh
$ podman compose up
```
или через `docker`
```sh
$ docker compose up
```

### Документация по API
Документация по всем эндпоинтам сервиса написана по спецификации OpenAPI и находится в `docs/main.yaml`.
