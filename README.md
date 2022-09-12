# echo-twitter-clone
A simple twitter clone using echo.

## Usage
- Create a mysql container(only once)
```
docker-compose up -d
```

- Start server
```
$ go run server.go
```
Browse to http://localhost:1323 and you should see 'Hello, this is a Twitter clone!' on the page.

- Post a message
```
curl -d "message=Hi!" http://localhost:1323/posts
```