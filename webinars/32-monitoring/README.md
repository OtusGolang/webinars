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
  ### !проверить запись!
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Метрики и мониторинг

### Дмитрий Смаль, Елена Граховац

---

# План занятия

.big-list[
* Observability и Operability
* Метрики и мониторинг
* Мониторинг ресурсов: LA, CPU, MEM, IO
* Мониторинг приложений
* Мониторинг баз данных
* Prometheus и DataDog
]

---

Пока идет теоретическая часть урока, можно скачать себе образы контейнеров ;)
<br><br>

## Prometheus
```
docker pull prom/prometheus
docker pull prom/node-exporter
docker pull grafana/grafana
```

Документация: https://prometheus.io/docs/prometheus/latest/installation/
<br><br>

## DataDog

```
docker pull datadog/agent
```

Попробовать DataDog: elena@grahovac.me

---

# Observability

Observability (наблюдаемость) - мера того, насколько по выходным данным можно
восстановить информацию о состоянии системы.
<br><br>
Примеры:
- логирование (zap, logrus -> fluentd -> kibana)
- мониторинг (zabbix, prometheus,
- алертинг (chatops, pagerduty, opsgenie)
- трейсинг (jaeger, zipkin, opentracing, datadog apm)
- профилирование (pprof)
- сбор ошибок и аварий (sentry)

---

# Operability

Operability (работоспособность) - мера того, насколько приложение умеет
сообщать о своем состоянии здоровья, а инфраструктура управлять этим состоянием.
<br><br>
Примеры:
- простейшие хелсчеки
- liveness и readiness в Kubernetes

---

# Зачем нужен мониторинг ?

.big-list[
* Отладка, решение текущих проблем
* SRE мир: SLA, SLO, SLI
* Алертинг
* Технические и бизнесовые A/B эксперименты
* Анализ трендов, прогнозирование
]

---

# Виды мониторинга

.big-list[
* Количественный / Событийный
* Whitebox / Blackbox
* Push / Pull
]


---

# Push vs Pull

*Push* - агент, работающий на окружении (например, сайдкар), подключается к серверу мониторинга
и отправляет данные.

Особенности:
* мониторинг специфических/одноразовых задач
* можно поставить хоть куда :)
* не нужно открывать никакие URL'ы/порты на стороне пирложения
* из приложения нужно конфигурировать подключение

Примеры: `graphite`, `statsd`
<br><br>

*Pull* - сервис мониторинга сам опрашивает инфраструктуры/сервисы и агрегирует статистику.

Фичи:
* не нужно подключаться к агенту на стороне приложения 
* нужно открывать URL или порт, который будет доступен сервису мониторинга

Примеры: `datadog-agent`, `prometheus`

---

# Мониторинг инфраструктуры

.big-list[
* LA (Load Average) - длинна очереди процессов в планировщике
* CPU (User/System/Wait) - время проводимое процессором в различных режимах
* Memory (RSS/Shared/Cached/Free) - распределение памяти в системе
* IO stats
* Network stats
]

---

# Load Average

LA - сложная метрика, ее можно интерпретировать как количество процессов(потоков) в ОС, 
находящихся в ожидании какого-либо ресурса (чаща всего CPU или диск). 
<br><br>
*Нормальной* считается загрузка когда LA ~ числу ядер процессора.
<br><br>

Как посмотреть:
* top
* iostat, dstat
  
---

# CPU

* User (`usr`, `us`) - процессор выполняет код программ. Высокое значение может быть признаком неоптимальных алгоритмов.
* System (`sys`, `sy`) - процессор выполняет код ядра. Высокое значение может означать большое кол-во операций ввода/вывода или сетевых пакетов.
* Wait (`wai`, `wa`) - процессор находится в ожидании ввода/вывода. Высокое значение означает недостаточную мощность дисковой системы.
* Idle (`id`)- процессор бездействует.

<br><br>
Как посмотреть:
* top, htop

---

# Memory

* Resident (`Res`) - память, занятая данными программ (как правило кучи). Высокое значение может говорить об утечках памяти в программах.
* Shared (`Shr`) - память, разделяемая между разными процессами (как правило сегменты кода).
* Cached - дисковый кеш операционный системы, в нагруженных системах (СУБД) состоянии занимает все свободное место.
* Free - не занятая память.
  
Как посмотреть:
* top
* free
  
---

# IO

* `%util` - процент времени в течение которого диск занят вводом/выводом.
* `r/s`, `w/s` - число запросов на чтение/запись в секунду.
* `rKb/s`, `wKb/s` - объем данных прочитанных/записанных в секунду.
* `await`, `r_await`, `w_await` - среднее время в мс. ожидания+выполнения запроса ввода/вывода. latency диска.
* `avgqu-sz` - средняя длинна очереди запросов к диску.

