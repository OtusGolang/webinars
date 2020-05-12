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

# Взаимодействие с OS

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем - задаем вопросы.
* ### Чат вижу - могу ответить не сразу.
* ### После занятия - оффтопик, ответы на любые вопросы.

---

# О чем будем говорить

.big-list[
* Обработка аргументов командной строки: flag, pflag, cobra
* Работа с переменными окружения
* Запуск внешних программ
* Работа с файловой системой
* Временные файлы
* Обработка сигналов
]

---

# Цель занятия

## Узнать, какие средства <br> взаимодействия с ОС есть в Go.

---

# flag

https://golang.org/pkg/flag/

```
func main() {
	var msg string

	verbose := flag.Bool("verbose", false, "verbose output")
	flag.StringVar(&msg, "msg", "hello world", "message to print")

	flag.Parse()

	if *verbose {
		fmt.Println("you say:", msg)
	} else {
		fmt.Println(msg)
	}
}
```

---

# pflag

https://github.com/spf13/pflag

```
func main() {
	var msg string

	verbose := pflag.BoolP("verbose", "v", false, "verbose output")
	pflag.StringVar(&msg, "msg", "hello world", "message to print")

	pflag.Parse()

	if *verbose {
		fmt.Println("you say:", msg)
	} else {
		fmt.Println(msg)
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

# cobra: сложные CLI приложения

```
git commit -m 123

docker pull

aws s3 ls s3://bucket-name
```

https://github.com/spf13/cobra/

---

# cobra: пример

```
$ cobra init git --pkg-name github.com/otus/git
$ cd git
$ gomod init github.com/otus/git
$ cobra add commit
```

Обычный флаг для команды:
```
var commitFlags struct {
	message string
}
```

```
commitCmd.Flags().StringVarP(&commitFlags.message, "message", "m", "", "commit message")
```

Общий флаг для команды и всех подкоманд:
```
rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/...yaml)")
```

---

# Переменные окружения

```
$ env

$ GOMAXPROCS=4 goapp

$ env -i NEWVAR=val prog
```

---

# Переменные окружения

```
env := os.Environ() // слайс строк "key=value"
fmt.Println(env[0]) // USER=rob

user, ok := os.LookupEnv("USER")
fmt.Println(user) // rob

os.Setenv("PASSWORD", "qwe123")                     // установить
os.Unsetenv("PASSWORD")                             // удалить
fmt.Println(os.ExpandEnv("$USER lives in ${CITY}")) // "шаблонизация"
```

---

# Запуск внешних программ

Для запуска внешних команд используется пакет `os/exec`. Основной тип - `Cmd`.

```
cmd := exec.Command("git", "commit", "-am", "fix")
```
```
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

# Запуск внешних программ

```
cmd := exec.Command("sleep", "1")

err := cmd.Run()
if err != nil {
  log.Fatal(err)
}
```
`cmd.Run()` = `cmd.Start()` + `cmd.Wait()`
```
cmd := exec.Command("sleep", "1")

err := cmd.Start()
if err != nil {
  log.Fatal(err)
}

log.Printf("Waiting for command to finish...")
err = cmd.Wait() // ошибка выполнения
if err != nil {
  log.Fatal(err)
}
```

---

# Запуск внешних программ: ввод/вывод

```
func main() {
	cmd := exec.Command("../env/env")
	cmd.Env = append(os.Environ(),
		"USER=petya",
		"CITY=SPb",
	)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
```

---

# Запуск внешних программ: ввод/вывод

```
func main() {
	cmd := exec.Command("../env/env")
	cmd.Env = append(os.Environ(),
		"USER=petya",
		"CITY=SPb",
	)
	cmd.Stdout = os.Stdout // <===

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
```

---

# Запуск внешних программ: pipe

```
ls | wc -l
```
<br>
```
lsCmd := exec.Command("ls")
wcCmd := exec.Command("wc", "-l")

pipe, _ := lsCmd.StdoutPipe()
wcCmd.Stdin = pipe
wcCmd.Stdout = os.Stdout

_ = lsCmd.Start()
_ = wcCmd.Start()
_ = lsCmd.Wait()
_ = wcCmd.Wait()
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

```
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

```
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
```

---


# Временные файлы

```
func main() {
	content := []byte("temporary file's content")

	tmpfile, err := ioutil.TempFile("", "example.")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
```

---

# Временные файлы: safefile

https://github.com/dchest/safefile

---

# Опрос

.left-text[
Заполните пожалуйста опрос.
<br><br>
Ссылка в чате.
]

.right-image[
![](img/gopher.png)
]

---

# Домашнее задание

Реализовать утилиту envdir на Go.
<br><br>

Она позволяет запускать программы получая переменные окружения из определенной директории.
<br><br>
Пример использования:

```
go-envdir /path/to/env/dir some_prog
```

Если в директории /path/to/env/dir содержатся файлы
* `A_ENV` с содержимым `123`
* `B_VAR` с содержимым `another_val`

То программа `some_prog` должать быть запущена с переменными окружения `A_ENV=123 B_VAR=another_val`

https://github.com/OtusGolang/home_work/tree/master/hw08_envdir_tool

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!
