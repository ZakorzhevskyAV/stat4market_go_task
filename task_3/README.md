## stat4market_task_3

clickhouse-client --user chuser --password chpass --host clickhouse-server --database default

Данный сервис создаёт таблицу в ClickHouse и позволяет передать данные в БД через POST-запрос.

Для разворачивания сервиса нужно ввести команду:

docker-compose up (-d)

Для сворачивания сервиса нужно ввести команду:

docker-compose down (--remove-orphans)

Сервис по-умолчанию принимает запросы через порт 8000.

Для передачи данных в ClickHouse нужно сделать запрос следующей формы:

POST localhost:8000/api/event

{
"eventType": "<login>",
"userID": <1>,
"eventTime": <"2023-04-09 13:00:00">,
"payload": <"{\"some_field\":\"some_value\"}">
}

где в треугольных скобках находятся передаваемые значения с их соответствующими названиями.

При записи данных в БД значение EventID генерируется случайно и является числом от 0 до 9.