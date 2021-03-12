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
    - если есть проблемы cо звуком или с видео]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Слайсы и словари

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем — задаем вопросы.
* ### Чат вижу — могу ответить не сразу.
* ### После занятия — оффтопик, ответы на любые вопросы.

---

# О чем будем говорить

* ### Массивы
* ### Слайсы
* ### Словари

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
![](img/gopher_science.png)
]

---

# Массивы

```
var arr [256]int         // фиксированная длина

var arr [10][10]string   // может быть многомерным

arr := [...]int{1, 2, 3} // автоматический подсчет длины

arr := [10]int{1, 2, 3, 4, 5}
```

---

# Массивы: операции

```
v := arr[1] // чтение

arr[3] = 1  // запись

len(arr)    // длина массива

arr[2:4]    // получение слайса
```

https://goplay.tools/snippet/fJIrJ-PciZU

---

# Слайсы

Слайсы — это те же "массивы", но переменной длины.

<br/>
Создание слайсов:

```
var s []int  // не-инициализированный слайс, nil

s := []int{} // с помощью литерала слайса

s := make([]int, 3) // с помощью функции make, s == {0,0,0}
```

---

# Слайсы: как они устроены?

```
// runtime/slice.go
type slice struct {
  array unsafe.Pointer
  len   int
  cap   int
}
```

```
l := len(s) // len — вернуть длину слайса
c := cap(s) // cap — вернуть емкость слайса
```

```
s := make([]int, 3, 10) // s == {?}
```

Отличное описание: https://blog.golang.org/go-slices-usage-and-internals

---

# Массивы: операции

```
v := a[1] // чтение

a[3] = 1  // запись

len(a)    // длина массива

a[2:4]    // получение слайса
```

---

# Слайсы: операции

```
v := s[1] // чтение

s[3] = 1  // запись

len(s)    // длина слайса

s[2:4]    // получение подслайса
```

---

# Слайсы: добавление элементов

```
s = append(s, 1)       // добавляет 1 в конец слайса

s = append(s, 1, 2, 3) // добавляет 1, 2, 3 в конец слайса

s = append(s, s2...)   // добавляет содержимое слайса s2 в конец s
```
```
var s []int            // s == nil
s = append(s, 1)       // s == {1} append умеет работать с nil-слайсами
```

https://goplay.tools/snippet/NXZWljF65zd

---

# Слайсы: добавление элементов

.full-image[
![](img/slice_1.png)
]

---

# Слайсы: добавление элементов

.full-image[
![](img/slice_2.png)
]

---

# Слайсы: добавление элементов

.full-image[
![](img/slice_3.png)
]

---

# Авто-увеличение слайса

.full-image[
![](img/slice_4.png)
]
---

# Авто-увеличение слайса

Если `len < cap` — увеличивается `len`

<br/>
Если `len = cap` — увеличивается `cap`, выделяется новый кусок памяти, данные копируются.

```
func main() {
	s := []int{1}
	for i := 0; i < 10; i++ {
		fmt.Printf("ptr %0x   len: %d \tcap %d  \t\n",
			&s[0], len(s), cap(s))
		s = append(s, i)
	}
}
```
https://goplay.tools/snippet/UjQR5fiudyO

---

# Получение под-слайса (нарезка)

`s[i:j]` — возвращает под-слайс, с `i` -ого элемента включительно, по `j` -ый не влючительно.

Длинна нового слайса будет `j-i`.

```
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

s2 := s[3:5] // ?

s2 := s[3:]  // ?

s2 := s[:5]  // ?

s2 := s[:]   // копия s (shallow)
```

https://goplay.tools/snippet/6nHKyMNjQbO

---

# Получение под-слайса (нарезка)



```
s := []byte{1, 2, 3, 4, 5}

s2 := s[2:5]
```

<br>

.left-image[
![](img/slice.png)
]

.right-image[
![](img/slice2.png)
]

---

# Неочевидные следствия

```
arr := []int{1, 2}
arr2 := arr // копируется только заголовок, массив остался общий
arr2[0] = 42

fmt.Println(arr[0]) // ?

arr2 = append(arr2, 3, 4, 5, 6, 7, 8) // реаллокация
arr2[0] = 1

fmt.Println(arr[0]) // ?
```

https://goplay.tools/snippet/QrBpZWTeaos

---

# Неочевидные следствия

.full-image[
![](img/slice_ex_1.png)
]

---

# Неочевидные следствия

.full-image[
![](img/slice_ex_2.png)
]

---

# Неочевидные следствия

.full-image[
![](img/slice_ex_3.png)
]

---

# Неочевидные следствия

.full-image[
![](img/slice_ex_4.png)
]

---

# Правила работы со слайсами

### Функции изменяющие слайс
- принимают shalow копии
- возвращают новый слайс

```
func AppendUniq(slice []int, slice2 []int) []int {
  ...
}

s = AppendUniq(s, s2)
```

### Копирование слайса

```
s := []int{1,2,3}
s2 := make([]int, len(s))
copy(s2, s)
```

