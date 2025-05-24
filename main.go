package main

import (
	"net/http"
	"time"
)

var Todo, _ = GetTodoFromFile()
var Cache = NewTodoCache()
var CacheRWMutex = NewAndFillCacheRWM()
var ShardMapCache = NewAndFillShardMap()

func main() {

	port := "8080"

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        routes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	server.ListenAndServe()
}
