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

# Функции и методы <br> в Go

### Алексей Бакин

---

# О чем будем говорить

- ### области видимости
- ### виды функций и зачем они нужны

---

# Настройка на занятие

.left-text[
Пожалуйста, пройдите небольшой тест.
<br><br>
Возможно, вы уже многое знаете про функции в Go.
<br><br>
Ссылка в чате
]

.right-image[
![](img/gopher9.png)
]

---

# Области видимости и блоки

```
var a = 1 // <- уровень пакета

func main() {
	fmt.Println("1: ", a)

	a := 2 // <-- уровень блока функции
	fmt.Println("2: ", a)
	{
		a := 3 // <-- уровень пустого блока
		fmt.Println("3: ", a)
	}
	fmt.Println("4: ", a) // <-- ???

	f()
}

func f() {
	fmt.Println("5: ", a) // <-- ???
}
```

https://play.golang.org/p/NcjESEYxQAN

---

# Неявные блоки: for, if, switch, case, select

```
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

```
if v, err := doSmth(); err != nil {
    fmt.Println(err)
} else {
    process(v)
}
```

```
switch i := 2; i * 4 {
case 8:
    j := 0
    fmt.Println(i, j)
default:
    // "j" is undefined here
    fmt.Println(i)
}
```

---

# Области видимости

- Всеобщий блок (universe block) - видят все.

```
bool, int32, int64, float64...
nil
make, new, panic...
true or false
```

- Блок пакета (package block) - видно во всех файлах пакета.

```
package mypackage

type Student struct {
	name string
}

const studentsCount = 10
```

---

# Области видимости

Импорты в блоке файла (file block).

```
// sandbox.go
package main

import "fmt”

func main() {
	fmt.Println("main”)
	f()
}
```

```
// utils.go

package main

// <-- тут тоже нужен import "fmt"

func f() {
	fmt.Println("f”)
}
```

---

# Области видимости

```
func main() {
    fmt.Println(v)
    v := 1 // undefined: v
}
```

```
func main() {
	{
		{
			var a = 22
			fmt.Println(a)
		}
		fmt.Println(a) // undefined: a
	}
}
```

---

# Контрольный вопрос: сколько раз объявлен x?

```
package main

import "fmt"

func f(x int) {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}

var x int

func main() {
	var x = 200
	f(x)
}
```

---

# Объявление функции

```
//   Имя функции            возвращаемые значения
//       |                       |       |
func TrySayHello(name string) (string, error)
//                |      |
//           параметр   тип параметра  
```

Интересное:
- в го нет дефолтных значений для параметров
- функция может возвращать несколько значений
- функция - first class value, можем работать как с обычным значением
- параметры в функцию передаются по значению

---

# Примеры функций

```
func Hello() {
    fmt.Println("Hello World!")
}

func greet(user string) {
        fmt.Println("Hello " + user)
}

func add(x int, y int) int {
        return x + y
}

func add(x, y int) int {
        return x + y
}
```

---

# Примеры функций: несколько значений

```
func addMult(a, b int) (int, int) {
        return a + b, a * b
}
```

```
func SquaresOfSumAndDiff(a int64, b int64) (int64, int64) {
        x, y := a + b, a - b
        s := x * x
        d := y * y
        return s, d
}
```

---

# Пример variadic функции

```
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    sum(5, 7)
    sum(3, 2, 1)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```

https://play.golang.org/p/c2jXVc4ts1-

---

# Анонимные функции

```
func main() {
	func() {
		fmt.Println("Hello ")
	}()

	sayWorld := func() {
		fmt.Println("World!")
	}
	sayWorld()
}
```

https://play.golang.org/p/3Ta6LGb1-tN

---

# Анонимные функции: зачем?

---

# Анонимные функции: зачем?

### Запуск горутины

```
go func() {
    ...
}()
```

### Управление поведением

```
people := []string{"Alice", "Bob", "Dave"}
sort.Slice(people, func(i, j int) bool {
    return len(people[i]) < len(people[j])
})
fmt.Println(people)
```

https://play.golang.org/p/TwoNgyWJNwM

---

# Замыкания

```
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()

    fmt.Println(nextInt()) // <-- ?
    fmt.Println(nextInt()) // <-- ?
    fmt.Println(nextInt()) // <-- ?

    newInts := intSeq()
    fmt.Println(newInts()) // <-- ?
}
```

https://play.golang.org/p/w-8lPCNFrbX

---

# Замыкания: middleware

```
package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/hello", hello)
  http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "<h1>Hello!</h1>")
}
```

---

# Замыкания: middleware

```
package main

import (
  "fmt"
  "net/http"
  "time"
)

func main() {
  http.HandleFunc("/hello", timed(hello))
  http.ListenAndServe(":3000", nil)
}

func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    f(w, r)
    end := time.Now()
    fmt.Println("The request took", end.Sub(start))
  }
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "<h1>Hello!</h1>")
}
```

---

# Замыкания: defer

```
func main() {
	s := "hello"

	defer fmt.Println(s)
	defer func() {
		fmt.Println(s)
	}()

	s = "world"
}
```
https://play.golang.org/p/LCBMgmREhjE


---

# Функции: именованные возвращаемые значения

```
func sum(a, b int) (s int) {
    s = a + b
    return
}
```

---

# Функции: именованные возвращаемые значения

```
func doSmth() (s int, err error) {
    defer func() {
        if err != nil {
            s = 0
        }
    }

    s = 1
    err = errors.New("test")
    // тут может быть много логики
    // и несколько выходов из функции с ошибкой

    return
}
```

https://play.golang.org/p/dUnEn-uG_Ra

---

# Функции: рекурсия

```
// n! = n×(n-1)! where n >0
func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	}
	return 1
}
```

В Go нет оптимизации хвостовой рекурсии и вряд ли будет:
https://github.com/golang/go/issues/22624

---

# Функции: сигнатура

Сигнатура - это "тип функции".

<br/>

Перегрузки нет.

<br/>

```
func()
func(x int) int
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
```

---

# Функции: сигнатура

```
package main

type SumFunc func(base int, arguments ...int) int

func main() {

	var summer SumFunc
	summer = func(a int, args ...int) int {
		for _, v := range(args) {
			a = a + v
		}
		return a
	}

	fmt.Println(summer(1, 2, 3, 4)) // 10
}
```

---

# Методы и функции

```
type Book struct {
	pages int
}
func (b Book) Pages() int {
	return b.pages
}
func (b *Book) SetPages(pages int) {
	b.pages = pages
}

// "the same"
func Pages(b Book) int {
	return b.pages
}

func SetPages(b *Book, pages int) {
	b.pages = pages
}
```

---

# Методы и функции

```
type Concater string

func (c Concater) do(i int) string {
	return fmt.Sprintf("%s: %d", c, i)
}
```

```
c := Concater("test")
fmt.Println(c.do(0)) // <---- Explicit call

f1 := Concater.do // <------- Method expressions
fmt.Println(f1(c, 1))

f2 := c.do // <-------------- Method value
fmt.Println(f2(2))

f3 := func(i int) string {
	return c.do(i) // <------ Closure
}
fmt.Println(f3(3))
```

https://play.golang.org/p/ggUP7aFZ5EG

---

# Повторение

.left-text[
Давайте проверим, что вы узнали за этот урок, а над чем стоит еще поработать.
<br><br>
Ссылка в чате
]

.right-image[
![](img/gopher9.png)
]

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
Ссылка в чате.
]

.right-image[
![](img/gopher.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
