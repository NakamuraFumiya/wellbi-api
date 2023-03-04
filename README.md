#### wellbi-api

## Introduction
#### 
First, set values in .env.sample

Next, rename .env.sample to .env
```
cd ~/wellbi-api 
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
Browse to http://localhost:1323 and you should see 'Hello, this is a wellbi-api!' on the page.

## Request
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