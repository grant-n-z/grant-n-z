# Authentication server

## Build
Git clone
```
$ git clone git@github.com:tomo0111/authentication-server.git
```

Run test
```
$ revel test authentication-server test
```

Docker build
```
$ docker build -t authentication-server:latest .
```

Docker container run
```
$ docker run -p 9000:9000 authentication-server
```

## Authentication and Authorization
JWT.
