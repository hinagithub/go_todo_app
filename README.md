# go_todo_app
TODO Web Application with AUTH by Go.

# test
```
go test -v ./...
```

# run
``` 
go run . 18000
``` 

```
curl localhost:18000/health

> {"status": "ok"}
```

# Docker
```
docker compose build --no-cache
make up
make down
```