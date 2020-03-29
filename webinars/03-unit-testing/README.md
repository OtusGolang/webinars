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

# Тестирование

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем - задаем вопросы.
* ### Чат вижу - могу ответить не сразу.
* ### После занятия - оффтопик, ответы на любые вопросы.

---

# О чем будем говорить

### 1. Зачем нужны тесты?
### 2. testing и testify.
### 3. Приемы тестирования.

---

# Зачем?

### Зачем нужны тесты?

---

# Зачем?

* Упрощают рефакторинг (safety net)


* Документируют код


* Отделение интерфейса от реализации (mocks), менее связный код


* Помогают найти неактуальный код


* Помогают найти новые кейсы


* Метрика для менеджмента (покрытие)


* Определяют контракт


* Качество кода

---

# Думай, как тестировщик

* Как хотелось бы, чтобы работало? (На что это похоже? Как бы я мог это использовать?) Не лазить в кишки.


* Как не должно работать? (Неправильные параметры, неправильный порядок вызовов) Негативные тест-кейзы.


* Граничные условия. (Самое маленькое/большое число, граница, на которой меняется состояние).


* А что если? Странные сценарии использования.

---

# Знакомтесь, тест в Go

```
strings_test.go // <-- ..._test.go
```

```
func TestIndex(t *testing.T) { // <-- Test...(t *testing.T)
    const s, sep, want = "chicken", "ken", 4
    got := strings.Index(s, sep)
    if got != want {
        t.Errorf("Index(%q,%q) = %v; want %v", s, sep, got, want)
    }
}
```

https://play.golang.org/p/0G3efzky6L0

---

# testing: Error vs Fatal

```
func TestAtoi(t *testing.T) {
	const str, want = "42", 42
	got, err := strconv.Atoi(str)
	if err != nil {
		t.Errorf("strconv.Atoi(%q) returns unexpeted error: %v", str, err)
	}
	if got != want {
		t.Errorf("strconv.Atoi(%q) = %v; want %v", str, got, want)
	}
}
```

https://play.golang.org/p/vjAsrBrQrxu


---

# testing: практика

## TitleCase
* Делает слова в строке с большой буквы.
* Кроме слов из второй строки.
* Первое слово всегда с первой буквы.

Пример:

<br/>

* `TitleCase("the quick fox in the bag", "")` -> `"The Quick Fox In The Bag"`
* `TitleCase("the quick fox in the bag", "in the")` -> `"The Quick Fox in the Bag"`

---

# testing: практика

## Задание
* ### Дописать существующие тесты.
* ### Придумать один новый тест.
* ### Не закрывайте playground - еще пригодится :)

https://play.golang.org/p/PQCd4_FqLeZ

---

# testify

https://github.com/stretchr/testify

```
func TestAtoi(t *testing.T) {
	const str, want = "42", 42
	got, err := strconv.Atoi(str)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
```

https://play.golang.org/p/5cpT652lEyy

---

# testify: assert vs require

## Простое правило - всегда используйте require.

---

# testify: изучение API

* ### `require.Equal()` vs `require.Equalf()`
* ### ` `
* ### ` `
* ### ` `

---

# testify: изучение API

* ### `require.Equal()` vs `require.Equalf()`
* ### `require.True(t, err == nil)`
* ### ` `
* ### ` `

---

# testify: изучение API

* ### `require.Equal()` vs `require.Equalf()`
* ### `require.True(t, err == nil)`
* ### `require.Nil(t, err)`
* ### ` `

---

# testify: изучение API

* ### `require.Equal()` vs `require.Equalf()`
* ### `require.True(t, err == nil)`
* ### `require.Nil(t, err)`
* ### `require.NoError(t, err)`

---

# testify: практика

## Задание
* ### Переписать тесты на testify.
* ### Не закрывайте playground - еще пригодится :)

---

# Табличные тесты

```
func TestParseInt(t *testing.T) {
	tests := []struct {
		str      string
		expected int64
	}{
		{"-128", -128},
		{"0", 0},
		{"127", 127},
	}

	for _, tc := range tests {
		got, err := strconv.ParseInt(tc.str, 10, 8)
		require.NoError(t, err)
		require.EqualValues(t, tc.expected, got)
	}
}

func TestParseIntErrors(t *testing.T) {
	tests := []string{"-129", "128", "byaka"}

	for _, str := range tests {
		_, err := strconv.ParseInt(str, 10, 8)
		require.Error(t, err)
	}
}
```
https://play.golang.org/p/YAgI798H8kj

---

# Табличные тесты: t.Run

```
func TestParseInt(t *testing.T) {
	tests := []struct {
		str      string
		expected int64
	}{
		{"-128", -128},
		{"0", 0},
		{"127", 127},
	}

	for _, tc := range tests {
		t.Run(tc.str, func(t *testing.T) {
			got, err := strconv.ParseInt(tc.str, 10, 8)
			require.NoError(t, err)
			require.EqualValues(t, tc.expected, got)
		})
	}
}
```

https://play.golang.org/p/Rr-i2UFYAXH

---

# Табличные тесты: практика

## Задание
* ### Переписать тесты на табличные.
* ### Постараться придумать еще один тест.
* ### Можно закрывать playground :)

---

# Blackbox тесты

```
package strings

func Contains(s, sub string) bool {
	...
}

func contains(s, p, sub, string) bool {
	...
}
```

