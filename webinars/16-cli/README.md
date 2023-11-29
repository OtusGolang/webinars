background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[CLI]
.tutor[Романовский Алексей]
.tutor_desc[Разработчик в Resolver Inc.]

---

background-image: url(../img/rules.svg)

---
# цели занятия

после занятия вы сможете:
работать с операционной системой из программы на Go.

---

# Краткое содержание

* обработка аргументов командной строки: flags, pflag, cobra;
* работа с переменными окружения;
* запуск внешних программ;
* создание временных файлов;
* обработка сигналов.

---

# соглащения и стандартны на CLI

* POSIX: https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html

```
utility_name[-a][-b][-c option_argument][-d|-e][-f[option_argument]][operand...]
```

* GNU: https://www.gnu.org/prep/standards/standards.html#Command_002dLine-Interfaces

```
utility_name -v
utility_name --verbose
```

---

# flag/pflag

* https://golang.org/pkg/flag/
* https://github.com/spf13/pflag

(пример на сл слайде)

---

```go
func main() {
	var msg, cfg string

	//verbose := flag.Bool("verbose", false, "verbose output")
	verbose := flag.Bool("verbose", "v", false, "verbose output")

	//flag.StringVar(&cfg, "cfg", "config.yaml", "config file")
	pflag.StringVar(&cfg, "cfg", "config.yaml", "config file")

	pflag.Parse() // flag.Parse() 

	if *verbose {
		fmt.Println("you say:", msg)
    } else {
		fmt.Println(msg, cfg)
	}
}
```

---

# pflag: флаги без значений

```
pflag.StringVar(&msg, "msg", "hello", "message to print")
pflag.Lookup("msg").NoOptDefVal = "bye"
```
<br><br>

.left-text[
|Флаг            |Значение
|:---------------|:-------
|--port=9999     |ip=9999
|--port	         |ip=80
|[nothing]	     |ip=8080
]

---

# Сложные CLI приложения

```sh
git commit -m 123

docker pull

aws s3 ls s3://bucket-name
```

* [Философия и best practices](https://clig.dev/)
* https://github.com/spf13/cobra/
* https://github.com/urfave/cli

---

# Демонстрация кода (cobra, cobra-cli)

---

# Переменные окружения (дополнительно)
(изучим это в IDE)

```go
env := os.Environ() // слайс строк "key=value"
fmt.Println(env[0]) // USER=rob

user, ok := os.LookupEnv("USER")
fmt.Println(user) // rob

os.Setenv("PASSWORD", "qwe123")                     // установить
os.Unsetenv("PASSWORD")                             // удалить (для новых процессов)
fmt.Println(os.ExpandEnv("$USER lives in ${CITY}")) // "шаблонизация"
```

---

# Запуск внешних программ
`os/exec`, (изучим это в IDE)

```go
cmd := exec.Command("git", "commit", "-am", "fix")
```
```go
type Cmd struct {
  Path string   // Путь к запускаемой программе
  Args []string // Аргументы командной строки
  Env  []string // Переменные окружения ("key=value")
  Dir  string   // Рабочая директория

  // Поток ввода, вывода и ошибок для программы (/dev/null если nil!)
  Stdin io.Reader
  Stdout io.Writer
  Stderr io.Writer
  ...
}
```

---

```
  cmd.CombinedOutput() // ждёт, отдаёт весь вывод как массив, обрабатывает код результата
  cmd.Run() // ждёт ответа, но ничего не делает с потоками
  cmd.Start() // не ждёт. Должны вызывть Wait()
```

---

# Сигналы

Сигналы - механизм OS, позволяющий посылать уведомления программе в особых ситуациях.
<br><br>

|Сигнал |Поведение|Применение|
|:------|:---------------------|:----------|
|SIGINT |Завершить|`Ctrl+C` в консоли|
|SIGKILL|Завершить|`kill -9`, остановка зависших программ|
|SIGHUP |Завершить|Сигнал для переоткрытия логов и перечитывания конфига|
|SIGUSR1|         |На усмотрение пользователя|
|SIGUSR2|         |На усмотрение пользователя|
|SIGPIPE|Завершить|Отправляется при записи в закрытый файловый дескриптор|
|SIGSTOP|Остановить|При использовании отладчика|
|SIGCONT|Продолжить|При использовании отладчика|

<br><br>
Некоторые сигналы, например `SIGINT`, `SIGUSR1`, `SIGHUP`, можно игнорировать или установить обработчик.
<br><br>
Некоторые, например `SIGKILL`, обработать нельзя.
---

# Обработка сигналов

```go
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)
	signal.Ignore(syscall.SIGTERM)

	for s := range c {
		fmt.Println("Got signal:", s)
	}
}
```

---

# Работа с файловой системой

В пакете `os` содержится большое количество функций для работы с файловой системой.

```go
// изменить права доступа к файлу
func Chmod(name string, mode FileMode) error

// изменить владельца
func Chown(name string, uid, gid int) error

// создать директорию
func Mkdir(name string, perm FileMode) error

// создать директорию (вместе с родительскими)
func MkdirAll(path string, perm FileMode) error

// переименовать файл/директорию
func Rename(oldpath, newpath string) error

// удалить файл (пустую директорию)
func Remove(name string) error

// удалить рекурсивно rm -rf
func RemoveAll(path string) error

// создать временный файл со случайным именем
func CreateTemp("dir", "prefix") (*os.File, error)
```

---

# Временные файлы: safefile

https://github.com/dchest/safefile

---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[1 января]
.announce_topic[Рефлексия]

---
background-image: url(../img/thanks.svg)

.tutor[Лектор]
.tutor_desc[Должность]
