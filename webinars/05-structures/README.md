.center.icon[![otus main](img/main.png)]

---


class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Структуры в Go

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
<!--
# Небольшой тест

.left-text[
Пожалуйста, пройдите небольшой тест. 
<br><br>
Возможно вы уже многое знаете про структуры в Go =)
<br><br>
[https://forms.gle/xLLab1NXH9NLKJij8](https://forms.gle/xLLab1NXH9NLKJij8)
]

.right-image[
![](img/gopher9.png)
]-->


# Структуры

Структуры - фиксированный набор именованных переменных. <br>
Переменные размещаются рядом в памяти и обычно используются совместно.

```
struct{}  // Пустая структура, не занимает памяти

type User struct { // Структура с именованными полями
  Id      int64
  Name    string
  Age     int
  friends []int64  // Приватный элемент
}
```
https://golang.org/ref/spec#Struct_types

---

# Литералы структур

```
var u0 User                      // Zero Value для типа User

u1 := User{}                     // Zero Value для типа User

u2 := &User{}                    // То же, но указатель

u3 := User{1, "Vasya", 23, nil}  // По номерам полей

u4 := User{                      // По именам полей
  Id:       1,
  Name:     "Vasya",
  friends:  []int64{1, 2, 3},
}

```

---

# Анонимные типы и структуры

Анонимные типы задаются литералом, у такого типа нет имени.<br>
Типичный сценарий использования: когда структура нужна только внутри одной функции. 

```
var wordCounts []struct{w string; n int}
```

```
var resp struct {
    Ok        bool `json:"ok"`
    Total     int  `json:"total"`
    Documents []struct{
        Id    int    `json:"id"`
        Title string `json:"title"`
    } `json:"documents"`
}
json.Unmarshal(data, &resp)
fmt.Println(resp.Documents[0].Title)
```
<br>
https://goplay.space/#rE-DsbSFgN1


---

# Анонимные типы и структуры

```
testCases := []struct{
    name     string
    input    string
    expected int
    err      error  
} {
    name: "case1",
    input: "aaa",
    expected: 10,
    err: nil,
}
```

---

# Размер и выравнивание структур

https://goplay.space/#0WdB68TTmkj <br>

```
unsafe.Sizeof(1)   // 8 на моей машине
unsafe.Sizeof("A") // 16 (длина + указатель)

var x struct {
    a bool   // 1 (offset 0)
    c bool   // 1 (offset 1)
    b string // 16 (offset 8)
}

unsafe.Sizeof(x) // 24!
```
![img/aling.png](img/align.png)

https://github.com/dominikh/go-tools/tree/master/cmd/structlayout <br>
https://en.wikipedia.org/wiki/Data_structure_alignment

---

# Указатели

Указатель - это адрес некоторого значения в памяти. <br>
Указатели строго типизированы. <br>
Zero Value для указателя - nil.

```
x := 1         // Тип int
xPtr := &x     // Тип *int
var p *int     // Тип *int, значение nil
```

https://goplay.space/#s-LG0fjQxmV

---

# Получение адреса

Можно получать адрес не только переменной, но и поля структуры или элемента массива или слайса. <br>
Получение адреса осуществляется с помощью оператора `&`.
```
var x struct {
    a int
    b string
    c [10]rune
}
bPtr := &x.b
c3Ptr := &x.c[2]
```

Но не значения в словаре!
```
dict := map[string]string{"a": "b"}
valPtr := &dict["a"]  // Не скомпилируется
```
https://github.com/golang/go/issues/11865
<br><br>

Также нельзя (и не нужно) получать указатель на функцию.

<br>
https://goplay.space/#5N5WqdIZVDS

---

# Разыменование указателей

Разыменование осуществляется с помощью оператора `*`:
```
a := "qwe"  // Тип string
aPtr := &a  // Тип *string
b := *aPtr  // Тип string, значение "qwe"

var n *int  // nil
nv := *n    // panic
```

В случае указателей на *структуры* вы можете обращаться к полям структуры без разыменования:
```
p := struct{x, y int }{1, 3}
pPtr := &p
fmt.Println(pPtr.x) // (*pPtr).x
fmt.Println((*pPtr).y)

pPtr = nil
fmt.Println(pPtr.x) // ?
```

https://goplay.space/#q3UDXozLcX9
<br>
https://golang.org/ref/spec#Selectors

---


# Копирование указателей и структур

При присвоении переменных типа структура - данные копируются.
```
a := struct{x, y int}{0, 0}
b := a
a.x = 1
fmt.Println(b.x) // ?
```

При присвоении указателей - копируется только адрес данных.
```
a := new(struct{x, y int})
b := a
a.x = 1
fmt.Println(b.x) // ?
```

```
a := struct{x *int}{new(int)}
b := a
*a.x = 1
fmt.Println(b.x) // ?
```

---

# Определение методов 

В Go можно определять методы у именованых типов (кроме интерфейсов)

```
type User struct {
    Id      int64
    Name    string
    Age     int
    friends []int64
}

func (u User) IsOk() bool {
    for _, fid := range u.friends {
        if u.Id == fid {
            return true
        }
    }
    return false
}

var u User
fmt.Println(u.IsOk()) // (User).IsOk(u)
```

https://golang.org/ref/spec#Method_declarations
<br>
https://goplay.space/#pp4iiJoQ8sO

---

# Методы типа и указателя на тип

Методы объявленные над типом получают копию объекта, поэтому не могут его изменять!
```
func (u User) HappyBirthday() {
  u.Age++ // Это изменение будет потеряно
}
```