Как посмотреть:
* iostat -x 1
* dstat

Проблемы:
* `%util` ~ 100% - вам не хватает мощности диска
* `%util` сильно отличается у разных дисков в RAID - неисправен диск?

---

# Траблшутинг

Алгоритм:
* идентифицировать проблему
* найти причину (?)
* решить  проблему

Расширенный алгоритм: управление инцидентами.
Гуглить по фразам "incident management" и "blameless postmortem".

<br><br>
Инструменты:
* `top`, `htop` - умеют сортировать процессы по CPU, RES
* `iotop` - умеет сортировать процессы по использованию диска
* `iftop` - трафик, по хостам 
* `atop` - записывает лог, позволяет исследовать ситуацию *post hoc*
  
---

# Мониторинг приложений

.big-list[
* RPS (request per second)
* Response time 
* Задержка между компонентами приложения (latency)
* Код ответа (HTTP status 200/500/5xx/4xx)
* Разделение по методам API
]
<br><br>
Для детального анализа: трейсинг

---

# Распределение значений

.main_image[
![img/latency.png](img/latency.png)
]

Среднее значение (`avg`, `mean`) или даже медиана (`median`) не отображают всей картины!
<br><br>
Полезно измерять *процентили* (percentile), например время в которое укладываются 95% или 99% запросов.

---

# Мониторинг баз данных

.big-list[
* TPS (transactions per second)
* QPS (queries per second)
* IO usage
* CPU usage
* Replication Lag
* Wait Events
* Active connections
]

---

# Основные группы метрик

.big-list[
* Latency - время задержки
* Traffic - количество запросов и объем трафика
* Errors - количество и характер ошибок
* Saturation - утилизация ресурсов
]

---


# Prometheus

.main-image[
![img/prometheus.png](img/prometheus.png)
]

---

# Prometheus - Infrastructure as Code

Дисклеймер: далее на слайдах будут примеры установки и настройки
на Linux-машинах. В реальной жизни вряд ли вы будете настраивать это руками :)
<br><br>
Куда смотреть для production-readiness: ansible, chef и всякое такое 
https://prometheus.io/docs/prometheus/latest/installation/

---

# Prometheus - установка сервера

Установка на Ubuntu:
```
$ apt get install prometheus
```

Настройка `/etc/prometheus/prometheus.yml`
```
global:
  scrape_interval:  15s  # как часто опрашивать exporter-ы

storage:
  tsdb:
    path: /var/lib/prometheus # где хранить данные
    retention:
      time:  15d              # как долго хранить данные

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']
  - job_name: 'app'
    static_configs:
      - targets: ['localhost:9100', 'localhost:9102', 'localhost:9103', 'localhost:9187'] 
```
---

# Prometheus - запуск

Запуск
```
$ service prometheus start
```

