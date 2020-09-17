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

# Тестирование. Часть 1

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


* Что там на краю обрыва? (Самое маленькое/большое число, граница, на которой меняется состояние). Граничные условия.


* А что если? Странные сценарии использования.

---

# Знакомтесь, тест в Go

```
strings_test.go // <-- ..._test.go
```

```
func TestIndex(t *testing.T) { // <-- Test...(t *testing.T)
    const s, sub, want = "chicken", "ken", 4
    got := strings.Index(s, sub)
    if got != want {
        t.Errorf("Index(%q,%q) = %v; want %v", s, sub, got, want)
    }
}
```

https://goplay.space/#yybc8Np1JjK

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

https://goplay.space/#vjAsrBrQrxu


---

# testing: практика

## TitleCase
* Делает слова в строке с большой буквы.
* Кроме слов из второй строки.
* Первое слово всегда с большой буквы.

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

https://goplay.space/#PQCd4_FqLeZ

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

https://goplay.space/#5cpT652lEyy

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
https://goplay.space/#YAgI798H8kj

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
		tc := tc
		t.Run(tc.str, func(t *testing.T) {
			got, err := strconv.ParseInt(tc.str, 10, 8)
			require.NoError(t, err)
			require.EqualValues(t, tc.expected, got)
		})
	}
}
```

https://goplay.space/#ssdm9GN6SiO

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

## Элементарные типы данных в Go

<br>
<br>
<br>

## 8 сентября, вторник

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
