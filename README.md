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
### Create a message
```
$ curl -X POST -d "message=sample message" http://localhost:1323/posts
```

### Read all message
```
$ curl -X GET http://localhost:1323/posts
```

### Read a message
```
$ curl -X GET http://localhost:1323/posts/1
```

### Update a message
```
$ curl -X PUT -d "message=update message" http://localhost:1323/posts/1
```

### Delete a message
```
$ curl -X DELETE http://localhost:1323/posts/1
```