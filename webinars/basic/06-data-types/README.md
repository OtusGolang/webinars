background-image: url(../../img/title.svg)

---

background-image: url(../../img/rec.svg)

---
background-image: url(../../img/topic.svg)

.topic[Переменные и элементарные типы данных]
.tutor[Алексей Романовский]
.tutor_desc[software developer @ Resolver]

---

background-image: url(../../img/rules.svg)

---
# Цели занятия

* познакомиться с системой типов Go;
* научиться работать с переменными и указателями.

---
# Краткое содержание

* различные системы счисления;
* типизация в Go;
* объявление переменной;
* операция присваивания;
* арифметические операции;
* указатели;
* расположения переменных в памяти;
* понятие "zero value";
* приведение типов;
* константы.


---
# Различные системы счисления

Считаем до 100!
* в бинарной системе счисления
* в троичной

Часто используемые системы счисления:
* двоичная (бинарная)
* десятичная
* шестнадцатеричная

---
# Типизация в Go;
Го - язык со статической типизацией.
* типы данных проверяются и на этапе компиляции и на этапе выполнения
* каждая переменная имеет не только значение, но и тип
* Как правило, типы не приводятся автоматически (неявно)

---

# Объявление переменной

```go
var Storage map[string]string         // zero value

var storage = make(map[string]string) // автовывод типа

func Answer() int {
  return 42
}

func main() {
  var i int = 10
  j := i  // короткое объявление, только внутри функций
}
```

---

# Группировка объявлений

```go
type (
	MyId   int
	UserId string
)

var (
	a int    = 4
	b        = 1.2
	c UserId = "admin"
)

func foo() {
	var (
		ctx = context.Background()
		rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	)
	// ...
}
```
---

# Публичные и приватные идентификаторы

Публичные идентификаторы - те, которые видны за пределами вашего *пакета*.
Публичные идентификаторы начинаются с заглавной буквы `Storage`, `Printf`.


Приватные идентификаторы - начинаются со строчной буквы `i`, `j` и видны только в вашем пакете.
Структуры могут содержать как приватные так и публичные поля.

```go
type User struct {
    Name     string  // Будет видно в json.Marshal.
    password string  // Не будет видно.
}
```
---

# Какие есть элементарные типы данных в Go?

- Логические: `bool`
- Целые: `int`, `uint`, `int8`, `uint8`, ... 
- С плавающей точкой: `float32`, `float64 `
- Комплексные: `complex64`, `complex128`
- Строки: `string`, `rune`
- Указатели: `uintptr`, `*int`, `*string`, ...
- Алиасы к другим элементарным типам: `byte` = `uint8`, `rune` = `int32`
- Decimal только через сторонние модули.

https://golang.org/ref/spec#Types

---

# Литералы


Целые числа:

* 2: `0b1111`
* 8: `017`, `0o17`, `0O17`
* 10: `19`
* 16: `0xF`, `0XF` 

Дробные:

* `1.23`, `01.239`, `.23`
* `1.23e+2`, `1.23e-2`, `1.23e2`
* `0x1p2`


---
# Литералы

В числах можно использовать подчеркивания: `100_000_000`

Другое:

* `bool`
* complex: `5+6.7i`
* руны: `'a', '\n', '\u9333'`


---

# Строковые литералы


Строки:
```go
s := "hello world"            // в двойных кавычках, на одной строке
s := "hello \n world \u9333"  // c непечатными символами

// если нужно включить в строку кавычки или переносы строки 
// - используем обратные кавычки
s := `hello
"cruel"
'world'
`  
```

Закрепим, рассмотрев код: https://go.dev/play/p/BocJQjJY0sP

---

# Особенности целых чисел в Go

- Есть значение "по умолчанию" - это `0`
- Типы `int` и `uint` могут занимать 32 и 64 бита на разных платформах
- Нет автоматического преобразования типов
- `uintptr` - целое число, не указатель

---

# Операции присваивания
* `=`
* `:=`
* `+=, -=, *=, /=, %=, <<=, >>=, &=, |=, ^=`


---

# Арифметические операции

Все так же стандартно

```go
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << unsigned integer
>>   right shift            integer >> unsigned integer
```

`++` и `--` не возвращают значения. `++i` - не существует.


---

# Указатели

Указатель - это адрес некоторого значения в памяти. 
Указатели строго типизированы. 
Zero Value для указателя - nil.
Нельзя приводить указатели к целым числам.

```go
x := 1         // Тип int
xPtr := &x     // Тип *int
var p *int     // Тип *int, значение nil
```

https://goplay.space/#s-LG0fjQxmV

---

# Расположение переменных в памяти

### https://go.dev/play/p/MZ3WUsAgVRY

* переменная - именованная область в памяти
* размер области зависит от типа
* big endian, little endian


---

# Понятие "zero value"

значения всегда инициализируются. 
Если вы не указали значение - оно будет нулевым.
У каждого типа есть свое нулевое значение.



`0, 0.0, "", false, nil, {...}`


---

# Приведение типов

### https://go.dev/play/p/dd1SWESaeBu

* неявные, явные
* приведение Т(х)
* проверки на переполнение - не происходит (для производительности)
* обрезка, дополнение для целых и дробных чисел
* в строки
* алиасы

См также:
* [Integer Overflow in Golang (Medium.com)](https://medium.com/@griffinish/integer-overflow-in-golang-9e13e274c8a5)
* [Conversions: complete list](https://yourbasic.org/golang/conversions/)

---

# Константы

* Константы - неизменяемые значения, доступные только во время компиляции.
* Только `rune`, `string`, `bool` и числа.

```go
const PI = 3             // принимает подходящий тип
const pi float32 = 3.14  // строгий тип

const (
  TheA = 1
  TheB = 2
)

const (
  X = iota   // 0
  Y          // 1
  Z          // 2
)

```

---

# Константы

### https://go.dev/play/p/w8mTXMI1AH2

См также:
* [A tour of go: constants](https://go.dev/tour/basics/15)

---


# Во время компиляции?

```go
package main

const HelloConst = 3

var HelloVar = 5

func main() {
    print(HelloVar, HelloConst)
}
```

Скомпилируем и посмотрим символы
```shell
$ go build -o 1.out 1.go
$ go tool nm 1.out  | grep Hello
 10bd148 D main.HelloVar
```

---

background-image: url(../../img/questions.svg)

---

background-image: url(../../img/poll.svg)

---

background-image: url(../../img/next_webinar.svg)
.announce_date[23 августа]
.announce_topic[Управляющие конструкции языка и отладка программ]


---
background-image: url(../../img/thanks.svg)


.announce_date[Ссылка на опрос:]
.announce_topic[https://otus.ru/polls/70293/]
.tutor[Алексей Романовский]
.tutor_desc[Преподаватель]
