# pingmen
Бот уведомлялка о мердж событиях гитлаба


1. [Запуск без сборки](#run)
2. [Сборка](#build)
3. [Флаги](#flags)
4. [Файл конфигурации](#cfgfile)
5. [Настройка telegram](#cfgtelegram)
6. [Настройка gitlab](#cfggitlab)
7. [TODO](#todo)

## Запуск без сборки <a name="run"></a>
```zsh
go mod tidy
go run main.go -c <your_config_file_name>.yaml
```

## Сборка <a name="build"></a>
```zsh
go mod tidy
env GOOS=<your_OS> go build 
```

### Запуск после сборки
```zsh
chmod +x pingmen
./pingmen -c <your_config_file_name>.yaml
```

## Флаги <a name="flags"></a>
```zsh
-c string
-c <path to сonfig file>
-cfg string
--config <path to сonfig file>
-h    help flag usage
-help
help flag usage
-l string
-l <path to log file>
-log string
--log <path to log file>
```

## Файл конфигурации <a name="cfgfile"></a>
```yaml
loglevel: "debug" # уровенть логирования, по умолчанию warning
gitlab: # настройки гитлаба
  token: "token" # токен для связи с гитлабом
  webhook_method: "webhook" # метод для хуков от гитлаба
  webhook_port: 80 # порт, на который ожидаем хуки
  actions: # события, которые присылает гитлаб при merge request
    - "open" # открыт новый merge request
    - "reopen" # merge request переоткрыт
    - "update" # merge request обновился
    - "close" # merge request закрыт
telegram: # настройки телеграмм бота
  token: "token" # токен телеграмм бота
  chat_id: -1 # идентификатор чата, у групп они идут с -
  workers_count: 2 # количество воркеров-отправщиков сообщений
  debug: true # режим отладки бота, рекомендовано выключить, много логов
users: # имена пользователей в телеграме, которых надо пингануть в чате, пишутся без @
  - "user1" # в чат выведет @user1
  - "user2"
projects: # имя проекта, сообщения о котором надо выводить
  - "project1"
  - "project2"
```

## Настройка telegram <a name="cfgtelegram"></a>
1. Написать в чат @BotFather команду ```/newbot```
2. Ввести имя бота
3. Ввести еще одно имя пользователя для бота, ```_bot``` (копирнуть токен)
4. Ввести в чат @BotFather команду ```/setprivacy```
5. Выбрать нужного бота
6. Выбрать ```Disabled```
7. Добавить бота в нужную группу
8. Узнать чат группы можно введя в строке браузера запрос
```html
https://api.telegram.org/bot<bot_token>/getUpdates
```
9. Как в чат прилетят какие то обновы, обновить станичку и в ```resul->chat->``` идентификатор чата

## Настройка gitlab <a name="cfggitlab"></a>
Проект -> ```Settings``` -> ```Webhooks```
1. Ввести в поле URL ```http://<host>:<webhook_port>/<webhook_method>```
2. Отметить галочку на ```Merge request events```
3. Ввести в поле Token ключ из конфиг файла 
4. Нажать кнопку ```Add webhook```
5. Созданный webhook можно протестировать -> кнопка ```Test```

# TODO: <a name="todo"></a>
1. Больше действий
2. Шаблон сообщений