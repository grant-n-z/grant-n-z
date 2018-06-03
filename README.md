# Revel application

## Build
Git clone
```
$ git clone git@github.com:tomo0111/revel-api.git
```

Run test
```
$ revel test revel-api test
```

Docker build
```
$ docker build -t auth-server:latest .
```

Docker container run
```
$ docker run -p 80:9000 auth-server
```

## Authentication and Authorization
JWT.