Методы объявленные над указателем на тип - могут.
```
func (u *User) HappyBirthday() {
  u.Age++ // OK
}
```

https://goplay.space/#XP7fc8wxQ3P
<br><br>

Метод типа можно вызывать у значения и у указателя. <br>
Метод указателя можно вызывать у указателя и у значения, если оно адресуемо.

<br>
https://github.com/golang/go/wiki/CodeReviewComments#receiver-type

---

# Экспортируемые и приватные элементы

Поля структур, начинающиеся со строчной буквы - **приватные**, они будут видны
только в том же пакете, где и структура. <br><br>
Поля, начинающиеся с заглавной - **публичные**, они будут видны везде.

```
type User struct {
  Id      int64
  Name    string   // Экспортируемое поле
  Age     int
  friends []int64  // Приватное поле
}
```

Не совсем очевидное следствие: пакеты стандартной библиотеки, например, `encoding/json` тоже не могут
работать с приватными полями :)<br><br>
Доступ к приватным элементам (на чтение!) все же можно получить с помощью пакета `reflect`.

<br>
https://goplay.space/#g9sldeRCgaO

---

# Функции-конструкторы

В Go принят подход Zero Value: постарайтесь сделать так, что бы
ваш тип работал без инициализации, как реализованы, например
```
var b strings.Builder
var wg sync.WaitGroup
```

Если ваш тип содержит словари, каналы или инициализация обязательна - скройте
ее от пользователя, создав функции-конструкторы:

```
func NewYourType() (*YourType) {
  // ...  
}
func NewYourTypeWithOption(option int) (*YourType) {
  // ...
}
```

https://goplay.space/#5lfGpAcfTyU

---

# nil receiver

```
type RateLimiter struct {
    ...
}

func (r *RateLimiter) Allow() bool {
    if r == nil {
        return true
    }
    return r.allow()
}
```


---

# Задачка

.left-code[
Реализовать тип `IntStack`, который содержит стэк целых чисел. 
У него должны быть методы `Push(i int)` и `Pop() int`.

<br><br>
https://goplay.space/#xhAGg8vtX8N
]

.right-image[
![](img/gopher9.png)
]

---

# Встроенные структуры

В Go есть возможность "встраивать" типы внутрь структур. <br>
При этом у элемента структуры НЕ задается имя.

```
type LinkStorage struct {
    sync.Mutex                  // Только тип!
    storage map[string]string   // Тип и имя
}
```

Обращение к элементам встроенных типов:
```
var storage LinkStorage
storage.Mutex.Lock()     // Имя типа используется 
storage.Mutex.Unlock()   // как имя элемента структуры
```

---

# Продвижение методов

При встраивании методы встроенных структур можно вызывать у ваших типов!

```
// Вместо
storage.Mutex.Lock()
// можно просто
storage.Lock()
```


---

# Но, это не наследование

```
type Base struct {}

func (b Base) Name() string {
    return "Base"
}

func (b Base) Say() {
    fmt.Println(b.Name())
}

type Child struct {
    Base
    Name string
}

func (c Child) Name() string {
    return "Child"
}

var c Child
c.Say() // Увы "Base" :(
```
https://goplay.space/#AOyLzYid61L


---

# Тэги элементов структуры

К элементам структуры можно добавлять метаинформацию - тэги. <br>
Тэг это просто литерал строки, но есть соглашение о структуре такой строки.

<br>
Например,
```
type User struct {
    Id      int64    `json:"-"`    // Игнорировать в encode/json
    Name    string   `json:"name"`
    Age     int      `json:"user_age" db:"how_old"`
    friends []int64 
}
```

Получить информацию о тэгах можно через `reflect`
```
var u User
ageField := reflect.TypeOf(u).FieldByName("Age")
jsonFieldName := ageField.Get("json")  // "user_age"
```

https://github.com/golang/go/wiki/Well-known-struct-tags

---

# Использование тэгов для JSON сериализации

Для работы с JSON используется пакет `encoding/json`

```
// Можно задать имя поля в JSON документе
Field int `json:"myName"`

// Не выводить в JSON поля у которых Zero Value
Author *User `json:"author,omitempty"`

// Использовать имя поля Author, но не выводить Zero Value
Author *User `json:"omitempty"`

// Игнорировать это поле при сериализации / десереализации
Field int `json:"-"`
```

---

# Использование тэгов для работы с СУБД

Зависит от пакета для работы с СУБД.<br>
Например, для `github.com/jmoiron/sqlx`
```
var user User
row := db.QueryRow("SELECT * FROM users WHERE id=?", 10)
err = row.Scan(&user)
```

Для ORM библиотеки GORM `github.com/jinzhu/gorm` фич намного больше
```
type User struct {
  gorm.Model
  Name         string
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // set field size to 255
  MemberNumber *string `gorm:"unique;not null"` // set member number to unique and not null
  Num          int     `gorm:"AUTO_INCREMENT"` // set num to auto incrementable
  Address      string  `gorm:"index:addr"` // create index with name `addr` for address
  IgnoreMe     int     `gorm:"-"` // ignore this field
}
```

---

# Ещё раз про пустые структуры

https://dave.cheney.net/2014/03/25/the-empty-struct

```
type Set map[int]struct{}
```

```
ch := make(chan struct{})
ch <- struct{}{}
```


<!--
# Небольшой тест

.left-text[
Проверим что мы узнали за этот урок
<br><br>
[https://forms.gle/xLLab1NXH9NLKJij8](https://forms.gle/xLLab1NXH9NLKJij8)
]

.right-image[
![](img/gopher9.png)
]-->

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
https://otus.ru/polls/????/
]

.right-image[
![](img/gopher7.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
