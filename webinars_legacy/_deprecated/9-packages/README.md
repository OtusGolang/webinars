.center.icon[![otus main](https://drive.google.com/uc?id=1NPIi9Hw5ZjA5SK24lTXckDjNAPSuFAHi)]

---


class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Пакеты в языке Go

### Иван Ремень

---

class: top white
background-image: url(tmp/sound.svg)
background-size: 130%
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

.sound-top[
  # Как меня слышно и видно?
]

.sound-bottom[
  ## > Напишите в чат
  ### **+** если все хорошо
  ### **–** если есть проблемы cо звуком или с видео
]

---

# Цель занятия 

.right-image[
![](tmp/gopher.png)
]

# 
- Узнать, зачем нужны пакеты в Go.
- Узнать, как устроены пакеты в Go.
- Научиться управлять видимостью элементов пакета.
- Разобраться с internal пакетами.
- Научиться использовать функцию init().
- Изучить порядок инициализации пакетов.
- Узнать как использовать синонимы имен пакетов.
- Написать свой первый пакет.

---

# А зачем?

.right-image[
![](tmp/gopher.png)
]

## Зачем нужны пакеты?

---

# Структура программы с пакетом

```
MBP-Remen:example bhychik$ tree
.
└── src
    ├── calc
    │   └── calc.go
    └── mathematica
        ├── mult.go
        └── sum.go

3 directories, 3 files
```

---

# sum.go

```
package mathematica

//Function Sum returns sum of several numbers
func Sum(args ...int) int {
	acc := int(0)
	for _, a := range args {
		acc += a
	}
	return acc
}
```

---

# mult.go

```
package mathematica

//Function Mult returns multiplication of several numbers
func Mult(args ...int) int {
	acc := int(1)
	for _, a := range args {
		acc *= a
	}
	return acc
}
```

---

# calc.go

```
package main

import (
	"fmt"

	"mathematica"
)

func main() {
	fmt.Println(mathematica.Sum(5, 7, 8))
	fmt.Println(mathematica.Mult(2, 3, 10))
}
```

---

# Запуск и сборка

```
MBP-Remen:example bhychik$ ls
src
MBP-Remen:example bhychik$ GOPATH=`pwd`
MBP-Remen:example bhychik$ go build calc
MBP-Remen:example bhychik$ ./calc
20
60
MBP-Remen:example bhychik$
```

---

# Псевдонимы пакетов

```
package main

import (
	"fmt"

	m "mathematica"
)

func main() {
	fmt.Println(m.Sum(5, 7, 8))
	fmt.Println(m.Mult(2, 3, 10))
}
```

---

# Использование пакетов с одинаковыми именами

```
package main

import (
	"fmt"

	m1 "github.com/bhychik/mathematica"
	m2 "github.com/i.remen/mathematica"
)

func main() {
	fmt.Println(m1.Sum(5, 7, 8))
	fmt.Println(m2.Mult(2, 3, 10))
}
```

---

# Функция init()

## Пакет

```
package matr

var m [200][200]int

func GetElement(i int, j int) int {
	return m[i][j]
}

func init() {
	for i, v := range m {
		for j := range v {
			m[i][j] = i*100 + j
		}
	}
}
```

---

# Функция init()
## Использование пакета

```
package main

import (
	"fmt"
	"matr"
)

func main() {
	fmt.Printf("%d\n", matr.GetElement(3, 6))
}
```

---

# Импорт пакета без использования
## Использование пакета

```
package main

import (
	_ "nousing"
)

func main() {

}
```

---

# Импорт пакета без использования
## Пример пакета

```
package nousing

import (
	"fmt"
)

func Welcome() {
	fmt.Printf("Welcome!\n")
}

func init() {
	fmt.Printf("I am init\n")
}
```

---

# Импорт пакета без использования

### Но зачем?

---

# Импорт пакета без использования
## Реальный пример

```
package main
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)
 
func main() { 
    db, err := sql.Open("mysql", "root:password@/testbase")
     
    if err != nil {
        panic(err)
    } 
    defer db.Close()
     
    fmt.Printf("Hello world")
}
```

---

#internal пакеты
## Проект

```
MBP-Remen:example bhychik$ tree
.
├── calc
├── main
├── mex
├── nousingmain
└── src
    ├── matr
    │   ├── internal
    │   │   └── matr.go
    │   └── matr.go
    └── mex
        └── mex.go
```

---

#internal пакеты

```
package matr

import(
	"matr/internal"
)

var m [200][200]int

func GetElement(i int, j int) int {
	return m[i][j]
}

func init() {
	for i, v := range m {
		for j := range v {
			m[i][j] = i*100 + j
		}
	}

	internal.Welcome()
}

```

---

#internal пакеты
## Некорректное использование

```
package main

import (
	"fmt"
	"matr"
	_ "matr/internal"
)

func main() {
	fmt.Printf("%d\n", matr.GetElement(3, 6))
}
```

```
MBP-Remen:example bhychik$ go build mex
src/mex/mex.go:6:2: use of internal package matr/internal not allowed
```

---

#Вендоринг пакетов

- dep
- gb
- glide
- Модули (golang >= 1.11)

---


# Практика

### Пишем пакет для логирования

---

# Тест

https://forms.gle/5ETZWNTZGfY9XLyq8

---


# На занятии

- Узнали, зачем нужны пакеты в Go.
- Узнали, как устроены пакеты в Go.
- Научились управлять видимостью элементов пакета.
- Разобрались с internal пакетами.
- Научились использовать функцию init().
- Изучили порядок инициализации пакетов.
- Узнали как использовать синонимы имен пакетов.
- Написали свой первый пакет.

---

## Вопросы?

---

# Опрос

Не заполните заполнить опрос. Ссылка на опрос будет в слаке.

---

class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Спасибо за внимание!