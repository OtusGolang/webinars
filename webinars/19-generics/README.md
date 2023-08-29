background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Тема]
.tutor[Лектор]
.tutor_desc[Должность]

---

background-image: url(../img/rules.svg)

---

# О чем будем говорить:
- Что такое дженерики
- Внутреннее устройство
- Базовые типы в контексте дженериков
- Интерфейсы в контексте дженериков
- Примеры использования

---

# Дженерики
- Дженерики - это инструмент обобщенного программирования, которые позволяет писать обощенный код для разных типов данных
без необходимости его дублирования или использования интерфейсов. По сути это такой placeholder для типов,
в который можно подставить нужный тип при разработке.

---

# Дженерики в Golang

- Появились в go1.18
- Продолжают развиваться
- Примеры актуальны для go1.21

---

# Особенности дженериков

- Влияют на время компиляции
- Поддерживают строгую типизацию
- Компилятор проверяет соотвествие типов заявленному интерфейсу
- Для базовых типов работают как обычные функции без дженериков

---

# Функция без дженериков

```

func IMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func SMax(a, b string) string {
    if a > b {
        return a
    }
    return b
}

```

---

# Функция с дженериком

```
func GMax[T interface{ string | int }](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```

https://go.dev/play/p/xKo8N-LwcOJ

---

# Дженерики как метатипы

```
type Numbers interface {
	int | int32 | int64 | float64
}

func NMax[T Numbers](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```

---

# Производные типы и дженерики

```
type Numbers interface {
	int
}

func NMax[T Numbers](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```

```go
type Price int
NMax(Price(2), Price(3))
```

https://go.dev/play/p/-lFCCr64Gsf

---


# Производные типы и дженерики

```
type Numbers interface {
	~int
}
```

https://go.dev/play/p/1tZ81nrFX_U

---

# Стандартные метатипы

- any (alias: interface{})
- comparable (для hash-map)
- cmp.Ordered

---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[1 января]
.announce_topic[Тема следующего вебинара]

---
background-image: url(../img/thanks.svg)

.tutor[Лектор]
.tutor_desc[Должность]
