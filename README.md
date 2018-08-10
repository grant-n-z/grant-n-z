# Grant Authentication and Authorization
WIP

## Environment variable
```bash
export DB_SOURCE="root:root@tcp(127.0.0.1:3306)/auth_server?charset=utf8&parseTime=True"
``` 
## Build
Git clone
```
$ git clone git@github.com:tomoyane/grant-n-z.git
```

Run test
```
$ revel test grant-n-z test
```

Docker build
```
$ docker build -t grant-n-z:latest .
```

Docker container run
```
$ docker run -p 9000:9000 grant-n-z
```

## Authentication and Authorization
JWT.
