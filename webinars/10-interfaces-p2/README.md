.center.icon[![otus main](https://drive.google.com/uc?id=1NPIi9Hw5ZjA5SK24lTXckDjNAPSuFAHi)]

---


class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Интерфейсы в Go. <br>Часть 2

### Антон Телышев

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

# О чем будем говорить:

* Значение типа интерфейс и ошибки, связанные с nil
* Правила присваивания значений переменным типа интерфейс
* Опасное и безопасное приведение типов (type cast)
* Использование switch в контексте интерфейсов
* Слайсы и словари с интерфейсами
* Реализация подхода обобщенного программирования (generics) через интерфейсы

---

# Интерфейсы - вспоминаем прошлое занятие

 - это набор сигнатур методов
 - который реализуется неявно
 - интерфейсы могут встраивать другие интерфейсы
 - имена методов не должны повторяться
 - интерфейс может быть пустым (не иметь методов), такому интерфейсу удовлетворяет любой тип

---

# Значение типа интерфейс

<br>состоит из динамического типа и значения
<br>мы можем их смотреть при помощи %v и %T

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
import (
	"fmt"
	"reflect"
)

type MyError struct {}

func (e MyError) Error() string {
	return "smth happened"
}

func main() {
	var e error
	e = MyError{}

	fmt.Println(reflect.TypeOf(e).Name()) // main MyError
	fmt.Printf("%T\n", e)                 // main MyError
}
```
https://goplay.tools/snippet/Xmsbk5DEdqE


---

# Значение типа интерфейс

nil - нулевое значение для интерфейсного типа

```
type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {
	var s Shape
	fmt.Println("value of s is", s)     // value of s is <nil>
	fmt.Printf("type of s is %T\n", s)  // type of s is <nil>
}
```
https://goplay.tools/snippet/sxE9AxAQ8lH


---

# Значение типа интерфейс
<br>

```
type Rect struct {
    width  float64
    height float64
}

func (r Rect) Area() float64 {
    return r.width * r.height
}

func (r Rect) Perimeter() float64 {
    return 2 * (r.width + r.height)
}

func main() {
    var s Shape
    s = Rect{5.0, 4.0}
    fmt.Printf("type of s is %T\n", s)          // type of s is main.Rect
    fmt.Printf("value of s is %v\n", s)         // value of s is {5 4}
    fmt.Println("area of rectange s", s.Area()) // area of rectange s 20
}
```
https://goplay.tools/snippet/wbmnTcriHJ-


---

# Значение типа интерфейс

<br>
Переменная типа интерфейс `I` может принимать значение любого типа, который реализует интерфейс `I`.

```
type I interface {
    method1()
}

type T1 struct{}
func (T1) method1() {}

type T2 struct{}
func (T2) method1() {}
func (T2) method2() {}

func main() {
    var i I = T1{}

    i = T2{}
    fmt.Println(i) // {}
}

```
https://goplay.tools/snippet/a8PLrfRQL02


---

#  Интерфейсы: nil

<br>
Значение интерфейсного типа равно `nil` тогда и только тогда, когда `nil` его статическая и динамическая части.

<br>

https://goplay.tools/snippet/E8_TX3Zwznn

<br>

http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#nil_in_nil_in_vals

---

# Значение типа интерфейс
<br>
Что выведет программа?

```
package main

import (
    "io"
    "log"
    "os"
    "strings"
)

func main() {
    var r io.Reader

    r = strings.NewReader("hello")
    r = io.LimitReader(r, 4)

    if _, err := io.Copy(os.Stdout, r); err != nil {
        log.Fatal(err)
    }
}
```
https://goplay.tools/snippet/Tkx-7sKZhYD


---

# Правила присваиваний (assignability rules):
<br>
- Если переменная `T` реализует интерфейс `Callable`, мы можем присвоить ее переменной типа интерфейс `Callable`.

```
type Callable interface {
   f() int
}

type T int

func (t T) f() int {
    return int(t)
}

var c Callable
var t T
c = t
```


https://medium.com/golangspec/assignability-in-go-27805bcd5874

---

# Интерфейсы: присваивание

<br>

```
type I1 interface {
    M1()
}

type I2 interface {
    M1()
}

type T struct{}

func (T) M1() {}

func main() {
    var v1 I1 = T{}
    var v2 I2 = v1
    _ = v2
}
```

<br> валидно?

<br>

https://goplay.tools/snippet/4nNy7tBRbNJ

---

# Интерфейсы: присваивание

<br>Структура (вложенность) не имеет значения - `v1` и `v2` удовлетворяют `I1`, `I2`.
Порядок методов также не имеет значения.

```

type I1 interface { M1(); M2() }

type I2 interface { M1(); I3 }

type I3 interface { M2() }

type T struct{}

func (T) M1() {}
func (T) M2() {}

func main() {
    var v1 I1 = T{}
    var v2 I2 = v1
    _ = v2
}

```

<br>

https://goplay.tools/snippet/M-5AXWN2Es4

---


# Интерфейсы: присваивание

<br> валидно?

```
package main

type I1 interface { M1() }

type I2 interface { M1(); M2() }

type T struct{}

func (T) M1() {}

func main() {
	var v1 I1 = T{}
	var v2 I2 = v1
	_ = v2
}
```

<br>

https://goplay.tools/snippet/HhAPdPUNrh7

---

# Интерфейсы: присваивание

Что, если мы хотим присвоить переменной конкретного типа - значение типа интерфейс?


```
type I1 interface {
    M1()
}

type T struct{}
func (T) M1() {}

func main() {
    var v1 I1 = T{}
    var v2 T = v1   // Boom!
    _ = v2
}
```

<br>

https://goplay.tools/snippet/GrpYzhBcPQr

---

# Интерфейсы: type assertion


`x.(T)` проверяет, что конкретная часть значения `x` имеет тип `T` и `x != nil`:

	- если T - не интерфейс, то проверяем, что динамический тип x это T
	- если T - интерфейс: то проверяем, что динамический тип x его реализует
---

# Интерфейсы: type assertion

<br>

```
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string) // hello true
	fmt.Println(s, ok)

	r, ok := i.(fmt.Stringer) // <nil> false
	fmt.Println(r, ok)

	f, ok := i.(float64) // 0 false
	fmt.Println(f, ok)
```

<br>

https://goplay.tools/snippet/4VFT1joBgB6

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

<br>

https://golangci-lint.run/usage/configuration/
```
check-type-assertions: true
```

---


# Интерфейсы: type switch

<br>
можем объединить проверку нескольких типов в один type switch:

<br>

https://goplay.tools/snippet/QS_yLkiajPp

---


# Интерфейсы: type switch


как и в обычном `switch` можем объединять типы:

```
    case T1, T2:
            fmt.Println("T1 or T2")
    }
