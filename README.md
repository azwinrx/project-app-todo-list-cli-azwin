# TODO List CLI - Golang

Aplikasi To-do List sederhana berbasis Command Line Interface (CLI) menggunakan Golang dan Cobra.

## Video

Link Youtube Video : https://youtu.be/ruemF5zIS5E

## Struktur Project

```
project-app-todo-list-cli-azwin/
├── cmd/                   # Command handlers (Cobra)
│   ├── add.go             # Command untuk menambah task
│   ├── list.go            # Command untuk list task
│   ├── flag.go            # Command untuk update, delete, complete
│   └── root.go            # Root command
├── data/                  # Folder untuk menyimpan data
│   └── data.json          # File JSON untuk storage
├── dto/                   # Data Transfer Objects
│   ├── request.go
│   └── response.go
├── model/                 # Data models
│   └── todo_model.go      # Struct Task
├── service/               # Business logic
│   └── todo_service.go    # Service CRUD operations
├── utils/                 # Utilities
│   └── file.go            # File operations (read/write JSON)
├── main.go                # Entry point aplikasi
└── go.mod                 # Go module dependencies
```