```
package strings_test

import (
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	strings.Contains(...)
}
```
---

# Как запускать тесты

Все тесты в пакете и подпакетах:
```
go test ./...

go test ./pkg1/...

go test github.com/otus/superapp/...
```

Конкретные тесты по имени:
```
go test -run TestFoo
```

По тегам (`// +build integration`):
```
go test -tags=integration
```
---

# Coverage

* ### `go test -cover` - посмотреть покрытие
* ### `go test -coverprofile=c.out` - записать отчет о покрытии
* ### `go tool cover -html=c.out` - посмотреть отчет о покрытии

https://blog.golang.org/cover

---

# Моки: DI

```
type PersonRepo struct {
	db DB
}

func (r *PersonRepo) ProcessAndStore(person Person) {
	// ...
}
```

```
func main() {
	repo := PersonRepo {
		db: NewDB(connstr),
	}
}
```

```
func TestProcessAndStore() {
	repo := PersonRepo {
		db: NewMockDB(),
	}
	repo.ProcessAndStore(testPerson)

}
```

---

# Моки

* ### https://pkg.go.dev/github.com/stretchr/testify/mock?tab=doc
* ### https://github.com/golang/mock

---

# Моки: отличие от стабов

* ### Стаб - это заглушка, "пустая" реализация интерфейса.
* ### Мок проверяет правильность вызовов интерфейса.

---

# testify: suite

https://pkg.go.dev/github.com/stretchr/testify/suite?tab=doc

```
type UsersTestSuite struct {
    suite.Suite
	db UsersDB
}

func (s *UsersTestSuite) SetupTest() {
	s.db = NewUsersDB(connstr)
}

func (s *UsersTestSuite) TestAddUser() {
	user1 := fakeUser()
	s.db.AddUser(user1)
	u, err := s.db.FindUser(user1.id)
    s.Require().NoError(err)
    s.Require().Equal(user1, u)
}

func UsersTestSuite(t *testing.T) {
    suite.Run(t, new(UsersTestSuite))
}
```

---

# faker

https://github.com/bxcodec/faker

```
type Person struct {
	Name   string `faker:"username"`
	Phone  string `faker:"phone_number"`
	Answer int64  `faker:"answer"`
}

func CustomGenerator() {
	faker.AddProvider("answer", func(v reflect.Value) (interface{}, error) {
		return int64(42), nil
	})
}

func main() {
	CustomGenerator()

	var p Person
	faker.FakeData(&p)
	fmt.Printf("%+v\n", p)
}
```

https://play.golang.org/p/BqOcrrUCZAn

---

# Golden files

https://medium.com/soon-london/testing-with-golden-files-in-go-7fccc71c43d3

```
var update = flag.Bool(“update”, false, “update golden files”)

func TestSomething(t *testing.T) {
	actual := doSomething()
	golden := filepath.Join(“test-fixtures”, ”expected.golden”)
	if *update {
		ioutil.WriteFile(golden, actual, 0644)
	}
    expected, _ := ioutil.ReadFile(golden)
    require.Equal(t, expected, actual)
}
```
---

# Тестирование мутированием

https://github.com/zimmski/go-mutesting

---

# assert vs require

## В чем проблема?

```
resp, err := client.Do(req)
assert.NoError(t, err)
assert.Equal(t, len(expectedBody), resp.ContentLength)
```

---

# assert vs require

## Простое правило - всегда используйте require.

```
resp, err := client.Do(req)
require.NoError(t, err)
require.Equal(t, len(expectedBody), resp.ContentLength)
```

---

# assert vs require

## В чем проблема?

```
for _, c := range cases {
	resp, err := client.Do(c.req)
	require.NoError(t, err)
	require.Equal(t, len(c.expectedBody), resp.ContentLength)
}
```

---

# assert vs require

```
for _, c := range cases {
	resp, err := client.Do(c.req)
	if assert.NoError(t, err) {
		assert.Equal(t, len(c.expectedBody), resp.ContentLength)
	}
}
```

```
--- FAIL: TestFoo (0.00s)

	Error Trace:	foo_test.go:37
	Error:      	Expected nil, but got: &errors.errorString{s:"test"}

	Error Trace:	foo_test.go:38
	Error:      	Not equal:
	            	expected: 2
	            	actual: 0

	Error Trace:	foo_test.go:38
	Error:      	Not equal:
	            	expected: 3
	            	actual: 0
```

---

# assert vs require

```
for _, c := range cases {
	t.Run(c.name, func(t *testing.T) {
		resp, err := client.Do(c.req)
		require.NoError(t, err)
		require.Equal(t, len(c.expectedBody), resp.ContentLength)
	}
}
```

```
--- FAIL: TestFoo (0.00s)
    --- FAIL: TestFoo/check_empty (0.00s)
	Error Trace:	foo_test.go:38
	Error:      	Expected nil, but got: &errors.errorString{s:"test"}

    --- FAIL: TestFoo/check_v1 (0.00s)
	Error Trace:	foo_test.go:39
	Error:      	Not equal:
	            	expected: 2
	            	actual: 0

    --- FAIL: TestFoo/check_V1 (0.00s)
	Error Trace:	foo_test.go:39
	Error:      	Not equal:
	            	expected: 3
	            	actual: 0
```

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
