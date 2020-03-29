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

# Очереди сообщений

### Александр Давыдов
### Антон Телышев

---

# План занятия

.big-list[
* Очереди сообщений
* Apache Kafka
* Событийно-ориентированная архитектура
* RabbitMQ
* Использование RabbitMQ
]


---

# Очередь сообщений / Message Broker

.main-image[
	![](img/messagebroker.png)	
]

* Yandex Message Queue (https://cloud.yandex.ru/services/message-queue)
* Amazon Web Services (AWS) Simple Queue Service (SQS)
* Apache ActiveMQ
* Apache Kafka
* Redis (pubsub)
* RabbitMQ

etc.


---

class: black
background-size: 50%
background-image: url(img/crawler.png)

# Пример: crawler (поисковый робот)


---


# Вернемся к проблемам микросервисов...

.main-image[
 ![img/hell.jpeg](img/hell.jpeg)
]

---


# Очереди сообщений

* Слабое связывание
* Масштабируемость
* Эластичность
* Отказоустойчивость
* Гарантированная доставка
* Гарантированный порядок доставки
* Буферизация
* Понимание потоков данных
* Асинхронная связь

---

# Kafka

* Распределённый программный брокер сообщений
* Написан на Java/Scala
* Придуман в Linkedin чтобы обрабатывать безумный объем данных
* Есть коммерческая поддержка (Confluent)
* Линейно масштабируемый
* С гаринтией упорядоченности
* Надежный (репликация)
* Высокодоступный (high availability)

---

# Kafka: что почитать

* Designing Event Driven Systems
<br>
http://www.benstopford.com/2018/04/27/book-designing-event-driven-systems/


* Kafka: The Definitive Guide
<br>
https://www.confluent.io/wp-content/uploads/confluent-kafka-definitive-guide-complete.pdf


---

background-size: 65%
background-image: url(img/herokukafka.png)
# Что такое Kafka

<br><br><br><br><br><br>
<br><br><br><br><br><br>
<br><br><br><br><br><br>
— Издатель (producer)

— Топик (topic), раздел (partition)

— Подписчики (consumer), группа (group)


---

class: black
background-size: 65%
background-image: url(img/kafka-apis.png)
# Kafka: core APIs
---

# Kafka: core APIs

- The Producer API allows an application to publish a stream of records to one or more Kafka topics


- The Consumer API allows an application to subscribe to one or more topics and process the stream of records produced to them


- The Streams API allows an application to act as a stream processor, consuming an input stream from one or more topics and producing an output stream to one or more output topics, effectively transforming the input streams to output streams


- The Connector API allows building and running reusable producers or consumers that connect Kafka topics to existing applications or data systems. For example, a connector to a relational database might capture every change to a table

***
<br>
- The Producer API - публикация записей в топики.

- The Consumer API - подпись на топики.

- The Streams API - манипуляции с потоками, конвертация из одного топика в другой

- The Connector API - коннекторы к сторонним системам


---

# Kafka: как устроен топик

.main-image[
 ![img/topic.png](img/producers.png)
]


---

# Kafka: сообщение

.main-image[
 ![img/consumer-groups.png](img/message.png)
]


---

# Kafka: партиции

.main-image[
 ![img/topic.png](img/topic.png)
]


---

# Kafka: партиции


.main-image[
 ![img/scalingpartitions1.png](img/scalingpartitions1.png)
]


---

# Kafka: партиции


.main-image[
 ![img/scalingpartitions1.png](img/scalingpartitions2.png)
]


---

background-size: 40%
background-image: url(img/consumergroup.png)

# Kafka: группа консьюмеров


---

# Kafka: сжатие лога

.main-image[
 ![img/consumer-groups.png](img/logcompaction.jpg)
]


---

# Kafka: по сути журнал коммитов
распределенный,
отказоустойчивый,
масштабируемый

.main-image[
 ![img/messagelog.png](img/messagelog.png)
]


---

# Kafka: гарантии

- Упорядоченность сообщений в партиции
- Консьюмер видит сообщения в порядке попадания их в лог
- Сохранение сообщений при N-1 падениях и replication factor N


---

# Kafka: гарантии доставки

- At-most-once delivery (“как максимум однократная доставка”). Сообщение не может быть доставлено больше одного раза. При этом сообщение может быть потеряно.


- At-least-once delivery (“как минимум однократная доставка”). Сообщение никогда не будет потеряно. При этом сообщение может быть доставлено более одного раза.


- Exactly-once delivery (“строго однократная доставка”). Святой грааль систем сообщений. Все сообщения доставляются строго единожды.


https://dzone.com/articles/interpreting-kafkas-exactly-once-semantics


---

# Kafka: для каких задач?

- Message broker (ActiveMQ / RabbitMQ)
- Трекинг активности в вебе (linkedin)
- Сбор метрик
- Агрегация логов
- Stream processing (Kafka Streams)
- Event sourcing (https://martinfowler.com/eaaDev/EventSourcing.html)
- Storage? https://www.confluent.io/blog/okay-store-data-apache-kafka/

---

# Kafka: Go драйверы

* https://github.com/confluentinc/confluent-kafka-go - отличная дока, librdkafka c lib
* https://github.com/Shopify/sarama - плохая дока, зато чистый go
* https://github.com/segmentio/kafka-go - отличный код, хорошая дока


---

# RabbitMQ

RabbitMQ — это распределенная система управления очередью сообщений

<br>
#### Advanced Message Queuing Protocol (AMQP)


.main-image[
	![img/amqp-about.png](img/amqp-about.png)
]


---

# RabbitMQ

- Сообщение (<b>message</b>) — единица передаваемых данных, основная его часть (содержание) никак не интерпретируется сервером, к сообщению могут быть присоединены структурированные заголовки.


- Точка обмена (<b>Exchange</b>) — в неё отправляются сообщения. Точка обмена распределяет сообщения в одну или несколько очередей. При этом в точке обмена сообщения не хранятся.


- Очередь (<b>queue</b>) — здесь хранятся сообщения до тех пор, пока не будут забраны клиентом. Клиент всегда забирает сообщения из одной или нескольких очередей.


- Связки (<b>bindings</b>) - правила для роутинга сообщений


---


background-size: 60%
background-image: url(img/exchangefanout.webp)
# RabbitMQ: fanout exchange

<br><br><br><br><br>
<br><br><br><br><br>
<br><br><br><br><br>
<br><br><br><br><br>

Fanout – полностью игнорирует ключи маршрутизации и отправляет сообщения во все привязанные очереди. Точки обмена этого типа используются для распространения сообщений нескольким клиентам (рассылки уведомлений, обновлений, конфигураций и т.п.).


---

background-size: 60%
background-image: url(img/exchange-direct.png)
# RabbitMQ: direct exchange

<br><br><br><br><br>
<br><br><br><br><br>
<br><br><br><br><br>
<br><br><br><br><br>

Direct – доставляет сообщения в очереди по ключам маршрутизации. Ключи маршрутизации – это дополнительные данные, которые определяют, в какую очередь нужно отправить сообщение. Обычно точки обмена такого типа используются в балансировке нагрузки round-robin.


---

background-size: 90%
background-image: url(img/topicexchange.png)
# RabbitMQ: topic exchange

<br><br><br><br><br>
<br><br><br><br><br>
<br><br><br><br><br>
<br><br><br><br><br>

Topic – используется в шаблонах pub/sub. В этом случае ключ маршрутизации используется вместе с привязкой очередей к точке обмена. например, app.notification.sms.# — в очередь будут доставлены все сообщения, отправленные с ключами, начинающимися с app.notification.sms.


---

# RabbitMQ vs Kafka

- https://jack-vanlightly.com/blog/2017/12/4/rabbitmq-vs-kafka-part-1-messaging-topologies


- https://content.pivotal.io/blog/understanding-when-to-use-rabbitmq-or-apache-kafka

---


# Событийно-ориентированная архитектура

.main-image[
 ![](img/gopher8.png)
]


---

# Паттерны: Event Notification

.main-image[
 ![img/hell.jpeg](img/notificationreq.png)
]

request-driven

---

# Паттерны: Event Notification

.main-image[
 ![img/hell.jpeg](img/notification.png)
]

---

# Паттерны: State Transfer

.main-image[
 ![img/statetransfer.jpeg](img/statetransfer.png)
]


---

# Паттерны: Event Collaboration

.main-image[
 ![img/statetransfer.jpeg](img/eventcollaboration.png)
]


---

# Паттерны: Event Sourcing

.main-image[
 ![img/eventsourcing.jpeg](img/eventsourcing.jpeg)
]


---

# Паттерн: Команды, События, Запросы

.main-image[
 ![img/eventsourcing.jpeg](img/commandsqueries.png)
]


---

# Паттерны: CQRS
Command Query Responsibility Segregation

.main-image[
 ![img/cqsp.png](img/cqsp.png)
]


---

# Ссылки

- Разработка транзакционных микросервисов с помощью Агрегатов, Event Sourcing и CQRS<br>
https://habr.com/ru/company/nix/blog/322214/


- Основы CQRS<br>
https://habr.com/ru/company/simbirsoft/blog/329970/


---

# Ссылки

- RabbitMQ против Kafka: два разных подхода к обмену сообщениями<br>
https://habr.com/ru/company/itsumma/blog/416629/


- Understanding When to use RabbitMQ or Apache Kafka<br>
https://content.pivotal.io/blog/understanding-when-to-use-rabbitmq-or-apache-kafka


- Apache Kafka: обзор<br>
http://habr.com/ru/company/piter/blog/352978/


- Kafka и микросервисы: обзор<br>
https://habr.com/ru/company/avito/blog/465315/


- Apache Kafka и миллионы сообщений в секунду<br>
https://habr.com/ru/company/tinkoff/blog/342892/


- Apache Kafka и RabbitMQ: семантика и гарантия доставки сообщений<br>
https://habr.com/ru/company/itsumma/blog/437446/


---

# Домашнее задание 

Реализовать "напоминания" о событиях с помощью RabbitMQ:

<br>

- создать процесс, который периодически сканирует основную базу данных, выбирая события о которых нужно напомнить;


- создать процесс, который читает сообщения из очереди и шлет уведомления.

https://github.com/rabbitmq/rabbitmq-tutorials/tree/master/go

```
$ docker run -d --name rb -p 15672:15672 -p 5672:5672 rabbitmq:3-management
$ docker ps
```

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
[https://otus.ru/polls/4915/](https://otus.ru/polls/4915/)
]

.right-image[
![](img/gopher7.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
