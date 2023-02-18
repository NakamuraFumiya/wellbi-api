package controller

import "echo-twitter-clone/presentation/api/controller/v1/post"

type AppController interface {
	post.PostController
}
