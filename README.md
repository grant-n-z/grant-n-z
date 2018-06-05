# Authentication server

## Build
Git clone
```
$ git clone git@github.com:tomo0111/authentication-server.git
```

Run test
```
$ revel test auth-server test
```

Docker build
```
$ docker build -t auth-server:latest .
```

Docker container run
```
$ docker run -p 9000:9000 auth-server
```

## Authentication and Authorization
JWT.
