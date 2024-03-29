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
	+ если все хорошо
	- если есть проблемы со звуком или с видео]


---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Кодогенерация и дженерики в Go

### Владимир Балун

---

# Как проходит занятие

* ### Активно участвуем — задаем вопросы.
* ### Чат вижу — могу ответить не сразу.
* ### После занятия — оффтопик, ответы на любые вопросы.

---

# План занятия

* ### Поговорим о кодогенерации
* ### Посмотрим, где она может нам помочь
* ### Посмотрим на Protocol Buffers
* ### Посмотрим на Generics

---

# Кодогенерация

* ### Что это такое?

---

# Кодогенерация
* ### Зачем она нужна?
* ### Какие задачи помогает решить?

---

# Зачем нужна кодогененрация

* ### Генерировать код по метаописанию (swagger, protobuf, ...)
* ### Генерировать обобщенный код (до Go 1.18)
* ### Генерировать заглушки для интерфейсов
* ### Встраивать данные в код

<br><br>
Пример в стандартной библиотеке:<br>
https://golang.org/src/unicode/tables.go

---

# Go generate

```
//go:generate echo "Hello, world!"

package main

import (
	"fmt"
)

func main() {
	fmt.Println("run any unix command in go:generate")
}
```

```
> go generate
Hello, world!
```

https://github.com/golang/go/blob/master/src/cmd/go/internal/generate/generate.go

---

# Цикл разработки пакета с go generate

```
	% edit …
	% go generate
	% go test

	% git add *.go  # коммитим сгенерированный код
	% git commit
```

---

# Принципы go generate

