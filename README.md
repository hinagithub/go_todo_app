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
curl localhost:18080/test

> Hello, test!
```

# Docker
```
docker compose build --no-cache
make up
make down
```