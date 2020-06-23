.center.icon[![otus main](https://drive.google.com/uc?id=1NPIi9Hw5ZjA5SK24lTXckDjNAPSuFAHi)]

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

# Интерфейсы в Go. <br>Часть 1

### Антон Телышев


---

# О чем будем говорить

* Определение и реализация интерфейсов
* Композиция интерфейсов
* Пустой интерфейс
* Внутреннее устройство интерфейсов
* Интерфейсы и производительность программы


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

Переменная **типа интерфейс** может содержать значение типа, реализующего этот интерфейс.

```
var s Stringer
s = time.Time{}
```

https://golang.org/doc/effective_go.html#interfaces_and_types

---

# Интерфейсы реализуются неявно

<br> Dog удовлетворяет интерфейсу Duck

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

https://goplay.space/#GWYHjaDPnLG

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

https://goplay.space/#ppTH6Ya-fX5

https://golang.org/src/fmt/print.go#611


---

# Тип может реализовать несколько интерфейсов

```
type Hound interface {
    Hunt()
}
type Poodle interface {
    Bark()
}

type GoldenRetriever struct{name string}

func (GoldenRetriever) Hunt() { fmt.Println("hunt") }
func (GoldenRetriever) Bark() { fmt.Println("bark") }

```

https://goplay.space/#h_7ODwUAXfM


---

# Одному интерфейсу могут соответствовать много типов

```
type Poodle interface {
    Bark()
}

type ScandinavianClip struct{name string}
func (ScandinavianClip) Bark() { fmt.Println("bark") }


type ToyPoodle struct{name string}
func (ToyPoodle) Bark() { fmt.Println("bark") }
```

https://goplay.space/#0Mjn_Yd9K5W


---

# Интерфейсы: композиция

Интерфейс может встраивать в себя другой (определенный пользователем или импортируемый) интерфейс:

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

# Интерфейсы: композиция


```
type error interface {
    Error() string
}

```

Пример из io:
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

# Интерфейсы: циклические зависимости

```
type I interface {
    J
    i()
}

type J interface {
    K
    j()
}

type K interface {
    k()
    I
}
```

```
./prog.go:15:6: invalid recursive type K
```

https://goplay.space/#2fDIbsBoZfv


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
./prog.go:7:2: duplicate method bark
```

https://goplay.space/#wMw2VKOIysx


---

# Интерфейс может быть пустым

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

# Интерфейсы: interface{}

Как быть?

```
func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    PrintAll(names)
}
```

https://goplay.space/#1w7ksGW0uXh


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

Структура устарела, но алгоритм такой же:<br>
https://www.tapirgames.com/blog/golang-interface-implementation


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

https://goplay.space/#rjboVEC3V6w

---

background-size: 80%
background-image: url(img/internalinterfaces.png)

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
var s Speaker = string("test") // compile-time error
var s Speaker = io.Reader // compile time error
var h string = Human{} // compile time error

// runtime error
var s interface{};
h = s.(Human)
```


---

# Интерфейсы изнутри

Что здесь происходит?
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


---

# Интерфейсы изнутри

Динамический тип и значение можно посмотреть с помощью `%v` и `%T`

```
type Temp int

func (t Temp) String() string {
    return strconv.Itoa(int(t)) + " °C"
}

func main() {
    var x fmt.Stringer
    x = Temp(24)
    fmt.Printf("%v | %T\n", x, x) // 24 °C | main.Temp
}
```

https://goplay.space/#9ldo_icbhj0


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
    adder := Adder{id: 6754}
    for i := 0; i < b.N; i++ {
        Addifier(adder).Add(10, 32)
    }
}
```


---

# Интерфейсы изнутри: benchmark

```
$ GOOS=linux GOARCH=amd64 go tool compile -m addifier.go

Addifier(adder) escapes to heap
```

```
$ GOOS=linux GOARCH=amd64 go test -bench=.              
BenchmarkDirect-8       2000000000    1.60 ns/op    0 B/op   0 allocs/op
BenchmarkInterface-8    100000000     15.0 ns/op    4 B/op   1 allocs/op
```


---

# Интерфейсы: еще раз о ресиверах

https://goplay.space/#jm1bKNLABnB
<br><br>
https://stackoverflow.com/a/45653986
<br><br>
https://stackoverflow.com/a/48874650

---

# Интерфейсы

- это набор сигнатур методов
- интерфейс реализуется неявно
- интерфейс может встраивать другие интерфейсы
- имена методов интерфейса не должны повторяться
- интерфейс может быть пустым (не иметь методов), такому интерфейсу удовлетворяет любой тип


---

# Интерфейсы: интерактив

Реализовать интерфейс Adult
<br><br>
https://goplay.space/#A48l0-8FQX0


---

# Домашнее задание

Реализовать LRU-кэш на основе двусвязного списка
<br><br>
https://github.com/OtusGolang/home_work/tree/master/hw04_lru_cache

---

# Опрос

https://otus.ru/polls/11418/


---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
