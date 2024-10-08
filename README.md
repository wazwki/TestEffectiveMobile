## ТЗ: "Реализация онлайн библиотеки песен"

### Необходимо реализовать следующее:

### 1. REST методы:
- Получение данных библиотеки с фильтрацией по всем полям и пагинацией
- Получение текста песни с пагинацией по куплетам
- Удаление песни
- Изменение данных песни
- Добавление новой песни в формате JSON

```json
{
 "group": "Muse",
 "song": "Supermassive Black Hole"
}
```

### 2. При Create запросе к создаваемому сервису - сделать запрос во внешний API, описанный Swagger'ом:

```yaml
openapi: 3.0.3
info:
  title: Music info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: group
          in: query
          required: true
          schema:
            type: string
        - name: song
          in: query
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/SongDetail'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    SongDetail:
      required:
        - releaseDate
        - text
        - link
      type: object
      properties:
        releaseDate:
          type: string
          example: 16.07.2006
        text:
          type: string
          example: Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight
        link:
          type: string
          example: https://www.youtube.com/watch?v=Xsp3_a-PMTw

```

### 3. Обогащенную информацию положить в БД postgres (структура БД должна быть создана путем миграций при старте сервиса).

### 4. Покрыть код debug- и info-логами.

### 5. Вынести конфигурационные данные в .env-файл.

### 6. Сгенерировать swagger на реализованное API.

--- 

## Реализация:

### Структура проекта:

```markdown

TestEffectiveMobile
│
├── cmd
│   └── main.go          # Точка входа в приложение
├── db
│   ├── migrations       # SQL миграции для создания таблиц
│   │   ├── 000001_create_songs_table.up.sql
│   │   └── 000001_create_songs_table.down.sql
│   └── db.go            # Инициализация подключения к базе данных
├── docs                 # Swagger-документация
│   ├── docs.go
│   ├── index.html
│   ├── swagger-initializer.js
│   ├── swagger-ui-bundle.js
│   ├── swagger-ui-standalone-preset.js
│   ├── swagger-ui.css
│   ├── swagger.json
│   └── swagger.yaml
├── internal
│   ├── handlers
│   │   └── handler.go   # REST-хэндлеры
│   ├── models
│   │   └── models.go    # Структуры данных
│   ├── service
│   │   └── service.go   # Бизнес-логика
│   └── repository
│       └── repo.go      # Работа с базой данных (PostgreSQL)
├── mock_api
│   └── server.go        # Mock api для запросов
├── pkg
│   └── logger
│       └── logger.go    # Логгирование
├── .dockerfile
├── .env                 # Конфигурационный файл
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── log.log              # Логи
└── README.md

```

### Использование:

#### Для запуска проекта выполните команду:

```bash

docker-compose up

```