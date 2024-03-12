# Backend

To build and run the server:
```sh
go build .
env GD_DSN="postgresql://postgres:123456@localhost:5432/postgres" ./backend
```

To test:
```sh
go test -v ./...
```

To start PostgreSQL server with docker or podman:
```sh
docker run -name postgres-test -d -e POSTGRES_PASSWORD=123456 -p 5432:5432 docker.io/library/postgres
```
