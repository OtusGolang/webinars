background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Интерфейсы изнутри]
.tutor[Алексей Семушкин]
.tutor_desc[Software engineer at Semrush]

---

background-image: url(../img/rules.svg)

---

# О чем будем говорить

- Определение и реализация интерфейсов
- Композиция интерфейсов
- Пустой интерфейс
- Внутреннее устройство интерфейсов
- Интерфейсы и производительность программы
- Значение типа интерфейс и ошибки, связанные с nil
- Правила присваивания значений переменным типа интерфейс
- Опасное и безопасное приведение типов (type cast)
- Использование switch в контексте интерфейсов

---

# Интерфейсы: определение

**Интерфейс** — набор методов, которые надо реализовать, чтобы удовлетворить интерфейсу. Ключевое слово: `interface`.

```
type Stringer interface {
    String() string
}

type Shape interface {
    Area() float64
    Perimeter() float64
}
```

- Одному интерфейсу могут соответствовать много типов
- Тип может реализовать несколько интерфейсов

---

# Интерфейсы и типы

Переменная **типа интерфейс** может содержать значение типа, реализующего этот интерфейс.

```
var s Stringer // статический тип
s = time.Time{} // динамический тип
```

https://go.dev/doc/effective_go#interfaces_and_types

---

# Интерфейсы и типы

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

# Интерфейсы и типы
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

# Интерфейсы реализуются неявно

```
type Duck interface {
    Talk() string
    Walk()
    Swim()
}

type Dog struct {
    name string
}

func (d Dog) Talk() string {
    return "AGGGRRRR"
}

func (d Dog) Walk() { }

func (d Dog) Swim() { }

```

https://goplay.tools/snippet/GWYHjaDPnLG

---

# Интерфейсы реализуются неявно

```
type MyVeryOwnStringer struct { s string}

func (s MyVeryOwnStringer) String() string {
    return "my string representation of MyVeryOwnStringer"
}


func main() {
    // my string representation of MyVeryOwnStringer{}
    fmt.Println(MyVeryOwnStringer{"hello"})
}
```

```
type Stringer interface {
    String() string
}
```

https://goplay.tools/snippet/ppTH6Ya-fX5

---

# Интерфейсы: композиция

```
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

type ReadCloser interface {
    Reader
    Closer
}
```

---

# Интерфейсы: композиция

```
import "fmt"

type Greeter interface {
     hello()
}

type Stranger interface {
    Bye() string
    Greeter
    fmt.Stringer
}
```

---

# Интерфейсы: имена методов

Имена методов не должны повторяться:

```
type Hound interface {
    destroy()
    bark(int)
}

type Retriever interface {
    Hound
    bark() // duplicate method
}

```

```
./prog.go:6:2: duplicate method bark
```

https://goplay.tools/snippet/wMw2VKOIysx


---

# Интерфейсы: any

Пустой интерфейс не содержит методов:

```
type Any interface{}
```

```
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
   ...
}
```

---

# interface{} is says nothing

https://go-proverbs.github.io/

---

# Интерфейсы изнутри

```
type Speaker interface {
    SayHello()
}

type Human struct {
    Greeting string
}

func (h Human) SayHello() {
    fmt.Println(h.Greeting)
}

...

var s Speaker
h := Human{Greeting: "Hello"}
s = Speaker(h)
s.SayHello()

```

https://goplay.tools/snippet/rjboVEC3V6w

---

background-size: 60%
background-image: url(img/internalinterfaces.png)

# Интерфейсы изнутри

---

# Интерфейсы изнутри: iface

```
type iface struct {
    tab  *itab // Информация об интерфейсе
    data unsafe.Pointer // Хранимые данные
}
```

```
// itab содержит тип интерфейса и информацию о хранимом типе.
type itab struct {
    inter *interfacetype // Метаданные интерфейса
    _type *_type // Go-шный тип хранимого интерфейсом значения
    hash  uint32
    _     [4]byte
    fun   [1]uintptr // Список методов типа, удовлетворяющих интерфейсу
}
```

https://github.com/teh-cmc/go-internals/blob/master/chapter2_interfaces/README.md#anatomy-of-an-interface
<br><br>

---

# Интерфейсы изнутри

На этапе компиляции:
- генерируются метаданные для каждого статического типа, включая его список методов
- генерируются метаданные для каждого интерфейса, включая его
список методов

И при компиляции и в рантайме в зависимости от выражения:
- сравниваются methodset'ы типа и интерфейса
- создается и кэшируется `itab`

```
// Создание интерфейса:
// - аллокация места для хранения адреса ресивера
// - получение itab:
//      - проверка кэша
//      - нахождение реализаций методов
// - создание iface: runtime.convT2I
s := Speaker(Human{Greeting: "Hello"})

// Динамический диспатчинг
// - для рантайма это вызов n-го метода s.Method_0()
// - превращается в вызов вида s.itab.fun[0](s.data)
s.SayHello()
```

---

background-size: 60%
background-image: url(img/emptyinterface.png)

# Интерфейсы изнутри: any

---

# Интерфейсы изнутри: benchmark

```

type Addifier interface {
    Add(a, b int32) int32
}

type Adder struct { id int32 }

func (adder Adder) Add(a, b int32) int32 {
    return a + b
}

func BenchmarkDirect(b *testing.B) {
    adder := Adder{id: 6754}
    for i := 0; i < b.N; i++ {
        adder.Add(10, 32)
    }
}

func BenchmarkInterface(b *testing.B) {
    adder := Addifier(Adder{id: 6754})
    for i := 0; i < b.N; i++ {
        adder.Add(10, 32)
    }
}
```


---

# Интерфейсы изнутри: benchmark

```
BenchmarkDirect-16      1000000000   0.2436 ns/op   0 B/op   0 allocs/op
BenchmarkInterface-16   957668390    1.157 ns/op    0 B/op   0 allocs/op
```

```
$ GOOS=linux GOARCH=amd64 go tool compile -m addifier.go

Addifier(adder) escapes to heap
```

---

# Интерфейсы: еще раз о ресиверах

https://goplay.tools/snippet/jm1bKNLABnB
<br><br>
https://stackoverflow.com/a/45653986
<br><br>
https://stackoverflow.com/a/48874650

---

# Практика

Реализовать интерфейс Adult
<br><br>
https://goplay.tools/snippet/A48l0-8FQX0

---

# Zero-value

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

# Практика

<br>
Необходимо реализовать функцию `processMessage`.

<br>

https://goplay.tools/snippet/EZ2pXx3DDKA

---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[25 августа]
.announce_topic[Горутины и каналы]

---
background-image: url(../img/thanks.svg)

.tutor[Алексей Семушкин]
.tutor_desc[Software engineer at Semrush]
