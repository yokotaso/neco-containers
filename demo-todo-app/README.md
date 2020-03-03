# Getting Started

### Build executable jar

```sh
$ ./gradlew bootJar
```

### Run

```sh
$ ./gradlew bootRun
```

### API

Get TODO
```sh
curl -X GET localhost:8080/todo/1
```

Save TODO
```sh
$ curl -X POST -H 'Content-Type:application/json' -d '{"todoTitle":"hoge","finished":false,"createdAt":"2020-03-03T09:00:00"}' http://localhost:8080/todo
```

Delete TODO
```sh
curl -X DELETE localhost:8080/todo/1
```

### Environment Variable

- ENV_DB_HOST:localhost
- ENV_DB_PORT:3306
- ENV_DB_NAME:todo
- ENV_DB_USER:root
- ENV_DB_PASSWORD:mysql