[Go generate: A Proposal](https://docs.google.com/document/d/1V03LUfjSADDooDMhe-_K59EgpTEm3V8uvQRuNMAEnjg/edit)

<br>

* ### Запускать на машине разработчика пакета, а не пользователя.
  * утилиты для генерации нужны только разработчику
  * генерация не происходит автоматически при go get

* ### Добавлять disclaimer.

```
/*
* CODE GENERATED AUTOMATICALLY WITH tool name
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/
```
* ### Работать только с .go-файлами, как часть тулкита go.

---

# Go generate

псевдоним:
```
//go:generate -command foo go tool foo
```

regexp:
```
go generate -run enums
```

выводить команды:
```
go generate -x

```

список команд к выполнению:
```
go generate -n
```

---

# Binary data

```
go get -u github.com/go-bindata/go-bindata/...
```

```
go-bindata -o myfile.go data/
```

```
//go:generate bindata -o jpegs.go pic1.jpg pic2.jpg pic3.jpg
```


```
b, err := Asset("pic1.jpg")
if err != nil {
	log.Fatalf("unable to get template: %v", err)
}
```

Примеры:
- статика (изображения, иконки и пр.)
- транзакции
- скрипты
- ...

---

# Go embed

https://golang.org/pkg/embed/

```
//go:embed static/gopher.png
var gopherPngBytes []byte
```

---

# Stringer

```
go get golang.org/x/tools/cmd/stringer
```

```
func (t T) String() string
```

```
//go:generate stringer -type=MessageStatus
type MessageStatus int

const (
	Sent MessageStatus = iota
	Received
	Rejected
)
```

```
func main() {
	status := Sent
	fmt.Printf("Message is %s", status) // Message is Sent
}
```

---

# JSON Enums

```
go get github.com/campoy/jsonenums
```

```
func (t T) MarshalJSON() ([]byte, error)
func (t *T) UnmarshalJSON([]byte) error
```

```
//go:generate jsonenums -type=Status
type Status int

const (
	Pending Status = iota
	Sent
	Received
	Rejected
)
```

---

# Генерация Marshal/Unmarshal: easyjson


```
go get -u github.com/mailru/easyjson/...
```

```
easyjson -all <file>.go
```

 <br>
генерирует MarshalEasyJSON / UnmarshalEasyJSON, для структур из файла
<br>
кратно быстрее за счет отсутствия рефлексии

<br><br>


***
P.S. https://github.com/json-iterator/go

---

# Генерация go структур из JSON

https://mholt.github.io/json-to-go/

```
go get github.com/ChimeraCoder/gojson/gojson
```

```
{
  "name" : "Alex",
  "age": 24,
  "courses": ["go", "python"]
}
```

```
cat schema.json | gojson -name Person

package main

type Person struct {
        Age     int64    `json:"age"`
        Courses []string `json:"courses"`
        Name    string   `json:"name"`
}
```

---

# Реализация интерфейсов: impl

```
go get -u github.com/josharian/impl
```

```
$ impl 'f *File' io.ReadWriteCloser
func (f *File) Read(p []byte) (n int, err error) {
    panic("not implemented")
}

func (f *File) Write(p []byte) (n int, err error) {
    panic("not implemented")
}

func (f *File) Close() error {
    panic("not implemented")
}
```


---

# Моки интерфейсов: gomock

```
GO111MODULE=on go get github.com/golang/mock/mockgen@latest
```

```
//go:generate mockgen -source=$GOFILE
//-destination ./mocks/mock_getter.go -package mocks Getter
type Getter interface {
    Get(url string) (resp *http.Response, err error)
}
```

---

# Generics: дилема

```
The generic dilemma is this: do you want slow programmers,
slow compilers and bloated binaries, or slow execution times?
(c) Russ Cox
```

---

# Generics: какие есть варианты до Go 1.18?


- copy & paste (см. пакеты strings and bytes)
- интерфейсы

```
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

- type assertions
- рефлексия
- go generate

---

# Generics до Go 1.18

```
go get github.com/cheekybits/genny
```

```
//go:generate genny -in=$GOFILE -out=gen-$GOFILE
gen "KeyType=string,int ValueType=string,int"
```

объявляем заглушки по типам:

```
type KeyType generic.Type
type ValueType generic.Type
```

пишем обычный код:

---

# Generics после Go 1.18

```
func existsInSlice[T comparable](val T, values []T) bool {
	for _, v := range values {
		if val == v {
			return true
		}
	}

	return false
}

func main() {
	result := existsInSlice[int](10, []int{9, 10, 11})
	fmt.Println(result)
}
```

https://go.dev/play/p/IcZKGVLboVk

---

# Констрейты (ограничения типов)

У каждого параметра-типа обязательно указывается ограничение типа. 
Констрейт - это интерфейс, который описывает, каким может быть тип. 
Этот интерфейс может быть обычным go-интерфейсом:

```
type OwnConstraint interface {
	String() string
}
```

А может быть интерфейсом, перечисляющим полный список типов, для которых он может быть использован, 
а **использован он может быть только внутри дженериков:**

```
type OwnConstraint interface {
	int | int8 | int16 | int32 | int64
}
```

```
type OwnConstraint interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
```

---

# Мешать их нельзя

```
type OwnConstraintWithMethods interface {
	String() string
}

type OwnConstraint interface {
	int | int8 | int16 | int32 | int64 | OwnConstraintWithMethods
}
```

```
cannot use main.OwnConstraintWithMethods in union (main.OwnConstraintWithMethods contains methods)
```

---

# Готовые констрейнты

```
// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}

// comparable is an interface that is implemented by all comparable types
// (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
// The comparable interface may only be used as a type parameter constraint,
// not as the type of a variable.
type comparable interface{ comparable }
```

https://pkg.go.dev/golang.org/x/exp/constraints

```
type Complex
type Float
type Integer
type Ordered
type Signed
type Unsigned
```

---

# Cтруктуры

```
type ListNode[T any] struct {
	value T
	next  *ListNode[T]
}

type List[T any] struct {
	head *ListNode[T]
}

func (l *List[T]) Push(value T) {
	l.head = &ListNode[T]{value, l.head}
}

func (l *List[T]) Pop() T {
	if l.head == nil {
		panic("list is empty")
	}

	value := l.head.value
	l.head = l.head.next
	return value
}
```

https://go.dev/play/p/60EyJgE6nYl

---

#  Не все так хорошо с методами, к сожалению...

```
type Data[T1, T2 any] struct {
	value1 T1
	value2 T2
}

func (l *Data[T1, T2]) Print() {
	fmt.Println(l.value1, l.value2)
}

func (l *Data[T1, T2]) PrintWith[T any](value T) {
	fmt.Println(l.value1, l.value2, value)
}

func main() {
	d := Data[int, float32]{}
	d.Print()
}
```

https://go.dev/play/p/4ooRw0u1w8S

```
syntax error: method must have no type parameters
```

---

#  Но хоть можно так

```
type Data[T1, T2 any] struct {
	value1 T1
	value2 T2
}

func (l *Data[T1, T2]) Print() {
	fmt.Println(l.value1, l.value2)
}

func PrintWith[T1, T2, T any](d *Data[T1, T2], value T) {
	fmt.Println(d.value1, d.value2, value)
}

func main() {
	d := Data[int, float32]{}
	PrintWith[int, float32, string](&d, "test")
	d.Print()
}
```

https://go.dev/play/p/GKtb9K0rPVq

---

# Иногда достаточно ldflags

```
package main

import (
	"fmt"
)

var VersionString = "unset"

func main() {
	fmt.Println("Version:", VersionString)
}
```

```
go run -ldflags '-X main.VersionString=1.0' main.go
```

***
```
go help build
        -ldflags '[pattern=]arg list'
                arguments to pass on each go tool link invocation.
```

---

# Что посмотрели:

- встраивание даных в код
- Stringer: String() для целочисленных типов: golang.org/x/tools/cmd/stringer
- Marshal/Unmarhsal для Enums:  github.com/campoy/jsonenums
- генерация структур из JSON: github.com/ChimeraCoder/gojson/gojson
- easyjson для быстрой работы с JSON
- моки интерфейсов: github.com/josharian/impl
- generics до и после Go 1.18

больше примеров для вдохновения:

https://github.com/avelino/awesome-go#generation-and-generics

---

# Protocol buffers


xml:
```
<person>
  <name>Elliot</name>
  <age>24</age>
</person>
```

json:
```
{
  "name": "Elliot",
  "age": 24
}
```

protobuf:
```
[10 6 69 108 108 105 111 116 16 24]
```

---

# Protocol buffers: формат протокола

https://developers.google.com/protocol-buffers/docs/proto3

```
syntax = "proto3";

package demo;


message People {
    repeated Person person = 1;
}

message Person {
    string name = 1;
    repeated Address address = 2;
    repeated string mobile = 3;
    repeated string email = 4;
}

message Address {
    string street = 1;
    int32 number = 2;
}
```

---

# Protocol buffers: как установить

1) Скачиваем нужный релиз proto-компилятора, кладём в PATH
https://github.com/protocolbuffers/protobuf/releases/tag/v3.12.4
```
$ protoc --version
libprotoc 3.21.12
```

<br>
2) Ставим генератор Go-кода

```
$ go install google.golang.org/cmd/protoc-gen-go
$ protoc-gen-go --version
protoc-gen-go v1.25.0
```
(не путать с https://github.com/golang/protobuf)


<br>
https://developers.google.com/protocol-buffers/docs/reference/go-generated


---

# Protocol buffers: кодогенерация


```
//go:generate protoc -go_out=. file.proto
```

globbing не поддерживается:

```
//go:generate protoc -go_out=. file1.proto file2.proto
```


---

# Protocol buffers: кодогенерация

```
message Foo {}
```

```
type Foo struct {
}

// Reset sets the proto's state to default values.
func (m *Foo) Reset()         { *m = Foo{} }

// String returns a string representation of the proto.
func (m *Foo) String() string { return proto.CompactTextString(m) }

// ProtoMessage acts as a tag to make sure no one accidentally implements the
// proto.Message interface.
func (*Foo) ProtoMessage()    {}
```

---

# Protocol buffers: запись и чтение

```
course := &myotus.Course{
	Title:   "Golang",
	Teacher: []*myotus.Teacher{
		{Name: "Dmitry Smal", Id: 1}, 
		{Name: "Alexander Davydov", Id: 2}
	},
}

out, err := proto.Marshal(course)
if err != nil {
	log.Fatalln("Failed to encode", err)
}
```

```
otusdb := &myotus.Otus{}
if err := proto.Unmarshal(in, otusdb); err != nil {
	log.Fatalln("Failed to parse otus database:", err)
}
```


---

# Protocol buffers: типы данных
<br><br>
https://developers.google.com/protocol-buffers/docs/proto3<br><br>

Скаляры: default, float, int{32,64}, string, bytes

Поля: одиночные, repeated (порядок сохраняется), reserved (полезно для удаленных полей)

```
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}
```

enums, должны начинаться с 0 как default value
```
enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
}
```

Maps:
```
map<string, Project> projects = 3;
```

---

# Protocol buffers: еще типы

https://github.com/gogo/protobuf/tree/master/protobuf/google/protobuf

---

# Следующее занятие

## Файлы конфигурации и логирование

<br>
<br>
<br>

## 20 мая, четверг

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
Ссылка в чате.
]

.right-image[
![](img/gopher7.png)
]

---

# Домашнее задание

Валидатор структур

https://github.com/OtusGolang/home_work/tree/master/hw09_struct_validator

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
