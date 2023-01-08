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

# chapter17
```
curl localhost:18000/health

curl -i -XGET localhost:18000/tasks

curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden

curl -i -XPOST localhost:18000/tasks -d @./handler/testdata/add_task/bad_req.json.golden

curl -i -XGET localhost:18000/tasks
```