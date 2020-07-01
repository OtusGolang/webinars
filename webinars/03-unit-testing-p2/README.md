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

# Тестирование. Часть 2

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем - задаем вопросы.
* ### Чат вижу - могу ответить не сразу.
* ### После занятия - оффтопик, ответы на любые вопросы.

---

# О чем будем говорить

* ### Моки (интерфейсы, time, фс)
* ### Faker
* ### Тестирование мутирвоанием
* ### Golden files

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

# Моки

* ### https://pkg.go.dev/github.com/stretchr/testify/mock?tab=doc
* ### https://github.com/golang/mock

---

# Моки: отличие от стабов

* ### Стаб - это заглушка, "пустая" реализация интерфейса.
* ### Мок фиксирует вызовы интерфейса. Позволяет проверить правильность его использования.

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

https://goplay.space/#BqOcrrUCZAn

---

# Тестирование мутированием

* ### https://github.com/zimmski/go-mutesting

<br>

Устарел. Поддержка модулей в ПРе - https://github.com/zimmski/go-mutesting/pull/77.

---

# Моки для времени

* ### https://github.com/cabify/timex
* ### https://github.com/facebookarchive/clock

---

# Моки для ФС

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
![](img/gopher_science.png)
]

---

# Примеры с занятия

* ### https://github.com/OtusGolang/webinars_practical_part/tree/master/03-unit-testing

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
