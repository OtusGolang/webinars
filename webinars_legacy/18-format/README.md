.center.icon[![otus main](../img/main.png)]

---

class: top white
background-image: url(../img/check.svg)
background-size: 130%
.top.icon[![otus main](../img/logo.png)]

.sound-top[
  # Как меня слышно и видно?
]

.sound-bottom[
	## > Напишите в чат
	+ если все хорошо
	- если есть проблемы со звуком или с видео]

---

class: white
background-image: url(../img/message.svg)
.top.icon[![otus main](../img/logo.png)]

# Форматирование данных

### Алексей Бакин

---

# Цель занятия

### Узнать про разные способы форматирования данных и их особенности,
### чтобы уметь выбирать подходящий инструмент для реальной задачи.

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
![](../img/gopher_science.png)
]

---

# А зачем?

.right-image[
![](../img/gopher.png)
]

### Что такое бинарные данные?
### .

.


---

# А зачем?

.right-image[
![](../img/gopher.png)
]

### Что такое бинарные данные?
### Зачем кодировать бинарные данные в текстовый вид?

https://en.wikipedia.org/wiki/Binary-to-text_encoding

---

# Кодировка quoted-printable

https://en.wikipedia.org/wiki/Quoted-printable

```
Привет, мир!
```

↓↓↓↓

```
=D0=9F=D1=80=D0=B8=D0=B2=D0=B5=D1=82, =D0=BC=D0=B8=D1=80!
```

Какие минусы?

<br>

.

---
# Кодировка quoted-printable

https://en.wikipedia.org/wiki/Quoted-printable

```
Привет, мир!
```

↓↓↓↓

```
=D0=9F=D1=80=D0=B8=D0=B2=D0=B5=D1=82, =D0=BC=D0=B8=D1=80!
```

Какие минусы?

<br>

Избыточность = 300%


---

# Кодировка base64

https://en.wikipedia.org/wiki/Base64

```
Man is distinguished, not only by his reason, but by this singular passion
from other animals, which is a lust of the mind, that by a perseverance of
delight in the continued and indefatigable generation of knowledge, exceeds
the short vehemence of any carnal pleasure.
```

```
TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0
aGlzIHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1
c3Qgb2YgdGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZGVsaWdodCBpbiB0
aGUgY29udGludWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd2xlZGdl
LCBleGNlZWRzIHRoZSBzaG9ydCB2ZWhlbWVuY2Ugb2YgYW55IGNhcm5hbCBwbGVhc3VyZS4=
```

Избыточность = 1/3

---

# Работа с base64
```
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hello world"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, err := base64.StdEncoding.DecodeString(sEnc)
	mustNil(err)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	mustNil(err)
	fmt.Println(string(uDec))
}
```
https://goplay.tools/snippet/469rXJqiq0k


---

# Поточная работа с base64 (кодирование)

```
package main

import (
	"encoding/base64"
	"os"
)

func main() {
	input := []byte("foo\x00bar")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	//
	// Зачем закрывать?
	//
	encoder.Close()
}
```

https://goplay.tools/snippet/T2e6bwI3h4g

---

# Поточная работа с base64 (кодирование)

```
package main

import (
	"encoding/base64"
	"os"
)

func main() {
	input := []byte("foo\x00bar")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
}
```

https://goplay.tools/snippet/T2e6bwI3h4g


---

# Поточная работа с base64 (декодирование)

```
package main

import (
	"encoding/base64"
	"io"
	"os"
	"strings"
)

func main() {
	input := "Zm9vAGJhcg=="
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))
	io.Copy(os.Stdout, r)
}

```

https://goplay.tools/snippet/oKiYPu6jfDj


---

# Текстовая сериализация
## JSON
## XML
## YAML

Какие цели у сериализации?

---

# JSON

```
{"widget": {
    "debug": "on",
    "window": {
        "title": "Sample Konfabulator Widget",
        "name": "main_window",
        "width": 500,
        "height": 500
    },
    "image": {
        "src": "Images/Sun.png",
        "name": "sun1",
        "hOffset": 250,
        "vOffset": 250,
        "alignment": "center"
    },
    "text": {
        "data": "Click Here",
        "size": 36,
        "style": "bold",
        "name": "text1",
        "hOffset": 250,
        "vOffset": 100,
        "alignment": "center",
        "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
    }
}}
```

---

# YAML

```
---
widget:
  debug: 'on'
  window:
    title: Sample Konfabulator Widget
    name: main_window
    width: 500
    height: 500
  image:
    src: Images/Sun.png
    name: sun1
    hOffset: 250
    vOffset: 250
    alignment: center
  text:
    data: Click Here
    size: 36
    style: bold
    name: text1
    hOffset: 250
    vOffset: 100
    alignment: center
    onMouseUp: sun1.opacity = (sun1.opacity / 100) * 90;
```

---

# XML

