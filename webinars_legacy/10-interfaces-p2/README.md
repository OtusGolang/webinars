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

# Интерфейсы в Go. <br>Часть 2

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем — задаем вопросы.
* ### Чат вижу — могу ответить не сразу.
* ### После занятия — оффтопик, ответы на любые вопросы.

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

# О чем будем говорить:

* Значение типа интерфейс и ошибки, связанные с nil
* Правила присваивания значений переменным типа интерфейс
* Опасное и безопасное приведение типов (type cast)
* Использование switch в контексте интерфейсов
* Реализация подхода обобщенного программирования (generics) через интерфейсы

---

# Вспоминаем прошлое занятие

### Что такое интерфейс?

---

# Вспоминаем прошлое занятие

Интерфейс:
 - это набор сигнатур методов
 - методы реализуются неявно

Свойства:
 - интерфейсы могут встраивать другие интерфейсы
 - имена методов не должны повторяться
 - интерфейс может быть пустым (не иметь методов `interface{}`), такому интерфейсу удовлетворяет любой тип

---

# Значение типа интерфейс

<br>Значение типа интерфейс состоит из динамического типа и значения.
<br>Мы можем их смотреть при помощи %v и %T

```
type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}


func main() {
	var x fmt.Stringer
	x = Temp(24)
	fmt.Printf("%v %T\n", x, x) // 24 °C main.Temp
}
```
https://goplay.tools/snippet/JjXQsIsXwac


---

# Статический и динамический типы
<br>
...или с помощью пакета reflect

```
package main

import (
    "fmt"
    "reflect"
)

type InternalErr struct{}

func (e InternalErr) Error() string {
    return "500 Internal Server Error"
}

func main() {
    var e error
    e = InternalErr{}

    fmt.Println(reflect.TypeOf(e).Name()) // InternalErr
    fmt.Printf("%T\n", e)                 // main.InternalErr
}
```
https://goplay.tools/snippet/lr6oueIJ4zF


---

# Значение типа интерфейс

### Какое zero-value для интерфейса?

---

# Значение типа интерфейс

nil — нулевое значение для интерфейсного типа

```
type IHTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}

func main() {
    var c IHTTPClient
    fmt.Println("value of client is", c)
    fmt.Printf("type of client is %T\n", c)
	fmt.Println("(c == nil) is", c == nil)
}
```
https://goplay.tools/snippet/uBwvZ4bLy7T

---

# Значение типа интерфейс
<br>

```
type IHTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}

func main() {
    var h *http.Client
    var c IHTTPClient = h
    fmt.Println("value of client is", c)
    fmt.Printf("type of client is %T\n", c)
	fmt.Println("(c == nil) is", c == nil)
}
```
https://goplay.tools/snippet/ZzDLxREAXV2

---

# Интерфейсы: опасный nil
<br>
Что выведет программа?

```
func ReadFile(fname string) error {
    var err *os.PathError // nil

    if len(fname) == 0 {
        return err
    }

    // Do some work...
    return err
}

func main() {
    if err := ReadFile(""); err != nil {
        log.Printf("ERR: (%T, %v)", err, err)
    } else {
        log.Println("OK")
    }
}
```
https://goplay.tools/snippet/AUJ57LjntXb

---

#  Интерфейсы: опасный nil

<br>
Значение интерфейсного типа равно `nil` тогда и только тогда, когда `nil` и тип, и значение.

<br>

http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#nil_in_nil_in_vals

---

# Правила присваиваний (assignability rules)
<br>
Если переменная реализует интерфейс `I`, то мы можем присвоить ее переменной типа интерфейс `I`.

```
type BaseStorage interface {
    Close()
}

type UsersStorage struct{}
func (UsersStorage) Close() {}

type TicketsStorage struct{}
func (TicketsStorage) Close()      {}
func (TicketsStorage) GetTickets() {}

func main() {
    var s BaseStorage

    s = UsersStorage{}
    s = TicketsStorage{}
    _ = s
}
```
https://goplay.tools/snippet/jccNcScVWMZ

https://medium.com/golangspec/assignability-in-go-27805bcd5874

---

# Интерфейсы: присваивание

<br>

```
type MetricCollector interface {
    Record()
}

type AudioRecorder interface {
    Record()
}

type DummyRecorder struct{}
func (DummyRecorder) Record() {}

func main() {
    var v1 MetricCollector = DummyRecorder{}
    var v2 AudioRecorder = v1
    _ = v2
}
```

<br> Валидно?

<br>

https://goplay.tools/snippet/cG0FfsygGnC

---

# Интерфейсы: присваивание

```
type MetricCollector interface {
    Record()
}

type AudioRecorder interface {
    Record()
    Play()
}

type DummyRecorder struct{}
func (DummyRecorder) Record() {}

func main() {
    var v1 MetricCollector = DummyRecorder{}
    var v2 AudioRecorder = v1
    _ = v2
}
```

<br> Валидно?

<br>

https://goplay.tools/snippet/KH5yf0PlPkJ

---

# Интерфейсы: присваивание

Что, если мы хотим `присвоить переменной` конкретного типа `значение типа интерфейс`?


