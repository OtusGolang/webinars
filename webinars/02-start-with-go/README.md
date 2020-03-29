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
  ### !проверить запись!
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Инструментарий и начало <br> работы с Go

### Антон Телышев

---

# План занятия

.big-list[
* Подробнее про GOPATH и GOROOT
* Сборка модулей и установка программ: go get, go build, go install
* Кросс-компиляция
* Модули и зависимости: go mod
* Форматирование кода: go fmt, goimports
* Линтеры: go vet, golint, gometalinter, golangci-lint
* Как сдавать домашние задания?
]

---

# Установка Go

<b>Getting Started</b></br>
https://golang.org/doc/install<br><br>

<b>Downloads</b><br>
https://golang.org/dl/<br><br>

Проще всего через `apt-get`
```
sudo apt-get update

sudo apt-get install golang
```

Или просто скачать с официального сайта
```
wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz

sudo ln -s /usr/local/go/bin/go /usr/bin/go
```

Готово!
---

# GOROOT

`GOROOT` - переменная, которая указывает где лежит ваш дистрибутив Go, т.е.
компилятор, утилиты и стандартная библиотека. В новых версия Go (&gt; 1.0) утилиты сами определяют расположение Go.

<br><br>

Однако, вы можете узнать `GOROOT`

```
$ go env | grep ROOT
GOROOT="/usr/local/go"
```

И можете посмотреть исходный код Go =)
```
vim /usr/local/go/src/runtime/slice.go
```

---

# TL;DR

.main-image[
![img/path.png](img/path.png)
]

---

# GOPATH

`GOPATH` - переменная окружения, показывает где лежит ваше дерево исходников.

<br><br>
Крайне желательно задать эту переменную явно, добавив в `.bashrc` например
```
export GOPATH=/path/your/go/projects
```

Однако, если не задать, то `GOPATH` будет предполагаться `/home/<username>/go`


---

# Загрузка пакетов

Давайте выполним команду:
```
go get -d github.com/OtusGolang/webinars_practical_part/1.1_first_go_program
```

```
$ tree -L 5 $GOPATH
/Users/anthony/go/
├── bin
├── pkg
└── src
    └──github.com/OtusGolang
        └── webinars_practical_part
            └── 1.1_first_go_program
                ├── README.md
                ├── feeds.go
                ├── go.mod
                ├── main.go
                └── searcher
                    ├── rss.go
                    └── search.go
3 directories, 6 files
```

---

# Установка пакетов

Теперь выполним команду
```
go install github.com/OtusGolang/webinars_practical_part/1.1_first_go_program
```

```
$ ls $GOPATH/bin
1.1_first_go_program
```

Можем запустить!
```
$ 1.1_first_go_program
...
```
---

# Базовые команды

`go get -d` - скачивает пакеты из Git репозиториев в `$GOPATH/src`.
<br><br>
`go install` собирает и устанавливает в указанные пакеты в `$GOPATH/pkg` и `$GOPATH/bin`.
<br><br>
`go get` (без флажка `-d`) - так же вызовет `install`.
<br><br>
`go run prog.go` - сборка и запуск программы.
<br><br><br>
### Многоточия

`go get github.com/golang/protobuf/...` - многоточие тут означает
"и все дочерние пакеты".
<br>
Это необходимо если в пакет сложный, и содержит подпакеты.
<br>
Для простых достаточно `go get github.com/beevik/ntp`

---

# Точечная сборка

`go build` - команда более низкого уровня, заново компилирующая выбранный пакет.
<br><br>

Например:
```
$ go build -o /tmp/thelib.a github.com/beevik/ntp

$ file /tmp/thelib.a
thelib.a: current ar archive
```

Или:
```
$ go build -o /tmp/prog github.com/golang/protobuf/protoc-gen-go

$ file /tmp/prog 
prog: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
```

Результат сборки зависит от пакета (`main` - executable, любой другой - библиотека)

---

# Кросс-компиляция

Go позволяет легко собирать программы для других архитектур и операционных систем.<br>
Для этого при сборке нужно переопределить переменные `GOARCH` и `GOOS`:

```
$ GOOS=windows go build -o /tmp/prog github.com/golang/protobuf/protoc-gen-go

$ file /tmp/prog
prog: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows

$ GOARCH=386 GOOS=darwin go build -o /tmp/prog github.com/golang/protobuf/protoc-gen-go

$ file /tmp/prog
prog: Mach-O i386 executable
```

Возможные значения `GOOS` и `GOARCH`
[https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)

---

# Единое дерево кода

Для работы в парадигме `GOPATH` нужно:
* Создать *публичный* проект `github.com/username/projectname`
* Скачать проект в `GOPATH` с помощью `go get github.com/username/projectname/...`
* Изменять, компилировать, и комитить проект из `$GOPATH/src/github.com/username/projectname`

Плюсы:
* Простота. Плоская структура
* Отсутствие версий (?) (master должен быть стабилен).

Минусы:
* Отсутствие версий (!)
* Иногда в проекте не только Go-код.
* Неудобно для корпоративных и других непубличных проектов.
  
---

# Go Modules

Начиная с Go 1.11 появилась поддержка модулей - системы версионирования и зависимостей, а также разработки вне `GOPATH`.
<br><br>
Стандартные команды (`go get`, `go install`, `go test` и т.д.) работают по-разному внутри модуля и внутри `GOPATH`.
<br><br>
Модуль - любая директория вне `GOPATH`, содержащая файл `go.mod`

