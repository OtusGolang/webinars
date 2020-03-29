.center.icon[![otus main](https://drive.google.com/uc?id=1NPIi9Hw5ZjA5SK24lTXckDjNAPSuFAHi)]

---

class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Методический воркшоп
### Александр Коржиков

---

class: top white
background-image: url(tmp/sound.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

.sound-top[
  # Как меня слышно и видно?
]

.sound-bottom[
  ## > Напишите в чат
  ### **+** если все хорошо
  ### **–** если есть проблемы cо звуком или с видео
]

---

# Темы

- Presentations
  - Google Slides Overview
  - Working with Markdown Notes
- Ideas
  - Arrange Materials with Folders Structure
  - Slackbot Notifications with Messages Materials

---

# Код

## Заголовок первого уровня
### Заголовок второго уровня
#### Пример кода

```
var a = 123
console.log('Hello World')
```

---

# Customized built-in elements

- reuse && extend - встроенное поведение
- extends && is - обязательные аттрибуты

```
class PlasticButton extends HTMLButtonElement {
  constructor() {
    super() // ... 
  }
}

customElements.define("plastic-button",
  PlasticButton, { extends: "button" }
)

document.createElement("button", {
  is: "plastic-button"
}) 
```

```html
<button is="plastic-button">Click Me!</button>
```

---

class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Modules Q&A

---

# Вопрос

### А Вы знаете какие модули включены в стандартный дистрибутив `Node`?

<br>

.center[
  ![question](https://drive.google.com/uc?id=1mTHfs6OkZ95ypZAOrbk3snF4Uv0n-f42)
]

---

# На занятии

- Попрактиковали `getters / setters`
- Разобрали веб спецификацию `Custom Elements`

---

# Table

| Header1 | Header2 |
|:---:|:---|
|text1|text2|

---

# Самостоятельная работа

![no homework](https://drive.google.com/uc?id=1ThBt-xuvffWmX7qVaG4QLOXCW38qvtf7)

---

class: white
background-image: url(tmp/title.svg)
.top.icon[![otus main](https://drive.google.com/uc?id=18Jw9bQvL3KHfhGWNjqyQ3ihR3fV3tmk8)]

# Спасибо за внимание!

### Вы верите в Markdown?

.black[
### Пожалуйста, пройдите [опрос](https://otus.ru/polls/2389/) в личном кабинете
]

