background-image: url(../img/title.svg)

---

background-image: url(../img/rec.svg)

---
background-image: url(../img/topic.svg)

.topic[Go внутри. Планировщик]
.tutor[Алексей Семушкин]
.tutor_desc[Software engineer at Semrush]

---

background-image: url(../img/rules.svg)

---

# О чем будем говорить

* ### Что такое планировщик.
* ### Особенности работы планировщика Go.
* ### Почему планировщик Go такой, какой он есть.

---

# Зачем нужен планировщик?

---

# Зачем нужен планировщик?

### Распределять вычислительные ресурсы между задачами.

---

# Зачем Go собственный планировщик?

* ### Уже есть планировщик процессов и потоков в ядре.

* ### Почему бы не запускать каждую горутину в отдельном потоке?

---

# Зачем Go собственный планировщик?

## Проблемы потоков

* ### Потоки дорогие по памяти (из-за стэка).
* ### Потоки дорого переключать (сисколы, контекст).

<br>

## Решения в Go

* ### Go использует growable stack.
* ### Go выбирает моменты, когда переключение дешевое.

---

# Проектируем планировщик: m-n threading

.full-image[
![](img/mn_1.png)
]

---

# Проектируем планировщик: m-n threading

.full-image[
![](img/mn_2.png)
]

Какие есть проблемы?

---

# Проектируем планировщик: m-n threading

.full-image[
![](img/mn_3.png)
]

.

---

# Проектируем планировщик: отдельные очереди

.full-image[
![](img/run_q.png)
]

Какие есть проблемы?

---

# Проектируем планировщик: закончилась очередь

.full-image[
![](img/run_q_ws_1.png)
]

---

# Проектируем планировщик: work stealing

.full-image[
![](img/run_q_ws_2.png)
]

http://supertech.csail.mit.edu/papers/steal.pdf

---

# Проектируем планировщик: syscall

.full-image[
![](img/syscall_1.png)
]

Тред заблокирован.

---

# Проектируем планировщик: syscall

.full-image[
![](img/syscall_2.png)
]

Создаем новый тред.

---

# Проектируем планировщик: syscall

.full-image[
![](img/syscall_3.png)
]

Куда деть горутину после syscall'a?


---

# Проектируем планировщик: глобальная очередь

.full-image[
![](img/pool_1.png)
]

---

# Проектируем планировщик: все идеи вместе

.full-image[
![](img/pool_2.png)
]

---

# Планировщик: go tool trace

* https://golang.org/cmd/trace/
* https://making.pusher.com/go-tool-trace/
* https://blog.gopheracademy.com/advent-2017/go-execution-tracer/

---

# Планировщик: честность

.full-image[
![](img/fifo_lifo_1.png)
]

Какой тип очереди честнее?

---

# Планировщик: честность

.full-image[
![](img/fifo_lifo_2.png)
]

Какие проблемы?

---

# Планировщик: честность, trade-off

## За честность приходится платить производительностью.

---

# Планировщик: честность, trade-off

.full-image[
![](img/run_q_fifo_lifo.png)
]

Одноэлементное LIFO улучшает использование кэша -> дешевле переключать горутины.

---

# Планировщик: голодание FIFO

* ### А что если две горутины постоянно ставят друг друга в LIFO?

.

---

# Планировщик: голодание FIFO

* ### А что если две горутины постоянно ставят друг друга в LIFO?

Считать время непрерывной работы цепочки горутин.

---

# Планировщик: голодание горутины

* ### А что если одна горутина находится в бесконечном цикле?

---

# Планировщик: голодание горутины

Go 1.14 Asynchronous Preemption:
* https://medium.com/a-journey-with-go/go-asynchronous-preemption-b5194227371c
* https://github.com/golang/proposal/blob/master/design/24543-non-cooperative-preemption.md

---

# Планировщик: порядок поиска работы

* ### Локальная очередь
* ### Глобальная очередь
* ### Work stealing

---

# Планировщик: голодание глобальной очереди

* ### А что если горутины в локальных очередях не кончаются?

.full-image[
![](img/glob_starv.png)
]