### SliceTricks

https://github.com/golang/go/wiki/SliceTricks

---

# Сортировка

```
  s := []int{3, 2, 1}
  sort.Ints(s)
```

```
  s := []string{"hello", "cruel", "world"}
  sort.Strings(s)
```
https://goplay.tools/snippet/hTEHP-bdemH

---

# Сортировка: типы

```
type User struct {
	Name string
	Age  int
}

func main() {
	s := []User{
		{"vasya", 19},
		{"petya", 18},
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Age < s[j].Age
	})
	fmt.Println(s)
}
```

https://goplay.tools/snippet/1K0s37F0z4I

---

# Слайсы: итерирование


```
// Индекс и значение
for i, v := range s {
  ...
}
```
```
// Только индекс
for i := range s {
  ...
}
```
```
// Только значение
for _, v := range s {
  ...
}
```

---

# Задачка

.left-code[
Написать функцию `Concat`, которая получает несколько слайсов и склеивает их в один длинный.
`{ {1, 2, 3}, {4, 5}, {6, 7} }  => {1, 2, 3, 4, 5, 6, 7}`
<br><br>

https://goplay.tools/snippet/GGlrv2nmqYb
]

.right-image[
![](img/gopher.png)
]

---

# Словари (map)

- Отображение ключ => значение.

- Реализованы как хэш-таблицы.

- Аналогичные типы в других языках: в Python — `dict`, в JavaScript — `Object`, в Java — `HashMap`, в C++ — `unordered_map`.


---

# Словари: cоздание

```
var cache map[string]string  // не-инициализированный словарь, nil

cache := map[string]string{} // с помощью литерала, len(cache) == 0

cache := map[string]string{  // литерал с первоначальным значением
	"one":   "один",
	"two":   "два",
	"three": "три",
}

cache := make(map[string]string)      // тоже что и map[string]string{}

cache := make(map[string]string, 100) // заранее выделить память
                                      // на 100 ключей
```

---

# Словари: операции

```
value := cache[key]     // получение значения,

value, ok := cache[key] // получить значение, и флаг того что ключ найден

_, ok := cache[key]     // проверить наличие ключа в словаре

cache[key] = value      // записать значение в инициализированный(!) словарь

delete(cache, key)      // удалить ключ из словаря, работает всегда
```

Подробное описание: https://blog.golang.org/go-maps-in-action

<br>

Про устройство мапы:
- https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html
- https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics

---

# Словари: итерирование

```
// Ключ и значение
for key, val := range cache {
  ...
}
```
```
// Только ключ
for key := range cache {
  ...
}
```
```
// Только значение
for _, val := range cache {
  ...
}
```

---

# Словари: списки ключей и значений

В Go нет функций, возвращающих списки ключей и значейний словаря. (Почему?)

<br/>
Получить ключи:
```
var keys []string
for key, _ := range cache {
  keys = append(keys, key)
}
```

<br/>
Получить значения:
```
values := make([]string, 0, len(cache))
for _, val := range cache {
  values = append(values, val)
}
```
---

# Словари: требования к ключам

Ключом может быть любой типа данных,
для которого определена операция сравнения `==` :
- строки, числовые типы, bool каналы (chan);
- интерфейсы;
- указатели;
- структуры или массивы содержащие сравнимые типы.

```
type User struct {
  Name string
  Host string
}
var cache map[User][]Permission
```

Подробнее https://golang.org/ref/spec#Comparison_operators

---

# Словари: порядок ключей

* ### Какой порядок итерирования по словарю?
* ### Что будет, если удалить ключ во время итерирования?
* ### Что будет, если добавить ключ во время итерирования?

https://goplay.tools/snippet/SmisQCUpCGb

---

# Использование Zero Values

Для слайсов и словарей, zero value — это nil .

<br/>
С таким значением будут работать функции и операции читающие данные, например:

```
var seq []string             // nil
var cache map[string]string  // nil
l := len(seq)       // 0
c := cap(seq)       // 0
l := len(cache)     // 0
v, ok := cache[key] // "", false
```

Для слайсов будет так же работать `append`
```
var seq []string             // nil
seq = append(seq, "hello")   // []string{"hello"}
```

---

# Использование Zero Values

Вместо
```
hostUsers := make(map[string][]string)
for _, user := range users {
  if _, ok := hostUsers[user.Host]; !ok {
    hostUsers[user.Host] = make([]string)
  }
  hostUsers[user.Host] = append(hostUsers[user.Host], user.Name)
}
```

Можно
```
hostUsers := make(map[string][]string)
for _, user := range users {
  hostUsers[user.Host] = append(hostUsers[user.Host], user.Name)
}
```

---

# Повторение

.left-text[
Давайте проверим, что вы узнали за этот урок, а над чем стоит еще поработать.
<br><br>
Ссылка в чате
]

.right-image[
![](img/gopher_science.png)
]

---

# Следующее занятие

## Структуры

<br>
<br>
<br>

## 16 марта, вторник

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
