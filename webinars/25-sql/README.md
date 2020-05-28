.center.icon[![otus main](img/main.png)]

---

class: top white
background-image: url(img/sound.svg)
background-size: 130%
.top.icon[![otus main](img/logo.png)]

.sound-top[
  # Как меня слышно и видно?
]

.sound-bottom[
  ## > Напишите в чат
  ### **+** если все хорошо
  ### **-** если есть проблемы cо звуком или с видео
  ### !проверить запись!
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Работа с SQL

### Елена Граховац, Антон Телышев

---

# План занятия

.big-list[
* Установка и работа с PostgreSQL
* Простейшие SQL запросы
* Подключение к СУБД и настройка пула подключений
* Выполнение запросов и получение результатов
* Стандартные интерфейсы sql.DB, sql.Rows, sql.Tx
* Использование транзакций
* SQL инъекции и борьба с ними
]

---

# Напишите в чат о вашем опыте с PostgreSQL

Например:

* есть опыт работы с PostgreSQL в Go
* есть опыт работы с другим SQL БД в Go
* есть опыт работы с PostgreSQL, но не в Go
* нет опыта работы с PostgreSQL, но есть опыт с другими SQL БД
* нет опыта работы с SQL БД

---

# Работаем с PostgreSQL локально

Устанавливаем сервер из консоли (пример для Ubuntu):
```
# обновить пакеты
$ sudo apt-get update

# установить PostgreSQL сервер и клиент
$ sudo apt-get install postgresql-10

# запустить PostgreSQL
$ sudo systemctl start postgresql

# подключиться под пользователем, созданным по умолчанию
$ sudo -u postgres psql
```

---

# Работаем с PostgreSQL локально

Работаем в клиенте СУБД:
```sql
-- создаем "проектного" пользователя СУБД
postgres=# create user myuser password 'mypass';
CREATE ROLE

-- создаем "проектную" базу данных
postgres=# create database mydb owner myuser;
CREATE DATABASE
```

Удобный клиент с графическим интерфейсом: https://www.pgadmin.org/download/

---

# Работаем с PostgreSQL локально через Docker

(См. https://hub.docker.com/_/postgres)

Создаем сеть для доступа между контейнерами:

```
docker network create postgres
```

Запускаем контейнер с сервером PostgreSQL:
```
docker run --network postgres --name dpostgres \
    -e POSTGRES_PASSWORD=password -d postgres
```

Подключаемся к серверу:
```
docker run --network postgres -it --rm postgres \
    psql -h dpostgres -U postgres
```

---


# Data Definition Language (DDL)

Создание простой таблицы и индекса (файл 001.sql):

```sql
create table events (
  id          serial primary key,
  owner       bigint,
  title       text,
  descr       text,
  start_date  date not null,
  start_time  time,
  end_date    date not null,
  end_time    time
);
create index owner_idx on events (owner);
create index start_idx on events using btree (start_date, start_time);
```

Выполнение SQL скрипта из консоли:
```
psql 'host=localhost user=myuser password=mypass dbname=mydb'  < 001.sql
```
---

# Data Manipulation Language (DML)

Добавление строки:

```sql
insert into events(owner, title, descr, start_date, end_date) 
values(42, 'new year', 'watch the irony of fate', 
        '2019-12-31', '2019-12-31')
returning id;
```
<br><br>
Обновление строки:
```sql
update events
set end_date = '2020-01-01'
where id = 1;
```

---

# Data Query Language (DQL)

Получение одной строки:
```sql
select * from events where id = 1
```
Получение нескольких строк:
```sql
select id, title, descr
from events
where owner = 42
  and start_date = '2020-01-01'
```

<br><br>
ResultSet:
```
 id(bigint) |  title(text)   |          descr(text)
------------+----------------+-------------------------
  2         | new year       | watch the irony of fate
  3         | prepare ny     | make some olive
```
---

# Подключение к PostgreSQL из Go, простейший вариант

Создание подключения:

```go
import "database/sql"
import _ "github.com/jackc/pgx/stdlib"

dsn := "..."
db, err := sql.Open("pgx", dsn)  // *sql.DB
if err != nil {
  log.Fatalf("failed to load driver: %v", err)
}
// создан пул соединений
```

Использование подключения:

```go
err := db.PingContext(ctx)
if err != nil {
  return xerrors.Errorf("failed to connect to db: %v", err)
}
// работаем с db
```

http://go-database-sql.org/importing.html <br>
http://go-database-sql.org/accessing.html

---

# DataSourceName

DSN - строка подключения к базе, содержит все необходимые опции. 
<br>
Синтаксис DSN зависит от используемой базы данных и драйвера.
<br><br>

Например для PostgreSQL:
```
"postgres://myuser:mypass@localhost:5432/mydb?sslmode=verify-full"
```
Или
```
"user=myuser dbname=mydb sslmode=verify-full password=mypass"
```
* `host` - Сервер базы данных или путь к UNIX-сокету (по-умолчанию localhost)
* `port` - Порт базы данных (по-умолчанию 5432)
* `dbname` - Имя базы данных
* `user` - Пользователь в СУБД (по умолчанию - пользователь OS)
* `password` - Пароль пользователя

Подробнее: [https://godoc.org/github.com/lib/pq](https://godoc.org/github.com/lib/pq)
---

# Пул соединений

`sql.DB` - это пул соединений с базой данных. Соединения будут открываться по мере необходимости.
<br><br>
`sql.DB` - безопасен для конкурентного использования (так же как `http.Client`)
<br><br>

Настройки пула:
```go
// Макс. число открытых соединений от этого процесса
db.SetMaxOpenConns(n int)

// Макс. число открытых неиспользуемых соединений
db.SetMaxIdleConns(n int)

// Макс. время жизни одного подключения
db.SetConnMaxLifetime(d time.Duration)
```

http://go-database-sql.org/connection-pool.html


---

# Пул соединений
.main-image[
 ![img/no_pool.png](img/no_pool.png)
]


---

# Пул соединений

.main-image[
 ![img/with_pool.png](img/with_pool.png)
]


---

# Выполнение запросов

```go
query := `insert into events(owner, title, descr, start_date, end_date) 
  values($1, $2, $3, $4, $5)`

result, err := db.ExecContext(ctx, query,
  42, "new year", "watch the irony of fate", "2019-12-31", "2019-12-31"
) // sql.Result
if err != nil {
  // обработать ошибку
}

// Авто-генерируемый ID (SERIAL)
eventId, err := result.LastInsertId()      // int64

// Количество измененных строк
rowsAffected, err := result.RowsAffected() // int64
```

http://go-database-sql.org/retrieving.html

---

# Получение результатов

```go
query := `
  select id, title, descr
  from events
  where owner = $1 and start_date = $2
`
rows, err := db.QueryContext(ctx, query, owner, date)
if err != nil {
  // ошибка при выполнении запроса
}
defer rows.Close()

for rows.Next() {
  var id int64
  var title, descr string
  if err := rows.Scan(&id, &title, &descr); err != nil {
    // ошибка сканирования
  }
  // обрабатываем строку
  fmt.Printf("%d %s %s\n", id, title, descr)
}
if err := rows.Err(); err != nil {
	// ошибка при получении результатов
}
```

---

# Объект sql.Rows

```go
// возвращает имена колонок в выборке
rows.Columns() ([]string, error)

// возвращает типы колонок в выборке
rows.ColumnTypes() ([]*ColumnType, error)

// переходит к следующей строке или возвращает false
rows.Next() bool

// заполняет переменные из текущей строки
rows.Scan(dest ...interface{}) error

// закрывает объект Rows
rows.Close()

// возвращает ошибку, встреченную при итерации
rows.Error() error
```

---

# Получение одной строки

```go
query := "select * from events where id = $1"

row := db.QueryRowContext(ctx, query, id)

var id int64
var title, descr string

err := row.Scan(&id, &title, &descr)

if err == sql.ErrNoRows {
  // строки не найдено
} else if err != nil {
  // "настоящая" ошибка
}
```

---

# PreparedStatements

*PreparedStatement* - это заранее разобранный запрос, который можно выполнять повторно.
*PreparedStatement* - временный объект, который *создается в СУБД* и живет в рамках сессии, или пока не будет закрыт.

```go
// создаем подготовленный запрос
stmt, err := db.Prepare("delete from events where id = $1") // *sql.Stmt
if err != nil {
	log.Fatal(err)
}

// освобождаем ресурсы в СУБД
defer stmt.Close()

// многократно выполняем запрос
for _, id := range ids {
  _, err := stmt.Exec(id)
  if err != nil {
    log.Fatal(err)
  }
}
```

http://go-database-sql.org/prepared.html

---

# Работа с соединением

`*sql.DB` - это пул соединений. Даже последовательные запросы могут использовать *разные* соединения с базой.
<br><br>
Если нужно получить одно конкретное соединение, то
```go
conn, err := db.Conn(ctx)  // *sql.Conn

// вернуть соединение в pool
defer conn.Close()

// далее - обычная работа как с *sql.DB
err := conn.ExecContext(ctx, query1, arg1, arg2)

rows, err := conn.QueryContext(ctx, query2, arg1, arg2)
```

---

# Транзакции

Транзакция - группа запросов, которые либо выполняются, либо не выполняются вместе. 
Внутри транзакции все запросы видят "согласованное" состояние данных.
<br><br>
На уровне SQL для транзакций используются отдельные запросы: `BEGIN`, `COMMIT`, `ROLLBACK`.
<br><br>
Работа с транзакциями в Go:
```go
tx, err := db.BeginTx(ctx, nil)  // *sql.Tx
if err != nil {
  log.Fatal(err)
}

// далее - обычная работа как с *sql.DB
err := tx.ExecContext(ctx, query1, arg1, arg2)
rows, err := tx.QueryContext(ctx, query2, arg1, arg2)

err := tx.Commit()  // или tx.Rollback()
if err != nil {
  // commit не прошел, данные не изменились
}

// далее объект tx не пригоден для использования
```

http://go-database-sql.org/modifying.html


---

# Основные методы

Определены у `*sql.DB`, `*sql.Conn`, `*sql.Tx`, `*sql.Stmt`:
```go
// изменение данных
ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)

// получение данных (select)
QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)

// получение одной строки
QueryRowContext(ctx context.Context, query string, args ...interface{}) *Row
```
<br><br>
Внимание, ошибка:
```go
_, err := db.QueryContext(ctx, "delete from events where id = $1", 42)
```
---

# NULL

В SQL базах любая колонка может быть объявлена к NULL / NOT NULL. 
NULL - это не 0 и не пустая строка, это отсутствие значения.
```sql
create table users (
  id          serial primary key,
  name        text not null,
  age         int null
);
```

Для обработки `NULL` в Go предлагается использовать специальные типы:
```go
var id, realAge int64
var name string
var age sql.NullInt64
err := db.QueryRowContext(ctx, "select * from users where id = 1").Scan(&id, &name, &age)

if age.Valid {
  realAge = age.Int64
} else {
	// обработка на ваше усмотрение
}
```

---

# SQL Injection

Опасно:
```go
query := "select * from users where name = '" + name + "'"
query := fmt.Sprintf("select * from users where name = '%s'", name)
```
Потому что в `name` может оказаться что-то типа:
```
"jack'; truncate users; select 'pawned"
```

Правильный подход - использовать `placeholders` для подстановки значений в SQL:
```go
row := db.QueryRowContext(ctx, "select * from users where name = $1", name)
```

Однако это не всегда возможно. Так работать не будет:
```go
db.QueryRowContext(ctx, "select * from $1 where name = $2", table, name)
db.QueryRowContext(ctx, "select * from user order by $1 limit 3", order)
```

Проверить код на инъекции (и другие проблемы безопасности): https://github.com/securego/gosec

---


# Проблемы database/sql

.big-list[
* placeholder зависят от базы: (`$1` в Postgres, `?` в MySQL, `:name` в Oracle)
* Есть только базовые типы, но нет, например `sql.NullDate`
* `rows.Scan(arg1, arg2, arg3)` - неудобен, нужно помнить порядок и типы колонок.
* Нет возможности `rows.StructScan(&event)`
]

---

# Расширение sqlx

`jmoiron/sqlx` - обертка, прозрачно расширяющая стандартную библиотеку `database/sql`.
<br><br>
* `sqlx.DB` - обертка над `*sql.DB`
* `sqlx.Tx` - обертка над `*sql.Tx`
* `sqlx.Stmt` - обертка над `*sql.Stmt`
* `sqlx.NamedStmt` - PreparedStatement с поддержкой именованых параметров

Подключение `jmoiron/sqlx`:

```go
import "github.com/jmoiron/sqlx"

db, err := sqlx.Open("pgx", dsn)  // *sqlx.DB

rows, err := db.QueryContext("select * from events") // *sqlx.Rows

...
```
---

# sqlx: именованные placeholder'ы

Можно передавать параметры запроса в виде словаря:
```go
sql := "select * from events where owner = :owner and start_date = :start"
rows, err := db.NamedQueryContext(ctx, sql, map[string]interface{}{
  "owner": 42,
  "start": "2019-12-31",
})
```

Или структуры:
```go
type QueryArgs{
  Owner int64
  Start string
}
sql := "select * from events where owner = :owner and start_date = :start"
rows, err := db.NamedQueryContext(ctx, sql, QueryArgs{
  Owner: 42,
  Start: "2019-12-31",
})
```
---

# sqlx: сканирование

Можно сканировать результаты в словарь:
```go
sql := "select * from events where start_date > $1"

rows, err := db.QueryContext(ctx, sql, "2020-01-01") // *sqlx.Rows

for rows.Next() {
  results := make(map[string]interface{})
  err := rows.MapScan(results)
  if err != nil {
    log.Fatal(err)
  }
  // обрабатываем result
}
```
---

# sqlx: сканирование

Можно сканировать результаты структуру:
```
type Event {
  Id          int64
  Title       string
  Description string   `db:"descr"`
}

sql := "select * from events where start_date > $1"

rows, err := db.NamedQueryContext(ctx, sql, "2020-01-01") // *sqlx.Rows

events := make([]Event)

for rows.Next() {
  var event Event
  err := rows.StructScan(&event)
  if err != nil {
    log.Fatal(err)
  }
  events = append(events, event)
}
```
---

# Драйверы для Postgres

* Лучший драйвер на текущий момент: [https://github.com/jackc/pgx](https://github.com/jackc/pgx)
* Другой часто используемый драйвер (менее производительный): [https://github.com/lib/pq](https://github.com/lib/pq)

---

# Миграции

* [https://github.com/pressly/goose](https://github.com/pressly/goose) - можно использовать как cli-тулзу и как библиотеку
* [https://flywaydb.org/](https://flywaydb.org/) - пожалуй, самая популярная штука для миграций

*Protip*: flyway можно запускать из докер-контейнера перед запуском основного приложения, см. https://github.com/flyway/flyway-docker

---

# ORM

* [https://gorm.io/](https://gorm.io/) - использует пустые интерфейсы :(
* [https://github.com/go-reform/reform](https://github.com/go-reform/reform) - использует кодогенерацию, но разработка немного заброшена

Пример использования (не идеальный код, но может пригодиться): https://github.com/rumyantseva/mif

---

# Другие ресурсы для изучения

.big-list[
* [en] [https://github.com/rumyantseva/pglocal](Пример окружения для локальной разработки)
* [ru] [https://habr.com/ru/company/oleg-bunin/blog/461935/](Статья и видео о тонкастях работы с Postgres в Go)
* [en] [http://go-database-sql.org/index.html](Полезная документация по работе с Postgres из Go)
* [en] [https://golang.org/pkg/database/sql](Документация к database/sql)
* [en] [https://jmoiron.github.io/sqlx](sqlx)
]

---

# Домашнее задание

Заготовка микросервиса «Календарь»
<br>
[задание](https://github.com/OtusGolang/home_work/blob/master/hw12_13_14_15_calendar/docs/12_README.md)


---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
[https://otus.ru/polls/8471/](https://otus.ru/polls/8471/)
]

.right-image[
![](img/gopher7.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
