# pingmen
Ping bot for gitlab merge requests

[На русском](README_RUS.md)

1. [Running without build](#run)
2. [Build](#build)
3. [Flags](#flags)
4. [Configuration file](#cfgfile)
5. [Configuration telegram](#cfgtelegram)
6. [Configuration gitlab](#cfggitlab)
7. [TODO](#todo)

## Running without build <a name="run"></a>
```zsh
go mod tidy
go run main.go -c <your_config_file_name>.yaml
```

## Build <a name="build"></a>
```zsh
go mod tidy
env GOOS=<your_OS> go build 
```

### Launch after build
```zsh
chmod +x pingmen
./pingmen -c <your_config_file_name>.yaml
```

## Flags <a name="flags"></a>
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

## Configuration file <a name="cfgfile"></a>
```yaml
gitlab: # gitlab settings
  token: "token" # token for communication with gitlab
  webhook_method: "webhook" # method for hooks from gitlab
  webhook_port: 80 # the port on which the hooks are expected
  actions: # events sent by gitlab when merge request
    - "open" # a new merge request is open
    - "reopen" # merge request reopened
    - "update" # merge request updated
    - "close" # merge request closed
telegram: # bot telegram settings
  token: "token" # bot telegram token
  chat_id: -1 # chat id, for groups they go with -
  workers_count: 2 # number of message-sending workers
  debug: true # bot debug mode, it is recommended to turn it off, many logs
users: # usernames in the telegram, which need to be pinged in the chat, are written without @
  - "user1" # will display @ user1 in the chat
  - "user2"
projects: # the name of the project to display messages about
  - "project1"
  - "project2"
```

## Configuration telegram <a name="cfgtelegram"></a>
1. Chat @BotFather command ```/newbot```
2. Enter bot name
3. Enter another username for the bot, ```_bot``` (copy token)
4. Enter the @BotFather command into the chat ```/setprivacy```
5. Select the bot you want
6. Choose ```Disabled```
7. Add a bot to the desired group
8. You can find out the group chat by entering a request in the browser line
```html
https://api.telegram.org/bot<bot_token>/getUpdates
```
9. How will some updates arrive in the chat, update the page and in ```resul->chat->``` id chat

## Configuration gitlab <a name="cfggitlab"></a>
Project -> ```Settings``` -> ```Webhooks```
1. Enter in field URL ```http://<host>:<webhook_port>/<webhook_method>```
2. Check box on ```Merge request events```
3. Press button ```Add webhook```
4. Created webhook can be tested -> button ```Test```

# TODO: <a name="todo"></a>
1. Log levels
2. More actions
3. Message templates