```
<widget>
    <debug>on</debug>
    <window title="Sample Konfabulator Widget">
        <name>main_window</name>
        <width>500</width>
        <height>500</height>
    </window>
    <image src="Images/Sun.png" name="sun1">
        <hOffset>250</hOffset>
        <vOffset>250</vOffset>
        <alignment>center</alignment>
    </image>
    <text data="Click Here" size="36" style="bold">
        <name>text1</name>
        <hOffset>250</hOffset>
        <vOffset>100</vOffset>
        <alignment>center</alignment>
        <onMouseUp>
            sun1.opacity = (sun1.opacity / 100) * 90;
        </onMouseUp>
    </text>
</widget>
```

---

# Работа с JSON (базовые возможности)

```
func main() {
	p1 := &Person{
		Name: "Vasya",
		age: 36,
		Job: struct {
			Department string
			Title      string
		}{Department: "Operations", Title: "Boss"},
	}

	j, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("p1 json %s\n", j)

	var p2 Person
	json.Unmarshal(j, &p2)
	fmt.Printf("p2: %v\n", p2)

}
```

https://goplay.tools/snippet/EwX5Dq2l60C (полная версия)

---

# Работа с JSON через interface{}

```
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	j := []byte(`{"Name":"Vasya",
		"Job":{"Department":"Operations","Title":"Boss"}}`)

	var p2 interface{}
	json.Unmarshal(j, &p2)
	fmt.Printf("p2: %v\n", p2)

	person := p2.(map[string]interface{})
	fmt.Printf("name=%s\n", person["Name"])
}
```

https://goplay.tools/snippet/8SgUuo2L23z

---

# Работа с тегами пакета json

```
type Person struct {
	Name    string `json:"fullname,omitempty"`
	Surname string `json:"familyname,omitempty"`
	Age     int    `json:"-"`
	Job     struct {
		Department string
		Title      string
	}
}
```

https://goplay.tools/snippet/RxcV-MjmgAm (полная версия)


---


# Быстрый json

https://github.com/mailru/easyjson

https://jsoniter.com/

---

# Кодирование в xml

```
type Address struct {
	City, State string
}
type Person struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}
```

https://goplay.tools/snippet/xCtXEHUgKAU (полная версия)

---

# Декодирование из xml

```
type Address struct {
	City, State string
}
type Person struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}
```

https://goplay.tools/snippet/FekJkpuj9KT (полная версия)

---

# XML

## А что если данные не помещаются в оперативную память?

---
# SAX Parser

https://en.wikipedia.org/wiki/Simple_API_for_XML

```
for {
		token, _ := decoder.Token()

		switch se := token.(type) {
		case xml.StartElement:
			fmt.Printf("Start element: %v Attr %s\n",
						se.Name.Local, se.Attr)
			inFullName = se.Name.Local == "FullName"
		case xml.EndElement:
			fmt.Printf("End element: %v\n", se.Name.Local)
		case xml.CharData:
			fmt.Printf("Data element: %v\n", string(se))
			if inFullName {
				names = append(names, string(se))
			}
		default:
			fmt.Printf("Unhandled element: %T", se)
		}

	}
```


https://goplay.tools/snippet/ADbUEs1PeUF

---

# Бинарные сериализаторы

## Что это такое?
## Какие плюсы?
## А какие минусы?

---

# Бинарные сериализаторы

- gob (https://golang.org/pkg/encoding/gob/)
- msgpack
- Protobuf

---

# msgpack (<a href="https://github.com/vmihailenco/msgpack">github.com/vmihailenco/msgpack</a>)

```
type Person struct {
	Name        string
	Surname     string
	Age         uint8
	ChildrenAge map[string]uint8
}

func main() {
	p := Person{Name: "Ivan", Surname: "Remen", Age: 27}
	p.ChildrenAge = make(map[string]uint8)
	p.ChildrenAge["Alex"] = 5
	p.ChildrenAge["Maria"] = 2

	marshaled, _ := msgpack.Marshal(&p)

	fmt.Printf("Length of marshaled: %v IMPL: %v\n", len(marshaled), string(marshaled))

	p2 := &Person{}
	msgpack.Unmarshal(marshaled, p2)
	fmt.Printf("Unmarshled: %v\n", p2)
}
```

https://goplay.tools/snippet/HFmhdDcmnx0

---

# Protobuf

```
syntax = "proto3";

package main;

message Person {
    string name = 1;
    string surname = 2;
    uint32 age = 3;

    map<string, uint32> children_age = 4;
}
```

#### Установка:
https://developers.google.com/protocol-buffers/docs/downloads
+
```
go get google.golang.org/protobuf/cmd/protoc-gen-go
```

#### Сборка:
```
protoc --go_out=. *.proto
```

---


# Повторение

.left-text[
Давайте проверим, что вы узнали за этот урок, а над чем стоит еще поработать.
<br><br>
Ссылка в чате
]

.right-image[
![](../img/gopher_science.png)
]

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
Ссылка в чате.
]

.right-image[
![](../img/gopher_boat.png)
]

---

# Следующее занятие

## Взаимодействие с OS

<br>
<br>
<br>

## 03 февраля, четверг

---

class: white
background-image: url(../img/message.svg)
.top.icon[![otus main](../img/logo.png)]

# Спасибо за внимание!