---

# Планировщик: голодание глобальной очереди

* ### А что если горутины в локальных очередях не кончаются?

Брать горутины из глобальной очереди каждый 61-й тик.

---

# Планировщик: network poller

### Проверка готовности горутины в очереди — это проверка значения в мапе.
* Это быстро и дешево.

<br>

### Проверка готовности сетевого io — это syscall
* Syscall паркует тред.
* Syscall это дорого.

<br>

## Что делать?

---

# Планировщик: network poller

* ## Завести для network poller отдельный тред

---

# Планировщик: честность, саммари

* ### Бесконечные горутины — sysmon
* ### Циклы из горутин — давать работать суммарно 10 мс
* ### Глобальная очередь — периодически брать горутины из нее
* ### Network poller — отдельный тред

---

# Каналы

https://golang.org/src/runtime/chan.go

---

# Каналы

.full-image[
![](img/write_read_make.png)
]

---

# Каналы: запись

.full-image[
![](img/write_read_w0.png)
]

---

# Каналы: запись

.full-image[
![](img/write_read_w1.png)
]

---

# Каналы: чтение

.full-image[
![](img/write_read_r0.png)
]

---

# Каналы: concurrency

.full-image[
![](img/concurrency_w_init.png)
]

---

# Каналы: concurrency

.full-image[
![](img/concurrency_w_g1_lock.png)
]

---

# Каналы: concurrency

.full-image[
![](img/concurrency_w_g1_copy.png)
]

---

# Каналы: concurrency

.full-image[
![](img/concurrency_w_g1_unlock.png)
]

---

# Каналы: concurrency

.full-image[
![](img/concurrency_w_g2_lock.png)
]

---

# Каналы: concurrency

.full-image[
![](img/concurrency_w_g2_copy.png)
]

---

# Каналы: concurrency

.full-image[
![](img/concurrency_w_g2_unlock.png)
]

---

# Каналы: concurrency (write full)

.full-image[
![](img/concurrency_wf_init.png)
]

---

# Каналы: concurrency (write full)

.full-image[
![](img/concurrency_wf_init_2.png)
]

---

# Каналы: concurrency (write full)

.full-image[
![](img/concurrency_wf_g1_lock.png)
]

---

# Каналы: concurrency (write full)

.full-image[
![](img/concurrency_wf_g2_copy.png)
]

---

# Каналы: concurrency (write full)

.full-image[
![](img/concurrency_wf_g2_g1_copy.png)
]

---

# Каналы: concurrency (write full)

.full-image[
![](img/concurrency_wf_g1_ready.png)
]

---

# Каналы: concurrency (read empty)

.full-image[
![](img/concurrency_re_init.png)
]

---

# Каналы: concurrency (read empty)

.full-image[
![](img/concurrency_re_init_2.png)
]

---

# Каналы: concurrency (read empty)

.full-image[
![](img/concurrency_re_g2_lock.png)
]

---

# Каналы: concurrency (read empty)

.full-image[
![](img/concurrency_re_g1_w1.png)
]

---

.full-image[
![](img/no_double_copy.jpg)
]

---

# Каналы: concurrency (read empty)

.full-image[
![](img/concurrency_re_g1_w2.png)
]

---

# Каналы: горутина пишет в стек другой горутины!

https://cs.opensource.google/go/go/+/master:src/runtime/chan.go;l=211

---

# Материалы

* ### Сага о планировщике — https://youtu.be/YHRO5WQGh0k
* ### Планировщик шаг за шагом — https://youtu.be/-K11rY57K7k
* ### Про каналы — https://www.youtube.com/watch?v=KBZlN0izeiY
* ### Про планировщик на русском — https://youtu.be/Gy6XEYWYht8

---

background-image: url(../img/questions.svg)

---

background-image: url(../img/poll.svg)

---

background-image: url(../img/next_webinar.svg)
.announce_date[13 сентября]
.announce_topic[Go внутри. Память и сборка мусора]

---
background-image: url(../img/thanks.svg)

.tutor[Алексей Семушкин]
.tutor_desc[Software engineer at Semrush]