```

и обрабатывать `default`:

```
var v I
switch v.(type) {
default:
        fmt.Println("fallback")
}
```

---


# Интерфейсы: type assertion & type switch

<br>
что-то такое происходит в пакете fmt:

```
type Stringer interface {
    String() string
}

func ToString(any interface{}) string {
    if v, ok := any.(Stringer); ok {
        return v.String()
    }
    switch v := any.(type) {
    case int:
        return strconv.Itoa(v)
    case float:
        return strconv.Ftoa(v, 'g', -1)
    }
    return "???"
}
```

---

# Значение типа интерфейс

<br>
реализовать функцию zoo

<br>

https://goplay.tools/snippet/XmnDh8X03nV


---

# Интерфейсы: type assertion: T(v)

<br>
interface type -> concrete type

```
type I interface {
    M()
}

type T struct {}
func (T) M() {}

func main() {
    var v I = T{}
    fmt.Println(T(v)) // Boom!
}
```

```
cannot convert v(type I) to type T: need type assertion
```

---

# Интерфейсы: type assertion: I2(v)

<br>
interface type -> interface type

```
type I1 interface {
    M()
}

type I2 interface {
    M()
    N()
}

func main() {
    var v I1
    fmt.Println(I2(v)) // Boom!
}
```

```
main.go:16: cannot convert v (type I1) to type I2:
	I1 does not implement I2 (missing N method)
