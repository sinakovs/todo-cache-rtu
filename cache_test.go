package main

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	val := &todo{
		ID:        "1",
		Item:      "Learn Go",
		Completed: false,
	}
	keys := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	cache := NewTodoCache()

	for i := range 10 {
		go func(i int, val *todo) {
			cache.Set(fmt.Sprint(keys[i]), val)
		}(i, val)
	}
}

func TestCacheRWM(t *testing.T) {
	val := &todo{
		ID:        "1",
		Item:      "Learn Go",
		Completed: false,
	}
	keys := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	cache := NewTodoCacheRWM()

	for i := range 10 {
		go func(i int, val *todo) {
			cache.Set(fmt.Sprint(keys[i]), val)
		}(i, val)
	}
}

func TestShardMap(t *testing.T) {
	val := todo{
		ID:        "1",
		Item:      "Learn Go",
		Completed: false,
	}
	keys := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	cache := NewShardMap(10)

	for i := range 10 {
		go func(i int, val todo) {
			cache.Set(fmt.Sprint(keys[i]), &val)
		}(i, val)
	}
}