---

# Создание Go модуля

* (Опционально) создайте и склонируйте (в любое место) репозиторий с проектом
```
git clone https://github.com/user/otus-go.git /home/user/otus-go
```
* Создайте внутри репозитория нужные вам директории
```
mkdir /home/user/otus-go/hw-1
```
* Зайдите в директорию и инициализируйте Go модуль
```
cd /home/user/otus-go/hw-1
go mod init github.com/user/otus-go/hw-1
```

Теперь `/home/user/otus-go/hw-1` - это Go модуль.

---

# Добавление зависимостей

Внутри модуля, вы так можете добавить пакет точно так же
```
$ go get github.com/beevik/ntp
go: finding golang.org/x/net latest
```

Но при этом пакет попадет не в `$GOPATH/src`, а в `$GOPATH/pkg/mod`
```
$ tree -L 4 ~/go/pkg
/home/mialinx/go/pkg
└── mod
    ├── cache
    │   └── download
    │       ├── github.com
    │       └── golang.org
    ├── github.com
    │   └── beevik
    │       └── ntp@v0.2.0
    └── golang.org
        └── x
            ├── net@v0.0.0-20190827160401-ba9fcec4b297
            └── sys@v0.0.0-20190215142949-d0b11bdaac8a
```
---

# Версии зависимостей

Внутри Go модуля обязательно находится файл `go.mod`, содержащий информацию о версии и зависимостях

```
$ cat go.mod
module github.com/mialinx/foobar

go 1.14

require (
	github.com/beevik/ntp v0.2.0 // indirect
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
)
```

А так же `go.sum`, содержащий чек-суммы зависимостей.

<br><br>

Внимание (!) версии зависимостей - фиксируются в момент добавления.

---

# Авто-добавление

Есть еще более простой способ управлять зависимостями: просто редактируйте код
```
package main

import (
	"fmt"
	"github.com/go-loremipsum/loremipsum"
)

func main() {
	fmt.Println(loremipsum.New().Word())
}
```
А потом запустите
```
$ go mod tidy 
```
Это добавит новые и удалит неиспользуемые зависимости
---

# Утилиты

.main-image[
  ![img/tools.jpeg](img/tools.jpeg)
]
---

# One-shot запуск

Запустить файл "как скрипт".

```
go run ./path/to/your/snippet.go
```

Удобно для проверки кода и синтаксиса.<br><br>

Так же можно использовать Go PlayGround: [https://play.golang.org/p/Fz3j-hbcocv](https://play.golang.org/p/Fz3j-hbcocv)

---

# Получение справки

```
$ go help
Go is a tool for managing Go source code.

Usage:

	go <command> [arguments]

The commands are:
	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
...

$ go help build
usage: go build [-o output] [-i] [build flags] [packages]

Build compiles the packages named by the import paths,
along with their dependencies, but it does not install the results.
...
```
---


# Форматирование кода

В Go нет style guide, зато есть `go fmt path/to/code.go`
<br><br>
.left-code[
Было:
```
package main
import "fmt"



const msg = "%d students in chat\n"
type Student struct{
	Name string
	Age int
}
func main() {
	for i:=99;i>0;i-- {
		fmt.Printf(msg, i, i)
		if i<10{
			break
		}

	}
}

```
]
.right-code[
Стало:
```
package main

import "fmt"

const msg = "%d students in chat\n"

type Student struct {
	Name string
	Age  int
}

func main() {
	for i := 99; i > 0; i-- {
		fmt.Printf(msg, i, i)
		if i < 10 {
			break
		}

	}
}
```
]
---

# Обновление импортов c помощью `goimports`

```
$ go get golang.org/x/tools/cmd/goimports
$ ~/go/bin/goimports -w path/to/code.go
```

```
import (
	"strings"
)

func main() {
	fmt.Println(loremipsum.New().Word())
}
```

```
import (
	"fmt"

	"github.com/go-loremipsum/loremipsum"
)

func main() {
	fmt.Println(loremipsum.New().Word())
}
```


---

# Линтеры

Линтер - программа, анализирующая код и сообщающая о потенциальных проблемах.
<br><br>
`go vet` - встроенный линтер
```
$ go vet ./run.go
# command-line-arguments
./run.go:14:3: Printf call needs 1 arg but has 2 args

$ echo $?
2
```

`golint` - популярный сторонний линтер
```
$ go get -u golang.org/x/lint/golint

$ ~/go/bin/golint -set_exit_status ./run.go
run.go:7:6: exported type Student should have comment or be unexported
Found 1 lint suggestions; failing.

$ echo $?
1
```

---

# Металинтеры

Металинтеры - обертка, запускающая несколько линтеров за один проход.
<br><br>

<b>Deprecated</b>:<br>
https://github.com/alecthomas/gometalinter
<br><br>

<b>Модно-молодёжно</b>:<br>
https://github.com/golangci/golangci-lint/

```
$ ~/go/bin/golangci-lint run ./run.go
run.go:14:3: printf: Printf call needs 1 arg but has 2 args (govet)
		fmt.Printf(msg, i, i)
		^
run.go:7:6: `Student` is unused (deadcode)
type Student struct {
     ^

$ echo $?
1
```

---

# Как сдавать домашние задания?

https://github.com/OtusGolang/home_work/blob/master/HOWTO.md

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
[https://otus.ru/polls/8448/](https://otus.ru/polls/8448/)
]

.right-image[
![](img/gopher7.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
