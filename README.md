# golang-first-pet
My Golang First Pet project

## Folder structure
```
.
├── Dockerfile
├── LICENSE
├── README.md
├── app
│   ├── controllers
│   │   ├── chat.go
│   │   ├── index.go
│   │   └── qrcode.go
│   ├── models
│   │   ├── client.go
│   │   ├── message.go
│   │   ├── page.go
│   │   └── room.go
│   ├── route
│   │   └── route.go
│   ├── shared
│   └── views
│       ├── layouts
│       │   ├── footer.gohtml
│       │   └── header.gohtml
│       ├── pages
│       │   ├── chat.gohtml
│       │   └── qrcode.gohtml
│       └── view.go
├── cmd
├── config
│   └── config.go
├── config.yml
├── config.yml.example
├── deployment
│   └── local-mac
│       └── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── static
│   ├── js
│   │   ├── chat.js
│   │   └── jquery.min.js
│   └── style
│       └── chatbox.css
└── utils
    ├── http
    │   └── http.go
    ├── logger
    │   └── logger.go
    ├── math
    │   └── math.go
    └── network
        └── network.go

```

## Start application on host machine (not containerized)
```
go run main.go
```
