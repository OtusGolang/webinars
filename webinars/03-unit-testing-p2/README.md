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
	- если есть проблемы со звуком или с видео]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Тестирование. Часть 2

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем — задаем вопросы.
* ### Чат вижу — могу ответить не сразу.
* ### После занятия — оффтопик, ответы на любые вопросы.

---

# О чем будем говорить

* ### Blackbox тесты
* ### Дубли (моки для интерфейсов, time, фс)
* ### Faker (сгенерированные данные)
* ### Тестирование мутированием
* ### Golden files

---

# Blackbox тесты

```
package router

func New(...) *Router {
	return &Router{
		// ...
	}
}

func (r *Router) run() { // <--- приватный метод будет недоступен
	// ...
}
```

```
package router_test

import (
	"strings"

	// тестируемый пакет импортируется
	"github.com/kulti/task-list/server/internal/router"
)

func TestRouter(t *testing.T) {
	r := router.New(...)
	// ....
}
```
---

# Моки: DI

```
type UserStore struct {
	db UsersDB
}

func (s *UserStore) Duplicate(userID string) (string, error) {
	// ...
}
```

```
func main() {
	repo := NewUserStore(NewDB(connstr))
}
```

```
func TestProcessAndStore() {
	// ...
	store := NewUserStore(mockDB)
	// ...
	newID, err := store.Duplicate(user1.ID)
}
```

---

# Моки: отличие от стабов

* ### Стаб — это заглушка, "простая" реализация интерфейса (обычно хранит состояние).
* ### Мок фиксирует вызовы интерфейса. Позволяет проверить правильность его использования.


https://martinfowler.com/articles/mocksArentStubs.html

---

# Моки: пакеты

* ### https://pkg.go.dev/github.com/stretchr/testify/mock
* ### https://github.com/golang/mock

---

# testify: suite

https://pkg.go.dev/github.com/stretchr/testify/suite

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

https://goplay.tools/snippet/BqOcrrUCZAn

---

# faker: seed

Чтобы в будущем было проще разбираться с упавшими тестами, выставляем seed явно:
```
	var seed int64 = time.Now().UnixNano()
	s.T().Logf("rand seed: %d", seed)
	rand.Seed(seed)
	s.genFakeData()
```

Когда тесты упадут, то увидим в логах:
```
--- FAIL: TestStoreSuire (0.00s)
    --- FAIL: TestStoreSuire/TestDuplicate (0.00s)
        store_test.go:45: rand seed: 1599764658164627786
```

И сможем воспроизвести тест:
```
	...
	var seed int64 = 1599764658164627786 //time.Now().UnixNano()
	...
```

---

# faker: валидация случайных данных

Случайные данные могут оказаться "невалидными". Частые кейзы:
- дубликаты
- зарезервированные слова

```
	argNames := []string{faker.Word(), faker.Word()}
	for isLuaReservedWord(argNames[0]) {
		argNames[0] = faker.Word()
	}
	for argNames[0] == argNames[1] || isLuaReservedWord(argNames[1]) {
		argNames[1] = faker.Word()
	}
```

Или включить уникальные имена внутри самого faker:
```
	faker.SetGenerateUniqueValues(true)
	...
	faker.SetGenerateUniqueValues(false) // Не забываем отключать!
	faker.ResetUnique()
```

---

# Тестирование мутированием

* ### https://github.com/zimmski/go-mutesting

---

# Моки для времени

* ### https://github.com/cabify/timex
* ### https://github.com/benbjohnson/clock

---

# Стаб для ФС

* ### https://github.com/spf13/afero

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

# Рефлексия

.left-text[
Ответьте, пожалуйста, на несколько вопросов.
<br><br>
Они помогут вспомнить и лучше запомнить, что было на занятии.
<br><br>
Ссылка в чате
]

.right-image[
![](img/gopher_meditation.png)
]

---

# Примеры с занятия

* ### https://github.com/OtusGolang/webinars_practical_part/tree/master/03-unit-testing


---

# Результаты первого модуля "Начало работы с Go"

.left-text[
Пройдите, пожалуйста, тест по первому модулю.
<br><br>
Он поможет вспомнить, что мы уже прошли, и покажет, какие области запомнились лучше, а какие хуже.
<br><br>
Ссылка в чате
]

.right-image[
![](img/gopher_science.png)
]

---

# Следующее занятие

## Горутины и каналы

<br>
<br>
<br>

## 20 июля, вторник

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
