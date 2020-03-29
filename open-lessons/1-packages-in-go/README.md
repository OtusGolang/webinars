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

	m "mathematica"
)

func main() {
	fmt.Println(m.Sum(5, 7, 8))
	fmt.Println(m.Mult(2, 3, 10))
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

# Практика

### Пишем пакет для логирования

---

# На занятии

- Разобрались зачем нужны пакеты
- Написали программу, использующую пакеты




## Вопросы?

---

class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Спасибо за внимание!