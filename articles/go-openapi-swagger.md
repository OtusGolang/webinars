# Использование Swagger в Go

В предыдущей статье я сделал короткое введение в спецификацию Open API и показали некоторые инструменты для работы с ней. Сейчас я более подробно рассмотрим go-swagger - утилиту для генерации Go кода из swagger файлов.

## go-swagger - кодогенрация для Go

[go-swagger](https://github.com/go-swagger/go-swagger) это инструмент для Go разработчиков, позволяющий автоматически генерировать Go код по swagger файлам. Он опирается на различные библиотеки из [проекта go-openapi](https://github.com/go-openapi) для работы с форматом swagger.

Я некоторое время следил за проектом. Он развивается очень активно, коммиты добавляются в master-ветку каждые несколько дней. Основные контрибьюторы очень быстро реагируют на возникающие проблемы. Проект ввыпускается релизами с конкретными версиями и поставляется в виде исполняемых файлов или docker-контейнеров.

## Пример

В первую очередь установить команду `swagger` [по инструкции с github](https://github.com/go-swagger/go-swagger/blob/master/docs/install.md). Далее мы будем использовать пример `swagger.yaml` файла из предыдущего поста.

## Создание REST сервера

Используйте следущие bash команды что бы создать и запустить сервер для вашего swagger файла. Единственные требование - это наличие swagger.yaml файла в текущей рабочей директории, и то что эта директория находится внутри `GOPATH`.

    $ # Validate the swagger file 
    $ swagger validate ./swagger.yaml
    The swagger spec at "./swagger.yaml" is valid against swagger specification 2.0
    $ # Generate server code
    $ swagger generate server
    $ # go get dependencies, alternatively you can use `dep init` or `dep ensure` to fix the dependencies.
    $ go get -u ./...
    $ # The structure of the generated code
    $ tree -L 1
    .
    ├── cmd
    ├── Makefile
    ├── models
    ├── restapi
    └── swagger.yaml
    $ # Run the server in a background process
    $ go run cmd/minimal-pet-store-example-server/main.go --port 8080 &
      09:40:12 Serving minimal pet store example at http://127.0.0.1:8080
    $ # go-swagger serves the swagger scheme on /swagger.json path:
    $ curl -s http://127.0.0.1:8080/swagger.json | head
      {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "schemes": [
          "http"
        ],
    $ # Test list pets
    $ curl -i http://127.0.0.1:8080/api/pets
    HTTP/1.1 501 Not Implemented
    Content-Type: application/json
    Content-Length: 50

    "operation pet.List has not yet been implemented"
    $ # Test enforcement of scheme - create a pet without a required property name.
    $ curl -i http://127.0.0.1:8080/api/pets \
        -H 'content-type: application/json' \
        -d '{"kind":"cat"}'
    HTTP/1.1 422 Unprocessable Entity
    Content-Type: application/json
    Content-Length: 49

    {"code":602,"message":"name in body is required"}

Все должно заработать автоматически! go-swagger создал следущие директории

* `cmd` - команда для запуска сервера, обработка параметров командной строки и конфигурации.
* `restapi` - логика маршрутизации запросов на основе секции `paths` в swagger файле.
* `models` - модели из раздела `definitions` в swagger файле.

## Создание REST клиента

Давайте попробуем создать клиент к нашему серверу

    $ swagger generate client

И напишем небольшую программу, которая использует сгенерированный клиент

    package main

    import (
        "context"
        "flag"
        "fmt"
        "log"

        "github.com/posener/swagger-example/client"
        "github.com/posener/swagger-example/client/pet"
    )

    var kind = flag.String("kind", "", "filter by kind")

    func main() {
        flag.Parse()
        c := client.Default
        params := &pet.ListParams{Context: context.Background()}
        if *kind != "" {
            params.Kind = kind
        }
        pets, err := c.Pet.List(params)
        if err != nil {
            log.Fatal(err)
        }
        for _, p := range pets.Payload {
            fmt.Printf("\t%d Kind=%v Name=%v\n", p.ID, p.Kind, *p.Name)
        }
    }

Когда мы запустим ее, мы получим ошибку HTTP 501:

    $ go run main.go 
      15:57:53 unknown error (status 501): {resp:0xc4204c2000}
    exit status 1

## Реализация метода API

Как вы можете видеть, сгенерированный сервер возвращает код ответ 501 (ошибка HTTP not implemented) для всех методов, которые мы определили в swagger файле. Реализация того что я называю "бизнес логикой" делается в файле `configure_minimal_pet_store_example.go`. Этот файл так же сгенерирован автоматически, но он особенный - он не будет перезаписан при следующем вызове команды `generate server`. Как предлагает go-swagger, вы можете и должны редактировать этот файл. Он создается только один раз при первом запуске `generate server`.

Для примера, давайте реализует метод `pet list`.

Для этого найдем в файле такой фрагмент кода:

    func configureAPI(api *operations.MinimalPetStoreExampleAPI) http.Handler {
        [...]
        api.PetListHandler = pet.ListHandlerFunc(func(params pet.ListParams) middleware.Responder {
            return middleware.NotImplemented("operation pet.List has not yet been implemented")
        })
        [...]
    }

Как и ожидалось, автоматически сгенерированный код просто возвращает `middleware.NotImplemented`, реализуя интерфейс `middleware.Responder`, который во многом похож на обычный `http.ResponseWriter` интерфейс:

    // Responder is an interface for types to implement
    // when they want to be considered for writing HTTP responses
    type Responder interface {
        WriteResponse(http.ResponseWriter, runtime.Producer)
    }
    
Для нашего удобства сгенерированный код содержит ответы для каждого метода, который мы определили в файле `swagger.yaml`. Давайте создадим фиксированный список животных для этого метода и будет фильтровать его по ключевому слову, переданному в параметрах запроса.

    var petList = []*models.Pet{
    {ID: 0, Name: swag.String("Bobby"), Kind: "dog"},
        {ID: 1, Name: swag.String("Lola"), Kind: "cat"},
        {ID: 2, Name: swag.String("Bella"), Kind: "dog"},
        {ID: 3, Name: swag.String("Maggie"), Kind: "cat"},
    }

    func configureAPI(api *operations.MinimalPetStoreExampleAPI) http.Handler {
        [...]
        api.PetListHandler = pet.ListHandlerFunc(func(params pet.ListParams) middleware.Responder {
            if params.Kind == nil {
                return pet.NewListOK().WithPayload(petList)
            }
            var pets []*models.Pet
            for _, pet := range petList {
                if *params.Kind == pet.Kind {
                    pets = append(pets, pet)
                }
            }
            return pet.NewListOK().WithPayload(pets)
        })
        [...]
    }

Перезапустим сервер и проверим его с помощью клиента

    $ go run main.go 
        0 Kind=dog Name=Bobby
        1 Kind=cat Name=Lola
        2 Kind=dog Name=Bella
        3 Kind=cat Name=Maggie
    $ go run main.go -kind=dog
        0 Kind=dog Name=Bobby
        2 Kind=dog Name=Bella
        

## Что можно улучшить

Ниже я приведу несколько вещей, котрые на мой взгляд нужно улучшить. Хочу подчеркнуть, что это только мое мнение и я считаю, что они могут сделать и так хороший go-swagger превосходным.

### Неудобный configure_\*.go

Файл `restapi/configure_*.go` весьма неудобен в использовании:
* Это автоматически генерируемый файл, но он генерируется только в первый раз - довольно странное поведение
* Когда API меняется/расширяется необходимо вручную его обновлять
* Все API методы управляются в одном файле.
* И последний, хотя весьма существенный недостаток: *невозможно использовать Dependency Injection что бы протестировать поведение API*

### Обязательные поля

В описании модели обязательные поля генрируется как указатели, а опциональные поля - как значения. Наприме, модель `Pet` из примера с обязательным полем `Name` и опциональными полями `Kind` и `ID` сгенерирована следующим образом (свойство `readonly` пока не поддерживается):

    type Pet struct {
        ID   int64   `json:"id,omitempty"`
        Kind string  `json:"kind,omitempty"`
        Name *string `json:"name"`
    }

Причина этого, насколько я понимаю - гарантировать, что обязательные поля на самом деле переданы. В Go SDK для Amazon Web Services использован похожий подход, но [для опциональных полей](https://github.com/aws/aws-sdk-go/issues/114). Но у этого подхода есть недостатки:

* Опциональные поля: если я получил модель `Pet` с полем `Kind == ""`, как мне понять - была ли это пустая строка или поле не было передано совсем ?

* Неудобно использовать: чтение и записей указателей может быть весьма утомляющим занятием и часто приходится писать вспомогательные функции. Пакет `swag`, например, как раз содержит такие.

### Сложно получить `http.Handler`

go-swagger генерирует законченный сервер, с функцией `main()`, разбором параметров командной строки и прочим, что позволяет очень быстро запустить сервис. Это действительно прикольно, но иногда бывает необходимо использовать свою функцию `main()` и свой фреймворк, который включает переменные окружения, логирование, трассировку и т.д. В такой ситуации в Go используется стандартный интерфейс `http.Handler`, и я ожидаю, что в сгенерированный код будет его предоставлять. Но в текущем дизайне go-swagger получить этот интерфейс не так то просто.

### Сложно использовать и настраивать сгенерированный клиент

Сгенерированный клиент работает "из коробки" - как было показано выше. Тем не менее, как-либо настроить или кастомизировать работу клиента сложно, в основном потому, что он использует нестандартные сущности. Например, в клиенте есть метод `SetTransport`, который принимает `runtime.ClientTransport` (специфичный для go-swagger тип). Установить новый транспорт с кастомизированным HTTP клиентом или своим URL - не простая задача.

Другая проблема в том, что клиент не предоставляет интерфейсы. Когда я разрабатываю пакет, использующий клиент, мне нужны интрефейсы, что подменить клиент на заглушку в юнит-тестах. Сгенерированный клиент таких интерфейсов не предоставляет.

### Не используется стандартный `time.Time`

По разным причинам, поля определенные как `type: string, format: date-time` являются значениями типа `strfmt.DateTime` из пакет ` github.com/go-openapi/strfmt`, а не стандартного `time.Time`. Это требует утомительных преобразований типа при работе с другими библиотеками, рассчитанными на `time.Time`

### Версионирование библиотек go-openapi

Библиотека go-openapi не версионируется и, к сожалению, у нее часто ломается обратная совместимость.