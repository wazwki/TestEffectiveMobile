## Реализация онлайн библиотеки песен

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
          example: Ooh baby, don't you know I suffer?
        patronymic:
          type: string
          example: https://www.youtube.com/watch?v=Xsp3_a-PMTw
```

### 3. Обогащенную информацию положить в БД postgres (структура БД должна быть создана путем миграций при старте сервиса)

### 4. Покрыть код debug- и info-логами

### 5. Вынести конфигурационные данные в .env-файл

### 6. Сгенерировать swagger на реализованное API


```markdown
TestEffectiveMobile
│
├── cmd
│   └── main.go          # Точка входа в приложение
├── db
│   ├── migrations       # SQL миграции для создания таблиц
│   └── db.go            # Инициализация подключения к базе данных
├── docs                 # Swagger-документация
├── internal
│   ├── handlers
│   │   └── handler.go   # REST-хэндлеры
│   ├── service
│   │   └── service.go   # Бизнес-логика
│   └── repository
│       └── repo.go      # Работа с базой данных (PostgreSQL)
├── pkg
│   └── logger
│       └── logger.go    # Логгирование
├── .env                 # Конфигурационный файл
├── .gitignore
├── go.mod
├── go.sum
├── log.log              # Логи
└── README.md
```