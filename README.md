#### echo-twitter-clone
A simple twitter clone using echo.

## Introduction
#### 
First, set values in .env.sample

Next, rename .env.sample to .env
```
cd ~/echo-twitter-clone 
```

```
 mv .env.sample .env
```

#### Create a mysql container
```
docker-compose up -d
```

#### Start server
```
go run server.go
```
Browse to http://localhost:1323 and you should see 'Hello, this is a Twitter clone!' on the page.

## Sample Request
##### Create a message
```
curl -X POST -d "message=sample message" http://localhost:1323/api/posts
```

##### Read all message
```
curl -X GET http://localhost:1323/api/posts
```

##### Read a message
```
curl -X GET http://localhost:1323/api/posts/1
```

##### Update a message
```
curl -X PUT -d "message=update message" http://localhost:1323/api/posts/1
```

##### Delete a message
```
curl -X DELETE http://localhost:1323/api/posts/1
```