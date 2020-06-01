# golang-first-pet
My Golang First Pet project

## Folder structure
```
.
├── app
│   ├── controllers
│   │   ├── chat.go
│   │   ├── index.go
│   │   └── qrcode.go
│   ├── models
│   │   ├── client.go
│   │   ├── message.go
│   │   └── room.go
│   ├── route
│   │   └── route.go
│   ├── shared
│   └── views
│       ├── layouts
│       │   ├── footer.gohtml
│       │   └── header.gohtml
│       ├── pages
│       │   ├── chat.gohtml
│       │   └── qrcode.gohtml
│       └── view.go
├── config
│   └── config.go
├── config.yml.example
├── deployment
│   └── local-mac
│       └── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
├── static
│   ├── js
│   │   ├── chat.js
│   │   └── jquery.min.js
│   └── style
│       └── chatbox.css
├── tree.txt
└── utils
    ├── http
    │   └── http.go
    ├── logger
    │   └── logger.go
    ├── math
    │   └── math.go
    └── network
        └── network.go

```

## Installation
1. Create config file
```
cp ./config.yml.example ./config.yml
```
2. Edit values of some variables in **config.yml**
```
host: <server_ip>
port: <server_port>
lifetime: <lifetime_of_an_message>
```

## Start application on host machine (not containerized)
```
go run main.go
```
