# not-for-prod/speedrun

> Speedrun is a bunch of dirty hacks that removes most boring parts of development process by generating code that is not perfect but works if you're lucky.

Installation:

```bash
go install github.com/not-for-prod/speedrun
```

# Main features:

## CRUD

Execute:

```bash
speedrun crud --src internal/crud/example/in/peach.go::Peach::Id --dst internal/crud/example/out
```

- Input: ex. [internal/crud/example/in/peach.go](internal/crud/example/in/peach.go) - simple typed entity

- Output:
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