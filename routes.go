package main

import (
	"net/http"
)

func routes() http.Handler {

	api := http.NewServeMux()

	api.HandleFunc("GET /", getTodo)
	api.HandleFunc("GET /todo/", getTodoById)
	api.HandleFunc("POST /", addTodo)

	api.HandleFunc("GET /cache", getTodoCache)
	api.HandleFunc("GET /cache/", getTodoCacheById)
	api.HandleFunc("POST /cache", addTodoCache)

	api.HandleFunc("GET /rwm", getTodoCacheRWM)
	api.HandleFunc("GET /rwm/", getTodoCacheRWMById)
	api.HandleFunc("POST /rwm", addTodoCacheRWM)

	api.HandleFunc("GET /shard", getTodoShardMap)
	api.HandleFunc("GET /shard/", getTodoShardMapById)
	api.HandleFunc("POST /shard", addTodoShardMap)

	return api
}
