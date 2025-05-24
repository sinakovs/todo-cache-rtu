package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Todo); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func getTodoCache(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Cache); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func getTodoCacheRWM(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(CacheRWMutex.data); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func getTodoShardMap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(showAllShards(*ShardMapCache)); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var temp todo

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	temp.ID = strconv.Itoa(len(Todo) + 1)
	Todo = append(Todo, temp)

	w.WriteHeader(http.StatusCreated)
}

func addTodoCache(w http.ResponseWriter, r *http.Request) {
	var temp todo

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	temp.ID = strconv.Itoa(len(Cache) + 1)
	Cache.Set(temp.ID, &temp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(temp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func addTodoCacheRWM(w http.ResponseWriter, r *http.Request) {
	var temp todo

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	temp.ID = strconv.Itoa(len(CacheRWMutex.data) + 1)
	CacheRWMutex.Set(temp.ID, &temp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(temp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func addTodoShardMap(w http.ResponseWriter, r *http.Request) {
	var temp todo

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	temp.ID = strconv.Itoa(getShardsLen(*ShardMapCache) + 1)
	ShardMapCache.Set(temp.ID, &temp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(temp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func getTodoById(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/todo/")
	if path == "" || strings.Contains(path, "/") {
		http.Error(w, "INvalid ID format", http.StatusBadRequest)
		return
	}

	temp, err := findTodoById(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(temp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func getTodoCacheById(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/cache/")
	if path == "" || strings.Contains(path, "/") {
		http.Error(w, "INvalid ID format", http.StatusBadRequest)
		return
	}

	temp, ok := Cache.Get(path)
	if !ok {
		err := fmt.Errorf("there is no element with id: %s", path)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(temp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func getTodoCacheRWMById(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/rwm/")
	if path == "" || strings.Contains(path, "/") {
		http.Error(w, "INvalid ID format", http.StatusBadRequest)
		return
	}

	temp, ok := CacheRWMutex.Get(path)
	if !ok {
		err := fmt.Errorf("there is no element with id: %s", path)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(temp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func getTodoShardMapById(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/shard/")
	if path == "" || strings.Contains(path, "/") {
		http.Error(w, "INvalid ID format", http.StatusBadRequest)
		return
	}

	temp, ok := ShardMapCache.Get(path)
	if !ok {
		err := fmt.Errorf("there is no element with id: %s", path)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(temp); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
