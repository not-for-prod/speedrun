# not-for-prod/speedrun

> Speedrun is a bunch of dirty hacks that removes most boring parts of development process by generating code that is not perfect but works if you're lucky.

Installation:

```bash
go install github.com/not-for-prod/speedrun
```

# Main features:

## CRUD

> generates postgres CRUD operations by entity

Usage:

```bash
speedrun crud --src internal/crud/example/in/peach.go::Peach::Id --dst internal/crud/example/out
```

- Input: ex. [internal/crud/example/in/peach.go](internal/crud/example/in/peach.go) - simple typed entity

- Output: ex. [internal/crud/example/out](internal/crud/example/out)
    In dst folder:
    ```
    └── <entity>
        ├── sql
        │   ├── create.sql
        │   ├── get.sql
        │   ├── update.sql
        │   ├── delete.sql
        │   └── sql.go
        ├── create.go
        ├── get.go
        ├── update.go
        ├── delete.go
        └── repository.go
    ```
    In root folder
    ```
    └── migrations
        ├── <uuid>_create_<entity>.up.sql
        └── <uuid>_create_<entity>.down.sql
    ```

## Layout

> generates project folders

Usage:

```bash
speedrun layout --svc aboba
```

Output: ex. [internal/layout/example](internal/layout/example)
```
├── cmd
│   └── <service_name>
├── internal
│   ├── config
│   ├── <service_name> 
│   │   ├── app
│   │   ├── domain
│   │   │   ├── entity
│   │   │   ├── valueobject
│   │   │   └── microtypes
│   │   └── infrastructure
│   └── pkg
└── pkg
```

## GRPC-Gateway

> generates grpc-gateway boilerplate

- buf
- swagger
- http and grpc handlers
- middlewares that are working both for http and grpc

## Additional

### Errore
### Burito