# echo-twitter-clone
A simple twitter clone using echo.

## Introduction
### Create a mysql container
```
$ docker-compose up -d
```

### Start server
```
$ go run server.go
```
Browse to http://localhost:1323 and you should see 'Hello, this is a Twitter clone!' on the page.

## Usage
### Post a message
```
$ curl -d "message=sample message" http://localhost:1323/posts
```

### Read all message
```
$ curl http://localhost:1323/posts
```