# UralCTF-visit
## Список хэгдлеров
`POST /api/teams/register` - Регистрирует новую команду с капитаном и участниками.

Валидация:
   - Название команды — обязательное, уникальное.
   - Город и учебное заведение — обязательные, город из Яндекс API, вуз привязан к городу.
   - ФИО капитана — обязательное.
   - Telegram капитана — @username.
   - Телефон — E.164, только +7.
   - Курс — 1–6.
   - Не более 7 участников (с капитаном).
   - Все три галочки согласий обязательны.

Действия:
   - Проверка уникальности команды.
   - Сохранение данных в БД (таблицы teams, participants).
   - Отправка письма капитану (SMTP/Yandex/SendGrid).
   - Возврат статуса.
   
`GET /api/teams/check-name` - Проверка доступности названия команды (для фронта в реальном времени).

Параметры: name (строка).Возврат: { available: true/false }.

`GET /api/cities/search` - Поиск городов через API Яндекс Карт (проксируем, чтобы не палить API-ключ фронту).

Параметры: query (строка). Возврат: массив городов.

`GET /api/universities/search` - Поиск учебных заведений в выбранном городе через API Яндекс Карт.
Параметры: city_id или query. Возврат: массив учебных заведений.

`GET /api/teams` - Возврат списка команд (например, для страницы "Участники").
Фильтры: по городу, учебному заведению, зачетной группе. Возврат: JSON-массив команд с участниками.

## Структура .env
```
SERVER_PORT=8080

DB_HOST="localhost"
DB_PORT=5432
DB_USER="uralctf"
DB_PASSWORD="uralctf_password"
DB_NAME="uralctf_visit"
DB_SSLMODE="disable"

SMTP_HOST="smtp.example.com"
SMTP_PORT=587
SMTP_USER="user"
SMTP_PASSWORD="password"
SMTP_FROM="ctf@utmn.ru"

YANDEX_API_KEY="your_yandex_api_key"

LOG_DIR="./logs"
LOG_MAX_SIZE=10
LOG_MAX_BACKUPS=5
LOG_MAX_AGE=30
```