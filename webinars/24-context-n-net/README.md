.center.icon[![otus main](https://drive.google.com/uc?id=1NPIi9Hw5ZjA5SK24lTXckDjNAPSuFAHi)]

---


class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Низкоуровневые протоколы
## TCP, UDP, DNS

### Иван Ремень

---

class: top white
background-image: url(tmp/sound.svg)
background-size: 130%
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

.sound-top[
  # Как меня слышно и видно?
]

.sound-bottom[
  ## > Напишите в чат
  ### **+** если все хорошо
  ### **–** если есть проблемы cо звуком или с видео
]

---

# Цель занятия 

.right-image[
![](tmp/gopher.png)
]

# 
- Изучить что такое контекст
- Изучить особенности протоколов TCP и UDP
- Изучить стандартные типы Conn и Dialer
- Узнать о типичных сетевых проблемах
- Научиться обеспечивать тайм-ауты
- Научиться отлаживать сетевые проблемы

---
# Контекст
Теперь это часть стандартной библиотеки
```
import "context"
```

Имеет природу матрешки: контексты вкладываются друг в друга
---

# Виды контекстов
```
func Background() Context
func TODO() Context
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key interface{}, val interface{}) Context
```
---

# Практика
Пишем эмулятор долгих операций

---

# UDP

- Доставка пакета не гарантируется
- Порядок сообщений не гарантируется
- Соединение не устанавливается
- Быстрый
- Подходит для потокового аудио и видео, статистики, игр

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
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

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

# DNS

- Служит для получение IP адреса по доменному имени (и не только)
- Работает как поверх UDP, так и поверх TCP
- Имеет рекурсивную природу
- Имеет механизмы для кеширования
- При высоких нагрузках можно использовать /etc/hosts

---
class: black
background-image: url(img/2880px-Example_of_an_iterative_DNS_resolver.svg.png)

# Рекурсивная природа DNS


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

- Потеря пакетов
- Недоступность
- Тайм-ауты
- Медленные соединения

Как с ними бороться?

---
# Практика

Изучаем инструменты отладки:
- tcpdump
- wireshark
- lsof
- netstat

---

# Тест

https://forms.gle/SiDmYTPUU5La3rA88

---


# На занятии

- Изучили что такое контекст
- Изучили особенности протоколов TCP и UDP
- Изучили стандартные типы Conn и Dialer
- Узнали о типичных сетевых проблемах
- Научились обеспечивать тайм-ауты
- Научились отлаживать сетевые проблемы

---

## Вопросы?

---

# Опрос

Не забудьте заполнить опрос. Ссылка на опрос будет в слаке.

---

class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Спасибо за внимание!