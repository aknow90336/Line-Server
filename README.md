# Line Official Account

`dependency`

原則上專案編譯時會自行安裝相關套件，但也可以先執行下列指令，安裝 module 套件。
```
go mod tidy
```

`directory tree`

    ├── api                   # api handler
    ├── config                # 配置檔
    ├── domain                # 資料結構
    ├── env                   # 初始化相關資源 (
    ├── repo                  # repository 層級介面與實作
    ├── .gitignore                  
    ├── docker-compose.yml           
    └── main.go               # 程式進入點