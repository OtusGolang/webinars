.center.icon[![otus main](../img/main.png)]

---

class: top white
background-image: url(../img/check.svg)
background-size: 130%
.top.icon[![otus main](../img/logo.png)]

.sound-top[
  # Как меня слышно и видно?
]

.sound-bottom[
	## > Напишите в чат
	+ если все хорошо
	- если есть проблемы со звуком или с видео]

---

class: white
background-image: url(../img/message.svg)
.top.icon[![otus main](../img/logo.png)]

# Контекст и низкоуровневые сетевые протоколы

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем — задаем вопросы.
* ### Чат вижу — могу ответить не сразу.
* ### После занятия — оффтопик, ответы на любые вопросы.

---

# О чем будем говорить

### 1. `context.Context`.
### 2. Сетевые протоколы.
### 3. Работа с ними в Go.

---

# Настройка на занятие

.left-text[
Пожалуйста, пройдите небольшой тест.
<br><br>
Он поможет понять, что вы уже знаете,
а&nbsp;что предстоит узнать во время занятия.
<br><br>
Ссылка в чате
]

.right-image[
![](../img/gopher_science.png)
]

---

# Контекст

# Что это?

* Что уже знаете?
* Какие идеи исходя из названия?

---

# Контекст

https://pkg.go.dev/context

https://go.dev/blog/context

---

# Пакет context

```
func Background() Context
func TODO() Context
```

```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

```
func WithValue(parent Context, key interface{}, val interface{}) Context
```

---

# Практика

# Пишем эмулятор долгих операций

---

# UDP

# Какие особенности вы знаете?

---

# UDP

- Доставка пакета не гарантируется
- Порядок сообщений не гарантируется
- Соединение не устанавливается
- Быстрый
- Подходит для потокового аудио и видео, статистики, игр

---

# TCP

# Какие особенности вы знаете?

---

# TCP

- Доставка пакета гарантируется (или получим ошибку)
- Порядок пакетов гарантируется
- Соединение устанавливается
- Есть overhead
- Подходит для http, электронной почты

---
class: top black
background-size: 35%
background-image: url(img/5_4.gif)

# TCP - диаграмма состояний

---

class: black
background-image: url(img/tcp-udp-otlichiya-5.png)

# TCP - установка соединения

---

class: black
background-size: 75%
background-image: url(img/TCP_CLOSE.svg)

# TCP - завершение соединения


---
# Пакет net
В Go за сетевые возможности отвечает пакет net и его подпакеты

---

# Dialer

Тип, задача которого установка соединений

Обладает следующим интерфейсом:

```
func (d *Dialer) Dial(network, address string) (Conn, error)

func (d *Dialer) DialContext(ctx context.Context,
                    network, address string) (Conn, error)
```

Можно использовать стандартные знаения параметров функцией
```
func Dial(network, address string) (Conn, error)
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
```

---

# Примеры установки соединений
```
Dial("tcp", "golang.org:http")
Dial("tcp", "192.0.2.1:http")
Dial("tcp", "198.51.100.1:80")
Dial("udp", "[2001:db8::1]:domain")
Dial("udp", "[fe80::1%lo0]:53")
Dial("tcp", ":80")
```

---

# Тип Conn

Является абстракцией над поточным сетевым соединением.

Является имплементацией Reader, Writer, Closer.

Это потокобезопасный тип.

---

# Практика
- Пишем чат сервер
- Учимся создавать многопоточные сервера

---

# Типичные сетевые проблемы

* Какие знаете?
* С какими сталкивались? Что делали?

---

# Типичные сетевые проблемы

- Потеря пакетов
- Недоступность
- Тайм-ауты
- Медленные соединения

Как с ними быть?

---
# Практика

Изучаем инструменты отладки:
- tcpdump
- wireshark
- lsof
- netstat

---

# Повторение

.left-text[
Давайте проверим, что вы узнали за этот урок, а над чем стоит еще поработать.
<br><br>
Ссылка в чате
]

.right-image[
![](../img/gopher_science.png)
]

---

# Примеры с занятия

https://github.com/OtusGolang/webinars_practical_part/tree/master/24-context-n-net

---

# Следующее занятие

## Работа с SQL

<br>
<br>
<br>

## 23 сентября, четверг

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
Ссылка в чате
]

.right-image[
![](../img/gopher_boat.png)
]

---

class: white
background-image: url(../img/message.svg)
.top.icon[![otus main](../img/logo.png)]

# Спасибо за внимание!
