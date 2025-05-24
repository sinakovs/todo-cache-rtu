package main

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

var dataPath = "./data/data.txt"

type todo struct {
	ID        string `json;"id"`
	Item      string `json;"item"`
	Completed bool   `json;"completed"`
}

type TodoCache map[string]*todo

type CacheRWM struct {
	sync.RWMutex
	data map[string]*todo
}

type Shard struct {
	sync.RWMutex
	data map[string]*todo
}

type ShardMap []*Shard

func NewTodoCache() TodoCache {
	todoCache := make(TodoCache)

	todos, _ := GetTodoFromFile()

	for i := range todos {
		todoCache[todos[i].ID] = &todos[i]
	}

	return todoCache
}

func NewTodoCacheRWM() CacheRWM {
	return CacheRWM{
		data: make(map[string]*todo),
	}
}
func NewShardMap(n int) ShardMap {
	shards := make([]*Shard, n)

	for i := 0; i < n; i++ {
		shards[i] = &Shard{
			data: make(map[string]*todo),
		}
	}
	return shards
}

func (m TodoCache) Get(key string) (*todo, bool) {
	val := m[key]
	return val, val != nil
}

func (m TodoCache) Set(key string, val *todo) {
	m[key] = val
}

func (m *CacheRWM) Get(key string) (*todo, bool) {
	m.RLock()
	defer m.RUnlock()

	val := m.data[key]
	return val, val != nil
}

func (m *CacheRWM) Set(key string, val *todo) {
	m.Lock()
	defer m.Unlock()

	m.data[key] = val
}

func (m ShardMap) getShardIndex(key string) int {
	checksum := sha1.Sum([]byte(key))
	hash := int(checksum[0])

	return hash % len(m)
}

func (m ShardMap) getShard(key string) *Shard {
	i := m.getShardIndex(key)
	return m[i]
}

func (m ShardMap) Get(key string) (*todo, bool) {
	shard := m.getShard(key)

	shard.RLock()
	defer shard.RUnlock()

	val := shard.data[key]
	return val, val != nil
}

func (m ShardMap) Set(key string, val *todo) {
	shard := m.getShard(key)

	shard.Lock()
	defer shard.Unlock()

	shard.data[key] = val
}

func (m *CacheRWM) FillWithData() {

	todos, _ := GetTodoFromFile()

	for i := range todos {
		m.Set(todos[i].ID, &todos[i])
	}

}

func (m ShardMap) FillWithData() {

	todos, _ := GetTodoFromFile()

	for i := range todos {
		go func(i int, val todo) {
			m.Set(fmt.Sprint(todos[i].ID), &todos[i])
		}(i, todos[i])
	}

}

func NewAndFillCacheRWM() *CacheRWM {
	data := NewTodoCacheRWM()

	todos, _ := GetTodoFromFile()

	for i := range todos {
		data.Set(fmt.Sprint(todos[i].ID), &todos[i])
	}
	return &data
}

func NewAndFillShardMap() *ShardMap {
	data := NewShardMap(10)

	todos, _ := GetTodoFromFile()

	for i := range todos {
		data.Set(fmt.Sprint(todos[i].ID), &todos[i])
	}

	return &data
}
