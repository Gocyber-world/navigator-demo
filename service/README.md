
### Config go project

```
go generate -x
```

### Run MySQL locally

```
docker run --name local-test-mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=localpassword -d mysql:5.7
```

### Init MySQL tables
```
go run main.go initMigrateMysql --config configs/config.local.yaml
```

### Generate API docs
```
swag init
```

### Start service
```
go run main.go serve --config configs/config.local.yaml
```

### Visit API docs
```
http://localhost:9000/swagger/index.html
```