С настройками по умолчанию Prometheus будет доступен на порту 9090: 
[http://127.0.0.1:9090/](http://127.0.0.1:9090/)

---

# Prometheus - мониторинг сервера

.main_image[
  ![img/prometheus_linux.png](img/prometheus_linux.png)
]

---

# Prometheus - мониторинг сервера

Установка и запуск collectd
```
$ apt-get install collectd

# В /etc/collectd/collectd.conf
LoadPlugin network
<Plugin network>
  Server "127.0.0.1" "25826"
</Plugin>

# запускаем collected
$ service start collectd

# скачиваем collected-exporter
$ wget https://github.com/prometheus/collectd_exporter/releases/download/v0.4.0/collectd_exporter-0.4.0.linux-amd64.tar.gz
$ tar -zxf collectd_exporter-0.4.0.linux-amd64.tar.gz

# запускаем exporter
$ ./collectd_exporter --collectd.listen-address="localhost:25826" \
                      --web.listen-address="localhost:9103"    
```
---

# Prometheus - мониторинг Postgres

Установка postgres-exporter
```
$ go get github.com/wrouesnel/postgres_exporter
$ cd ${GOPATH-$HOME/go}/src/github.com/wrouesnel/postgres_exporter
$ go run mage.go binary
```

Запустить с указанием connection string
```
$ export DATA_SOURCE_NAME="postgresql://login:password@hostname:port/dbname"
$ ./postgres_exporter
```
---

# Prometheus - протокол

Простой способ исследовать: `wget -O - http://localhost:9103/metrics`

```
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 1.036096e+06

collectd_processes_ps_state{instance="mialinx-test-ub.ru-central1.internal",processes="blocked"} 0
collectd_processes_ps_state{instance="mialinx-test-ub.ru-central1.internal",processes="paging"} 0
collectd_processes_ps_state{instance="mialinx-test-ub.ru-central1.internal",processes="running"} 1
collectd_processes_ps_state{instance="mialinx-test-ub.ru-central1.internal",processes="sleeping"} 57
collectd_processes_ps_state{instance="mialinx-test-ub.ru-central1.internal",processes="stopped"} 0
collectd_processes_ps_state{instance="mialinx-test-ub.ru-central1.internal",processes="zombies"} 0

go_gc_duration_seconds{quantile="0"} 4.0147e-05
go_gc_duration_seconds{quantile="0.25"} 6.950600000000001e-05
go_gc_duration_seconds{quantile="0.5"} 0.000108126
go_gc_duration_seconds{quantile="0.75"} 0.001107202
go_gc_duration_seconds{quantile="1"} 0.039212351
go_gc_duration_seconds_sum 0.49406203400000004
go_gc_duration_seconds_count 282
```
---

# Prometheus - типы метрик

.big-list[
* `Counter` - монотонно возрастающее число, например число запросов
* `Gauge` - текущее значение, например потребление памяти
* `Histogram` - распределение значений по бакетам (сколько раз значение попало в интервал)
* `Summary` - похоже на `histogram`, но по квантилям
* Векторные типы для подсчета данных по меткам
]

Документация: https://prometheus.io/docs/concepts/metric_types/

Отличная документация в godoc: https://godoc.org/github.com/prometheus/client_golang/prometheus

---

# Prometheus - мониторинг Go HTTP сервисов

```
import (
  "log"
  "net/http"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  metrics "github.com/slok/go-http-metrics/metrics/prometheus"
  "github.com/slok/go-http-metrics/middleware"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("hello world!"))
}

func main() {
  // middleware для мониторинг
  mdlw := middleware.New(middleware.Config{
    Recorder: metrics.NewRecorder(metrics.Config{}),
  })
  h := mdlw.Handler("", http.HandlerFunc(myHandler))
  // HTTP exporter для prometheus
  go http.ListenAndServe(":9102", promhttp.Handler())
  // Ваш основной HTTP сервис
  if err := http.ListenAndServe(":8080", h); err != nil {
    log.Panicf("error while serving: %s", err)
  }
}
```

---

# Prometheus - собственные метрики

```
import "github.com/prometheus/client_golang/prometheus"

var regCounter = prometheus.NewCounter(prometheus.CounterOpts{
  Name: "business_registration",
  Help: "Client registration event",
})

func init() {
  prometheus.MustRegister(regCounter)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Hello, world!"))
  regCounter.Inc()
}

```
---

# DataDog - типы метрик

.big-list[
* `Count` (*) - в отличие от Prometheus может и возрастать, и убывать
* `Gauge` (*) - текущее значение, например потребление памяти
* `Histogram` - считает агрегированные avg, count, median, 95percentile, max
* `Distribution`  - больше агрегаций, чем histogram
* `Set` -  счетчик уникальных значений
* `Rate` (*) - скорость изменения значения
]

(*) - основные типы метрик, на их основе хранятся остальные

Документация: https://docs.datadoghq.com/developers/metrics/types/

---

# Домашнее задание

Обеспечить простейший мониторинг проекта с помощью prometheus

Prometheus запустить в docker контейнере рядом с остальными сервисами.

Для API сервиса необходимо измерять:
* Requests per second
* Latency
* Коды ошибок
* Все это в разделении по методам (использовать отдельный тэг prometheus для каждого метода API)

Для баз данных:
* Количество записей в таблице events (данные брать из pg_stat_user_tables)
* Стандартные метрики базы: Transactions per second, количество подключений (использовать готовый exporter)

Для расслыльщика:
* RPS (кол-во отправленных сообщений в сек)

---

# Опрос

.left-text[
Заполните пожалуйста опрос
<br><br>
[https://otus.ru/polls/6330/](https://otus.ru/polls/6330/)
]

.right-image[
![](img/gopher7.png)
]

---

class: white
background-image: url(img/message.svg)
.top.icon[![otus main](img/logo.png)]

# Спасибо за внимание!

---

# Дополнительные материалы

SRE book:<br>
  [https://landing.google.com/sre/sre-book/chapters/monitoring-distributed-systems/](https://landing.google.com/sre/sre-book/chapters/monitoring-distributed-systems/)
<br><br>
Monitoring and Observability:<br>
  [https://medium.com/@copyconstruct/monitoring-and-observability-8417d1952e1c](https://medium.com/@copyconstruct/monitoring-and-observability-8417d1952e1c)
  