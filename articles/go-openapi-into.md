Перевод https://posener.github.io/openapi-intro/

В следующей паре статей мы обсудим подход "API first", спецификацию Open API, а так же утилиту go-swagger, упрощающую создание REST API.

# OpenAPI

OpenAPI, изначально известное как Swagger это DSL (Domain Specific Language, специализированный язык) для описания REST API. Спецификации Open API могут быть описанны в виде JSON или YAML документов. Последняя версия языка swagger на данный момент - 3, разработана группой Open API Initiative, коллаборативным open-source проектом, входящим в Linux Foundation.

Основную документацию вы можете найти на сайте [swagger.io](https://swagger.io/). Сейчас существует большое количество инструментов и экосистем, разработанных вокруг Open API, отличная документация и, насколько я понимаю, Open API - это де-факто стандарт описания REST API (хотя есть альтернативы, см ниже).

Разработка REST сервиса с открытым API требует определенной методологии, хорошей методологии. В первую очередь определяется API сервиса, и только потом вокруг него создается все остальное. Есть множество инструментов что бы помочь вам сделать это правильно, которые будут описаны ниже. Условно инструменты можно разделить на:
* Редакторы - для создания/изменения swagger файлов
* Генераторы клиентов к API на любом языке, с валидацией параметров запроса
* Генераторы серверов на любом языке, с автоматической валидацией входящих запросов и исходящих ответов.
* Серверы-заглушки (mock) для тестирования.
* Генераторы клиентов для командной строки (CLI)

# Аспекты разработки сервисов

Обычно есть три заинтересованные стороны при разработке сервиса: архитекторы, разработчики сервиса и его потребители. Когда в процессе участвует множество сторон, неизбежно каждый из участников представляет себе сервис по-разному и предъявляет к нему разные требования. Open API может помочь сделать разработку сервисов масштабируемой, обеспечивая:

* Сотрудничество - все участники полностью понимаю что они разрабатывают. Они понимают что именно необходимо сделать, как это будет интегрировано и как будет работать.

* Согласованность - программы-клиенты созданные по swagger файлу будут знать как работать с любым сервером созданным по тому же файлу.

* Параллельная разработка - используя утилиты и swagger спецификацию, разработчик и потребитель сервиса могут разрабатывать параллельно, даже тестировать свою разработку независимо друг от друга и непосредственно взаимодействовать только на этапе интеграции для окончательной проверки своих реализаций. Разработчик сервиса может сконцентрироваться на бизнес-логике, используя сгенерированные клиенты и CLI для тестирования. Потребитель сервиса может использовать сгенерированные клиент и Mock-сервер для тестирования со своей стороны.

* Легкость изменения - API может быть легко изменено за счет авто-генерации кода, отвечающего за протокол.

* Прозрачность - сервис имеет четко определенный API, которая может быть легко задокументировано для потребителей.

* Независимость от языка - сервер и клиент могут быть реализованы на различных языках программирования.

Если в вашем процессе разработке есть необходимость любого из этих свойст - вам стоит попробовать OpenAPI.

# Альтернативы

Существует [RAML](https://raml.org/), который, так же как и Swagger является DSL для описания REST API. Если же вы не ограничены в использовании REST, то неплохой альтернативой является [gRPC](https://grpc.io/), которое использует protobuf в качеcтве языка описания API.

# Прежде чем начать

Требуется некоторое время что бы разобраться со swagger спецификацией, но после некоторых усилий понимать и использовать её становится очень просто.

Формальная спецификация огромна, в ней есть множество трюков и исключений из правил, а один и тот же результат можно получить разными способами. С моей точки зрения лучшей способ разобратьс в ней это:

1. Разобраться с инструментами и их ограничениями. Начните использовать инструменты и посмотрите насколько они соответствуют спецификации Open API (не все инструменты совместимы со всеми нюансами и версиями спецификации)

2. Подумайте над балансом между читаемостью swagger файлов и удобством использования сгенерированного кода.

3. Прийдите к "соглашениям" в использовании спецификации, согласуйте их со своими колегами и внедрите в организации.

# Пример

Давайте рассмотрим [пример с магазином животных](https://editor.swagger.io/) c сайте [swagger.io](https://swagger.io/). Удивительно, но он не "идеальный", в нем есть некоторые "не-RESTful" методы, например:

* Названия сущностей в путях должны быть во множественном числе, а в примере - в единственном, т.е. должно быть `/pets/` вместо `/pet`

* Присутствуют методы, которые на самом деле должны быть просто параметрами запроса для других методов, например вместо `/pet/findByStatus` должно быть `/pets?status=<status>`

* Для обновления деталей по животному используется метод `PUT /pet`, а ID передается в теле запроса. В то время как стандартный способ - `PUT /pets/{petID}`. 

* Операция создания объекта должно возвращать созданный объект, в то время как в примере ничего не возвращается.

Тем не менее пример дает хорошее понимание того как создать swagger файл.

Ниже очень короткий пример API магазина, включая три метода: pet-list, pet-create, pet-get и описание объекта pet. Пример довольно понятный поэтому я не буду описывать детально создать такой файл.

    swagger: '2.0'
    info:
      version: '1.0.0'
      title: Minimal Pet Store Example
    schemes: [http]
    host: example.org
    basePath: /api
    consumes: [application/json]
    produces: [application/json]
    paths:
      /pets:
        post:
          tags: [pet]
          operationId: Create
          parameters:
          - in: body
            name: pet
            required: true
            schema:
              $ref: '#/definitions/Pet'
          responses:
            201:
              description: Pet Created
              schema:
                $ref: '#/definitions/Pet'
            400:
              description: Bad Request
        get:
          tags: [pet]
          operationId: List
          parameters:
          - in: query
            name: kind
            type: string
          responses:
            200:
              description: 'Pet list'
              schema:
                type: array
                items:
                    $ref: '#/definitions/Pet'
      /pets/{petId}:
        get:
          tags: [pet]
          operationId: Get
          parameters:
          - name: petId
            in: path
            required: true
            type: integer
            format: int64
          responses:
            200:
              description: Pet get
              schema:
                $ref: '#/definitions/Pet'
            400:
              description: Bad Request
            404:
              description: Pet Not Found

    definitions:
      Pet:
        type: object
        required:
        - name
        properties:
          id:
            type: integer
            format: int64
            readOnly: true
          kind:
            type: string
            example: dog
          name:
            type: string
            example: Bobby
            
# Инструменты

Ниже список некоторых инструментов, которые применяются в Stratoscale, вы можете найти некоторые из них полезными

## Редактирование Swagger файлов

* [Swagger Editor](https://editor.swagger.io/) - Онлайн редактор

* [Jetbrain’s Swagger Editor](https://plugins.jetbrains.com/plugin/8347-swagger-plugin) - действительно хороший редактор swagger файлов - предоставляет автодополнение и создает Web UI для API.

## Генерация кода - клиент или сервер

* [Swagger Codegen](https://swagger.io/tools/swagger-codegen) - генератор кода на разных языках

* [go-swagger](https://github.com/go-swagger/go-swagger) - генерирует код только на Go, рассмотрим детальнее в следующей статье.

## Сервера - заглушки (Mock)

* [prism](https://github.com/stoplightio/prism) - возвращает случайные ответы и проверяет схему. Не сохраняется состояние - созданные объект `pet` не появится в списке `/pets`

* [imposter-openapi](https://github.com/outofcoffee/imposter) - еще один Mock сервер, для которого нужны образцы валидных ответов. В противном случае он может возвращать только 404 или 200 в зависимости от того найден или нет путь.

## Автоматическая генерация CLI

[open-cli](https://github.com/sharbov/open-cli) - это утилита автогенерации CLI для Open API серверов. Большинство серверов, сгенерированных из swagger файла предоставляют метод `GET /swagger.json`, который возвращает тот самый swagger файл из которого сгенрирован серевер. Утилита open-cli использует метод предоставляе сервером для построения CLI. Если сервер не предоставляет такой метод, но у вас есть swagger файл локально, то утилита так же может использовать его.


В следующей статье я более детально рассмотрю [go-swagger](https://github.com/go-swagger/go-swagger) - утилиту для генерации Go кода из swagger файлов. 