```
type MetricCollector interface {
    Record()
}

type DummyRecorder struct{}
func (DummyRecorder) Record() {}

func main() {
    var v1 MetricCollector
    var v2 DummyRecorder = v1
    _ = v2
}
```

<br>

https://goplay.tools/snippet/SYe5kK0nz-5

---

# Интерфейсы: type assertion


Выражение `x.(T)` проверяет, что интерфейс `x != nil` и конкретная часть `x` имеет тип `T`:

	- если T не интерфейс, то проверяем, что динамический тип x это T
	- если T интерфейс: то проверяем, что динамический тип x его реализует
---

# Интерфейсы: type assertion

<br>

```
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	r, ok := i.(fmt.Stringer)
	fmt.Println(r, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
```

<br>

https://goplay.tools/snippet/x-NbzVMZMUp

<br>

---

# Интерфейсы: type assertion

```
    var i interface{} = "hello"

	f, ok := i.(float64) // 0 false
	fmt.Println(f, ok)

	f = i.(float64) // panic: interface conversion:
					// interface {} is string, not float64
	fmt.Println(f)
```

Проверка типа возможна только для интерфейса:

```
	s := 5
    // Invalid type assertion: s.(int) (non-interface type int on left)
	i := s.(int)
```

---

# Интерфейсы: type assertion линтеры

https://golangci-lint.run/usage/configuration/
```yml
linters:
  enable:
    - errcheck
    - forcetypeassert

linters-settings:
  errcheck:
    check-type-assertions: true
```

---


# Интерфейсы: type switch

<br>
Мы можем объединить проверку нескольких типов в один `type switch`:

```
func checkSignature(/* ... */, publicKey crypto.PublicKey) (err error) {
    // ...

    switch pub := publicKey.(type) {
    case *rsa.PublicKey:
        // ...
    case *ecdsa.PublicKey:
        // ...
    case ed25519.PublicKey:
        // ...
    }
    return ErrUnsupportedAlgorithm
}
```
[src/crypto/x509/x509.go](https://github.com/golang/go/blob/283d8a3d53ac1c7e1d7e297497480bf0071b6300/src/crypto/x509/x509.go#L837)

---


# Интерфейсы: type switch


Как и в обычном `switch` мы можем объединять типы:

```
    case *rsa.PublicKey, *ecdsa.PublicKey:
        // Do some work...
    }
```

и обрабатывать `default`:

```
switch publicKey.(type) {
default:
    // No case for input type...
}
```

---


# Интерфейсы: type assertion & type switch

<br>
Заглянем в пакет fmt:

```
func (p *pp) printArg(arg interface{}, verb rune) {
    // ...

    switch f := arg.(type) {
    case bool:
        p.fmtBool(f, verb)
    case float32:
        p.fmtFloat(float64(f), 32, verb)
    case float64:
        p.fmtFloat(f, 64, verb)
    case complex64:
        p.fmtComplex(complex128(f), 64, verb)
    case complex128:
        p.fmtComplex(f, 128, verb)
    case int:
        p.fmtInteger(uint64(f), signed, verb)
    // ...
```
[src/fmt/print.go](https://github.com/golang/go/blob/283d8a3d53ac1c7e1d7e297497480bf0071b6300/src/fmt/print.go#L660)

---

# Немного практики

<br>
Необходимо реализовать функцию `processMessage`.

<br>

https://goplay.tools/snippet/EZ2pXx3DDKA

---

# Интерфейсы: приведение друг к другу


```
type BaseStorage interface {
    Close()
}

type SyncStorage interface {
    Close()
    Sync()
}

func main() {
    var s BaseStorage
    _ = SyncStorage(s)
}
```

Валидно? <br>
А наоборот? <br><br>

https://goplay.tools/snippet/Olph29QStlp

---


# Интерфейсы: почти дженерики

Дженерики на уровне языка: `map`, `slice`, etc.

<br>

Дженерики на уровне пользователя:
- https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md
- https://github.com/golang/go/issues/43651#issuecomment-776944155


---

# Интерфейсы: почти дженерики


Для реализации общих алгоритмов мы можем воспользоваться интерфейсами (или кодогенерацией):

```
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int

    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool

    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
```

---

# Интерфейсы: почти дженерики

```
type Person struct {
    Name string
    Age  int
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ...

people := []Person{
    {"Bob", 31},
    {"John", 42},
    {"Michael", 17},
    {"Jenny", 26},
}

sort.Sort(ByAge(people))
```

<br>

https://goplay.tools/snippet/SHZXfLu-ulF

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

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
Ссылка в чате.
]

.right-image[
![](../img/gopher_boat.png)
]

---

# Домашнее задание

Реализовать LRU-кэш на основе двусвязного списка
<br><br>
https://github.com/OtusGolang/home_work/tree/master/hw04_lru_cache

---

# Следующее занятие

## Обработка ошибок. Понятие паники

<br>
<br>
<br>

## 16 декабря, четверг

---

class: white
background-image: url(../img/message.svg)
.top.icon[![otus main](../img/logo.png)]

# Спасибо за внимание!
