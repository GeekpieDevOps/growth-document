# Backend

```sh
go build github.com/GeekpieDevOps/growth-document/backend
env GD_DSN="host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai" ./backend
```

To temporarily start postgresql server:
```sh
docker run -id --name=postgres-test -v postgre-data:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=123456 -e LANG=C.UTF-8 postgres
```
