# pingmen
Ping bot for gitlab merge requests

[На русском](README_RUS.md)

## Readme TODO:
1. Configuration
2. Gitlab hook
3. Telegram bot steps

## Run
```zsh
go mod tidy
go run main.go -c <your_config_file_name>.yaml
```

## Build
```zsh
go mod tidy
env GOOS=<your_OS> go build 
```

### Run after build
```zsh
chmod +x pingmen
./pingmen -c <your_config_file_name>.yaml
```

# TODO:
1. Log levels
2. More actions
3. Message templates