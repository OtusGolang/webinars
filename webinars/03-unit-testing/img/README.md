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

# Форматирование данных

### Алексей Бакин

---

# Как проходит занятие

* ### Активно участвуем - задаем вопросы.
* ### Чат вижу - могу ответить не сразу.
* ### После занятия - оффтопик, ответы на любые вопросы.

---

# О чем будем говорить

### 1. Примитивы синхронизации (WaitGroup, Once, Mutex, Cond)
### 2. Модель памяти в Go
### 3. Race-детектор

---

# Цель занятия

.right-image[
![](tmp/gopher.png)
]

#
- Изучить возможности кодирования бинарных данных в текстовой форме
- Научиться использовать стандартную библиотеку для кодирования в формате base64
- Изучить форматы JSON, XML, YAML.
- Изучить подходы к парсингу XML.
- Научиться парсить JSON через стандартную библиотеку
- Изучить библиотеку easyjson
- Изучить библиотеки для работы с MsgPack и Protobuf

---

# А зачем?

.right-image[
![](tmp/gopher.png)
]

### Зачем кодировать бинарные данные в текстовый вид?

---

# Кодировка base64


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

import b64 "encoding/base64"
import "fmt"

func main() {

    data := "Hello world"

    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
```
https://play.golang.org/p/4oFM2M2Sirq

---

# Работа с base64

А какие недостатки?

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
https://play.golang.org/p/GwrvXsSzeN7

---

# Поточная работа с base64 (декодирование)

```
package main

import (
	"encoding/base64"
	"os"
	"io"
	"strings"
)

func main() {
	input := "Zm9vAGJhcg=="
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))
        io.Copy(os.Stdout, r)
}
```

https://play.golang.org/p/uxmmi_OX42i


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

https://play.golang.org/p/p9uRcgPUX8B (полная версия)

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
https://play.golang.org/p/mGVtP-hSQjq

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


https://play.golang.org/p/RxcV-MjmgAm (полная версия)


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

https://play.golang.org/p/QbfwL44vjJU (полная версия)

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
https://play.golang.org/p/FekJkpuj9KT (полная версия)

---

# XML

## А что если данные не помещаются в оперативную память?

---
# SAX Parser

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

https://play.golang.org/p/cuSIsVyZpD-

---

# Бинарные сериализаторы

## Что это такое?
## Какие плюсы?
## А какие минусы?

---

# Бинарные сериализаторы

- gob
- msgpack
- protobuf

---

# msgpack (github.com/vmihailenco/msgpack)

```
type Person struct {
	Name        string
	Surname     string
	Age         uint8
	ChildrenAge map[string]uint8
}
func main() {
	p := Person{Name:  "Ivan",
		Surname: "Remen", Age: 27,
	}
	p.ChildrenAge = make(map[string]uint8)
	p.ChildrenAge["Alex"] = 5
	p.ChildrenAge["Maria"] = 2

	marshaled, _ := msgpack.Marshal(&p)

	fmt.Printf("Length of marshaled: %v
	   IMPL: %v\n", len(marshaled), string(marshaled))

	p2 := &Person{}
	msgpack.Unmarshal(marshaled, p2)
	fmt.Printf("Unmarshled: %v\n", p2)
}
```
https://play.golang.org/p/4pYvh-Qa_wg

---

#protobuf (proto-файл)

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

Сборка: protoc --go_out=. *.proto

---

#Работа с protobuf

```
package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &Person{
		Age:         27,
		Name:        "Ivan",
		Surname:     "Remen",
		ChildrenAge: make(map[string]uint32),
	}
	p.ChildrenAge["Maria"] = 2
	p.ChildrenAge["Alex"] = 5

	marshaled, _ := proto.Marshal(p)

	fmt.Printf("marshaled len %d message = %s\n", len(marshaled), string(marshaled))

	p2 := &Person{}
	proto.Unmarshal(marshaled, p2)

	fmt.Printf("Unmarshaled %v", p2)

}
```

---

# Практика

### Изучаем easyjson

---

# Тест

https://forms.gle/SiDmYTPUU5La3rA88

---


# На занятии

- Изучили возможности кодирования бинарных данных в текстовой форме
- Научились использовать стандартную библиотеку для кодирования в формате base64
- Изучили форматы JSON, XML, YAML.
- Изучили подходы к парсингу XML.
- Научились использовать стандартную библиотеку для кодирования в формате base64
- Научиться парсить JSON через стандартную библиотеку
- Изучили библиотеку easyjson
- Изучили библиотеки для работы с MsgPack и Protobuf

---

## Вопросы?

---

# Опрос

Не заполните заполнить опрос. Ссылка на опрос будет в слаке.

---

class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Спасибо за внимание!
