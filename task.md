Тестовое задание
Очередь на go с REST интерфейсом
Реализовать брокер очередей в виде веб сервиса. Сервис должен обрабатывать 2 метода:

PUT /queue/:queue

{
    "message": "data"
}
Положить сообщение message в очередь с именем :queue (имя очереди может быть любое), пример:

curl -XPUT http://localhost/queue/pet
curl -XPUT http://localhost/queue/role

в ответ {пустое тело + статус 200 (ok)} в случае если тело не по формату или отсутствует отдавать статус 400 (bad request)

Желательно предусмотреть настройку максимального количества сообщений в очереди и общее количество очередей.

GET /queue/:queue

Забрать (по принципу FIFO) из очереди с названием :queue сообщение и вернуть в теле http запроса, пример (результат, который должен быть при выполненных put’ах выше)

при GET-запросах сделать возможность задавать аргумент timeout

curl http://localhost/queue/pet?timeout=N

если в очереди нет готового сообщения получатель должен ждать либо до момента прихода сообщения либо до истечения таймаута (N - кол-во секунд). В случае, если сообщение так и не появилось - возвращать код 404. Получатели должны получать сообщения в том же порядке как от них поступал запрос, если 2 получателя ждут сообщения (используют таймаут), то первое сообщение должен получить тот, кто первый запросил.

Порт, на котором будет слушать сервис, должен задаваться в аргументах командной строки. Настройка таймаута по умолчанию, должена задаваться или отключаться в аргументах командной строки как и остальные параметры.

Запрещается пользоваться какими либо сторонними пакетами кроме стандартных библиотек. (задача в написании кода, а не в использовании чужого)

Желательно организовать код используя принципы гексоганальной архитектуры. Комментарии приветствуются и помогут нам понять ход Ваших мыслей при разработке.

Лаконичность кода будет восприниматься крайне положительно, не нужна "гибкость" больше, чем требуется для решения именно этой задачи, не нужны логи процесса работы программы (только обработка ошибок), никакого дебага и т.д... чем меньше кода - тем лучше!

Оцениваться корректность реализации (заданные условия выполняются),архитектурная составляющая (нет лишних действий в программе, только решающие задачи программы), лаконичность и понятность кода (субъективно, конечно, но думайте о том, насколько будет понятен ваш. код для других, это куда более важно в командной разработке, чем сложный "крутой" код).