```

А наоборот?

---

# Интерфейсы: type assertion: T = v1

```
type I1 interface {
    M1()
}

type T struct{}

func (T) M1() {}

func main() {
    var v1 I1 = T{}
    var v2 T = v1 // Boom!
    _ = v2
}
```

```
cannot convert v (type I) to type T: need type assertion
```


---

# Интерфейсы: type assertion для конкретных типов


<br>
Для обычных типов:

```
type I interface {
    M()
}

type T struct{}

func (T) M() {}

func main() {
    var v1 I = T{}
    v2 := v1.(T)
    fmt.Printf("%T\n", v2) // main.T
}
```


---

# Интерфейсы: type assertion для конкретных типов


<br>
Для интерфейсов:

```

type I interface {
	M()
}

type T1 struct{}

func (T1) M() {}

type T2 struct{}

func main() {
	var v1 I = T1{}
	v2 := v1.(T2) // compile time error: impossible type assertion: 
				  // T2 does not implement I (missing M method)
	fmt.Printf("%T\n", v2)
}
```

---

# Интерфейсы: type assertion для конкретных типов

<br> Динамические части не совпадают:

```
type I interface {
    M()
}

type T1 struct{}
func (T1) M() {}

type T2 struct{}
func (T2) M() {}

func main() {
    var v1 I = T1{}
    v2 := v1.(T2) // runtime error.
    fmt.Printf("%T\n", v2)
}
```

```
panic: interface conversion: main.I is main.T1, not main.T2
```

---

# Интерфейсы: type assertion для конкретных типов


Можем проверить, выполнится ли приведение при помощи
multi-valued type assertion:

```
type I interface {
    M()
}

type T1 struct{}
func (T1) M() {}

type T2 struct{}
func (T2) M() {}

func main() {
    var v1 I = T1{}
    v2, ok := v1.(T2) // Boom!
    if !ok {
        fmt.Printf("ok: %v\n", ok)      // ok: false
        fmt.Printf("%v,  %T\n", v2, v2) // {},  main.T2
    }
}
```

---


# Интерфейсы: type assertion для интерфейсов

```
type I1 interface {
    M()
}

type I2 interface {
    I1
    N()
}

type T struct{
    name string
}
func (T) M() {}
func (T) N() {}

func main() {
    var v1 I1 = T{"foo"}
    var v2 I2
    v2, ok := v1.(I2)
    fmt.Printf("%T %v %v\n", v2, v2, ok) // main.T {foo} true
}
```

---


# Интерфейсы: type assertion для интерфейсов

```
type I1 interface {
    M()
}

type I2 interface {
    N()
}

type T struct {}
func (T) M() {}

func main() {
    var v1 I1 = T{}
    var v2 I2
    v2, ok := v1.(I2)
    fmt.Printf("%T %v %v\n", v2, v2, ok) // <nil> <nil> false
}
```

---

# Интерфейсы: type assertion для интерфейсов

<br>
nil всегда паникует

```
type I interface {
    M()
}

type T struct{}

func (T) M() {}

func main() {
    var v1 I
    v2 := v1.(T) // panic: interface conversion: main.I is nil, not main.T
    fmt.Printf("%T\n", v2)
}
```

---

# Интерфейсы: почти дженерики

есть: map, slice, etc.

https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md


---

# Интерфейсы: почти дженерики


Чтобы реализовать общие алгоритмы мы можем воспользоваться интерфейсами:

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
...
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

# Опрос

.left-text[
Заполните пожалуйста опрос
<br>
https://otus.ru/polls/19013/
]

.right-image[
![](img/gopher7.png)
]


---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
