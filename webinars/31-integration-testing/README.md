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
]


---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Тестирование микросервисов

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем - задаем вопросы.
* ### Чат вижу - могу ответить не сразу.
* ### После занятия - оффтопик, ответы на любые вопросы.

---

# О чем будем говорить

.big-list[
* Юнит-тестирование vs интеграционное
* TDD
* Настройка окружения
* BDD
]

---

# Цель занятия

## Узнать об&nbsp;инструментах и&nbsp;подходах к&nbsp;интеграционному тестированию в&nbsp;Go.

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
![](img/gopher_science.png)
]

---

# Зачем?

### Зачем нужны тесты?

---

# Зачем?

* Упростить рефакторинг (safety net)

* Документировать код, определить контракт

* Отделить интерфейса от реализации (mocks)

* Найти найти неактуальный код

* Найти найти новые кейсы

* Повысить качество кода

---

# Думай, как тестировщик

* Как хотелось бы, чтобы работало? (На что это похоже? Как бы я мог это использовать?) Не лазить в кишки.

* Как не должно работать? (Неправильные параметры, неправильный порядок вызовов) Негативные тест-кейзы.

* Что там на краю обрыва? (Самое маленькое/большое число, граница, на которой меняется состояние). Граничные условия.

* А что если? Странные сценарии использования.

---

class: bottom
background-image: url(img/unit_vs_integr.png)
background-size: 95%

# Модульные VS Интеграционные

https://habr.com/ru/post/358950/
<br>
https://habr.com/ru/post/358178/


---

class: bottom
background-image: url(img/tdd.png)
background-size: 70%

# TDD (Test-Driven Development)

https://ru.wikipedia.org/wiki/Разработка_через_тестирование


---

# TDD: пример

## Задача
### Написать функцию Join, которая склеивает слайс рун в строку, игнорируя пробелы.

---

# TDD: исправление багов

## Ситуация

* Вам сообщили о проблеме и вы экстренно её решили.

* Написали тест на решение и он прошёл.

* Через какое-то время проблема повторилась.

## Почему?

---

class: bottom
background-image: url(img/unit_vs_integr.png)
background-size: 95%

# Интеграционные тесты

---

# Интеграционные тесты: окружение

## Варианты:
- Поднимаем сервисы, базу, кеши и пр. локально

- У нас есть виртуалка или админы любезно предоставили нам тестовое окружение, куда мы можем раскатиться

- **docker-compose**:
  - поднимаем всю инфраструктуру или только часть
  - не храним состояние между запусками, если не хотим

Что делать с сервисами, которые ходят во внешнюю сеть (стороннее API и пр.)?

---

# docker-compose: полезные команды

```
docker-compose [-f file] up [–d] [–build] [--exit-code-from service]
docker-compose [-f file] down
docker-compose logs [–f service]
docker-compose ps [–a]
docker-compose [-f file] run service [command]
docker-compose [-f file] exec service [command]
```

---

# Интеграционные тесты: пример

https://github.com/kulti/task-list

---

# Интеграционные тесты: резюме

- Используйте окружение максимально похожее на проду
- Не используйте тестируемый код
- Используйте спецификации и кодогенеренных клиентов
- Мокайте ненужные сервисы
- Пишите интеграционные тесты на том языке, на каком удобнее

---

class: bottom
background-image: url(img/bdd.png)
background-size: 80%

# Behavior-Driven Development (BDD)

https://en.wikipedia.org/wiki/Behavior-driven_development

---

# BDD

- Похож на TDD.

- Описание идёт через спецификацию поведения.

- Стандарт для спецификации de facto – язык Gherkin.

- Наиболее известная компания, продвигающая фреймворки для BDD - Cucumber.

- BDD придуман, чтобы бизнесу был ближе к программистам. (как на самом деле?)


---

# BDD: язык Gherkin

```
Feature: Guess the word

  # The first example has two steps
  Scenario: Maker starts a game
    When the Maker starts a game
    Then the Maker waits for a Breaker to join

  # The second example has three steps
  Scenario: Breaker joins a game
    Given the Maker has started a game with the word "silky"
    When the Breaker joins the Maker's game
    Then the Breaker must guess a word with 5 characters
```

https://cucumber.io/docs/gherkin/reference/


---

# BDD: возможный тест на Gherkin

```
История: Отсылка email-уведомления

  Как клиент API сервиса регистрации
  Чтобы понимать, что пользователю приходит подтверждение регистрации
  Я хочу получать события из соответствующей очереди

  Сценарий: Получаем событие от сервиса уведомлений
  Когда я отсылаю POST-запрос с пользовательским JSON в сервис регистрации
  Тогда ответ от сервиса должен быть 200 ОК
  И я должен получить событие из очереди, содержащее email-пользователя
```

---

# BDD: пример

1. Клиент API посылает запрос на регистрацию пользователя в **RegistrationService**

2. **RegistrationService** сохраняет пользователя в базу и публикует событие, что произошла новая регистрация

3. **NotificationService** уведомляет пользователя о регистрации (например смс, email и пр.) и публикует событие,
что такой-то пользователь был проинформирован.

---

class: bottom
background-image: url(img/example.png)
background-size: 90%

# BDD: пример

---

# BDD: реализация


Реализация примера:<br>
https://github.com/OtusGolang/webinars_practical_part/tree/master/31-integration-testing

<br><br>
Для BDD используем godog (читайте внимательно README):<br>
https://github.com/DATA-DOG/godog

---

# BDD: плюсы-минусы

## Плюсы
- Позволяет взглянуть со стороны (мышление тестировщика)
- Тулинг для тестов отличается от тулинга для продукта

<br>

## Минусы
- Сложно написать нетривиальные кейсы
- Неудобный рефакторинг

---

# Повторение

.left-text[
Давайте проверим, что вы узнали за этот урок, а над чем стоит еще поработать.
<br><br>
Ссылка в чате
]

.right-image[
![](img/gopher_science.png)
]

---

# Опрос

.left-text[
  Заполните пожалуйста опрос о занятии.
  <br><br>
  Ссылка в чате.
]

.right-image[
  ![](img/gopher7.png)
]


---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
